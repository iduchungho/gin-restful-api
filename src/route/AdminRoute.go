package route

import (
	"cse-foodcourt/src/controller"
	"cse-foodcourt/src/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.Engine) {
	r.POST("/admin/signup", middleware.RequireAdminKey, controller.SignupAdmin)
}
