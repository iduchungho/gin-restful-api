package middleware

import "github.com/gin-gonic/gin"

func RequireAuth(c *gin.Context) {
	// require auth to do something
	//TODO
	c.Next()
}

func RequireAdmin(c *gin.Context) {
	// require admin to do something
	//TODO
	c.Next()
}

func RequireAdminKey(c *gin.Context) {
	// require admin key to do something
	//TODO
	c.Next()
}
