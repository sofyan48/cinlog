package routes

import "github.com/gin-gonic/gin"

// LOOGERROUTES ...
const LOOGERROUTES = VERSION + "/logger"

func (rLoader *V1RouterLoader) initLogger(router *gin.Engine) {
	group := router.Group(LOOGERROUTES)
	group.POST("", rLoader.Logger.CreateLogger)
	group.POST("get", rLoader.Logger.GetLogger)
}
