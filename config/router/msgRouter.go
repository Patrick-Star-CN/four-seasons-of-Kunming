package router

import (
	"four-seasons-of-Kunming/app/controllers/msgController"
	"github.com/gin-gonic/gin"
)

func msgRouterInit(r *gin.RouterGroup) {
	r.GET("/msg", msgController.GetMsg)
	r.POST("/msg", msgController.CreateMsg)
}
