package router

import (
	"DecoderProtocol/controller"
	"github.com/gin-gonic/gin"
)

// InitRouter 方法创建 gin 路由，并设置相关路由规则
func InitRouter() *gin.Engine {
	// 创建 gin router
	r := gin.Default()

	// 添加路由
	r.GET("/getarmy", controller.GetArmy)
	r.GET("/getrarity", controller.GetRarity)
	r.GET("/getcombatpoints", controller.GetCombatPoints)
	r.GET("/getarmybycvc", controller.GetArmyByCVC)
	r.GET("/getarmybyunlockarena", controller.GetArmyByUnlockArena)

	return r
}
