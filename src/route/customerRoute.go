package route

import (
	"cse-foodcourt/src/controller"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	r.GET("/api/getAllCustomer", controller.GetAllCustomer)
}
