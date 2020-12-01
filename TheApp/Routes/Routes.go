package Routes

import (

	"github.com/gin-gonic/gin"
	"theapp/Controllers"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()

	usergroup := r.Group("/userapi")
	{
		usergroup.GET("user", Controllers.GetUsers)
		usergroup.GET("user/:id", Controllers.GetUserByID)
	}

	admingroup := r.Group("/admin",  gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"bar":   "foo",
	}))
	{
		admingroup.POST("user", Controllers.CreateUser)
		admingroup.PUT("user/:id", Controllers.UpdateUser)
		admingroup.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}