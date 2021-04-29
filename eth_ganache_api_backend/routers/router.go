package routers

import (
	"api_backend/middleware"
	"api_backend/middleware/jwt"
	"api_backend/routers/api/system"

	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	r := gin.New()
	//跨域
	r.Use(middleware.Cors())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.POST("/api/login",system.Login)
	r.POST("/api/register",system.Register)

	apiv1Base := r.Group("/api/system")
	apiv1Base.Use(jwt.JWT())
	{
		// 系统的相关接口路由
		// 获取用户信息
		apiv1Base.POST("/userinfo",system.UserInfo)
		//充值
		apiv1Base.POST("/recharge",system.Recharge)
		apiv1Base.POST("/transfer",system.Transfer)
		// 查询交易
		apiv1Base.POST("/gettx",system.GetTx)

		// 往智能合约存储一个数据
		apiv1Base.POST("/set",system.Set)
		// 从智能合约获取数据
		apiv1Base.POST("/get",system.Get)
	}

	return r
}
