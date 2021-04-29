package system

import (
	"api_backend/models"
	"api_backend/pkg/app"
	"api_backend/pkg/e"
	"api_backend/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type login struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}
type RegisterParam struct {
	User     string `form:"user" binding:"required,min=6,max=12" `
	Password string `form:"password" binding:"required,min=6,max=12"`
}

type UserInfoParam struct {
	UserID string `valid:"Required; MaxSize(50)"`
	UserType string `valid:"Required; MaxSize(50)"`
}

// 注册
func Register(c *gin.Context){
	var param RegisterParam
	err := c.Bind(&param)
	if err != nil {
		app.JsonError(c,e.INVALID_PARAMS,err.Error())
		return
	}

	if models.CheckIsExist(param.User) {
		app.JsonError(c,e.ERROR,"账户已存在")
		return
	}
	// 生成以太坊地址
	addr,key := NewAddress()
	err = models.AddUser(param.User,param.Password,addr,key)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}


	app.JsonSuccess(c,e.SUCCESS,"")
}

// 登录
func Login(c *gin.Context) {
	var param RegisterParam
	err := c.Bind(&param)
	if err != nil {
		app.JsonError(c,e.INVALID_PARAMS,err.Error())
	}
	_, err= models.Checklogin(param.User,param.Password)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	token , err := util.GenerateToken(param.User,param.Password)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"token":token,
	})
}

// 充值  给自己充值
func Recharge(c *gin.Context) {
	token := c.Query("token")
	u,err := models.GetUserBytoken(token)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
	}
	if u == nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}
	// 自动充值
	txhash,err := RechargeToMe(u.Address)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	// 添加数据库
	err = models.AddTxs("admin",u.User,txhash,models.RechargeAddr,u.Address,models.RechargeAmount,"transfer")
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}

	app.JsonSuccess(c,e.SUCCESS,"")
}


// 获取用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	u,err := models.GetUserBytoken(token)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
	}
	if u == nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}

	value,err:= GetAddrBalance(u.Address)
	if err != nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}
	valuerf ,err := strconv.ParseFloat(value, 64)
	if err != nil {
		app.JsonError(c,e.ERROR,err)
		return
	}
	// 保留两位小数
	r := fmt.Sprintf("%.2f", valuerf/1000000000000000000)

	res := models.GetAllTx(u.Address)

	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"addr":u.Address,
		"value":r,
		"tx_list":res,
	})
}



func Transfer(c *gin.Context) {

	token := c.Query("token")
	to := c.Query("to")
	value := c.Query("value")

	u,err := models.GetUserBytoken(token)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
	}
	if u == nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}

	if !models.CheckIsExist(to) {
		app.JsonError(c,e.ERROR,"接收地址")
		return
	}

	t,err :=models.GetUserByName(to)
	if err != nil {
		app.JsonError(c,e.ERROR,"接收地址")
		return
	}

	tx,err := Trans(u.Address,u.Key, t.Address,value)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	err = models.AddTxs(u.User,t.User,tx,u.Address,t.Address,value,"transfer")
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}

	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"txhash": tx,
	})
	return

}



func GetTx(c *gin.Context){
	hash := c.Query("hash")

	res,err :=GetTxByHash(hash)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"tx":res,
	})
}


func Set(c *gin.Context) {
	token := c.Query("token")
	value := c.Query("value")
	u,err := models.GetUserBytoken(token)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
	}
	if u == nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}
	tx,err :=Onchain(int(u.ID),value)
	if err != nil {
		app.JsonError(c,e.ERROR,"合约set失败")
		return
	}
	err = models.AddTxs(u.User,"智能合约",tx,u.Address,models.ContractAddr,value,"set")
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}

	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"txhash": tx,
	})

}

//  查询合约 不算交易 所以不需要更新数据库
func Get(c *gin.Context) {
	token := c.Query("token")
	u,err := models.GetUserBytoken(token)
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
	}
	if u == nil {
		app.JsonError(c,e.ERROR,"无效的token")
		return
	}

	res,err :=GetChain(int(u.ID))
	if err != nil {
		app.JsonError(c,e.ERROR,err.Error())
		return
	}
	app.JsonSuccess(c,e.SUCCESS,gin.H{
		"value": res,
	})
}