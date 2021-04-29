package models

import (
	"api_backend/pkg/setting"
	"api_backend/pkg/util"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	Db  *gorm.DB
)

// 用户表
type TxUser struct {
	ID        uint `gorm:"primary_key"`
	User string	`gorm:"column:user"`
	PassWord string `gorm:"column:password"`
	// 对应以太坊地址
	Address   string `gorm:"column:address"`
	// 对应私钥
	Key     string	`gorm:"column:key"`
}
// 用户交易表
type Txs struct {
	gorm.Model
	SendName string `gorm:column:sendname`
	ReciveName string `gorm:column:recivename`
	Hash  string    `gorm:column:hash`
	SendAddr  string    `gorm:"column:sendaddr"`
	ReciveAddr  string	 `gorm:"column:reciveaddr"`
	Value  string	`gorm:"column:value"`
	Typa   string   `gorm:"column:typa"`
}

func (this *TxUser) TableName() string{
	return "tx_user"
}

func (this *Txs) TableName() string{
	return "tx_s"
}

// 初始化 建表语句
func MysqlInit(){
	// 建表
	if Db.HasTable(&TxUser{}) {
	}else{
		Db.CreateTable(&TxUser{})
		Db.CreateTable(&Txs{})
	}
}

func InitMysql() {
	var err error
	var url,user,passwd string
	secEth, err := setting.Cfg.GetSection("mysql")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	url = secEth.Key("URL").String()
	user =  secEth.Key("USER").String()
	passwd  =  secEth.Key("PASSWD").String()

	Db, err = gorm.Open("mysql", user+":" +passwd+"@tcp("+url+")/work?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db.SingularTable(true)
	fmt.Println("connect mysql success")
	MysqlInit()
}

func AddUser(user,pass,addr,key string) error {
	 var u TxUser
	// 预先处理密码加密，然后存储在数据库
	 _pass := Base64Md5(pass)
	 u = TxUser{
		 User:     user,
		 PassWord: _pass,
		 Address:  addr,
		 Key:      key,
	 }
	 err := Db.Create(&u).Error
	 return  err
}

func Checklogin(user,pass string) (*TxUser,error){
	var u TxUser
	sql := "select * from tx_user where  user = ? "

	res := Db.Raw(sql,user).First(&u)
	if res.Error == nil {
		// 账户密码验证成功
		if len(u.PassWord) > 0 && (u.PassWord == Base64Md5(pass)) {
			return &u,nil
		}
	}
	fmt.Println(res.Error)
	return nil,res.Error
}

func CheckIsExist(user string) bool {

	var u TxUser
	Db.Where("user = ?", user).First(&u)
	// 该账户名已存在
	if u.PassWord != "" {
		return  true
	}
	return  false
}

func GetUserBytoken(token string) (*TxUser,error){
	res,err :=util.ParseToken(token)
	if err != nil {
		return nil,err
	}
	var u TxUser
	err = Db.Where("user = ? ",res.Username).First(&u).Error
	if err != nil {
		return  nil,err
	}
	return &u,nil
}

func GetUserByName(name string)(*TxUser,error) {
	var u  TxUser
	err := Db.Where("user = ? ",name).First(&u).Error
	if err != nil {
		return  nil,err
	}
	return &u,nil
}

func Transfer(from,to ,amount string) error {
	return nil
}


func AddTxs(sendName,reciveName,hash,from,to,value,typa string) error{
	var tx Txs
	tx = Txs{
		SendName: sendName,
		ReciveName:reciveName,
		Hash:hash,
		SendAddr:from,
		ReciveAddr: to,
		Value:value,
		Typa:typa,
	}
	err := Db.Create(&tx).Error
	return err
}

func GetAllTx(addr string) *[]Txs {
	fmt.Println(addr)
	var tx []Txs
	 Db.Where("sendaddr = ?",addr).Or("reciveaddr = ?",addr).Find(&tx)
	//fmt.Println(err.Error())
	return &tx
}