package route

import (
	"cse-foodcourt/src/controller"
	"cse-foodcourt/src/middleware"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	r.POST("/signup", controller.Signup)
	r.GET("/me/info/:id", middleware.RequireAuth, controller.GetInfo)
}
