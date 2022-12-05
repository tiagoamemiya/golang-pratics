package routes

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest(r *gin.Engine) {
	r.GET("/health", controllers.HealthHandler)
	r.GET("/test/:param", controllers.TestParamsUrl)
}
