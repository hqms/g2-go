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
		usergroup.POST("user", Controllers.CreateUser)
		usergroup.GET("user/:id", Controllers.GetUserByID)
		usergroup.PUT("user/:id", Controllers.UpdateUser)
		usergroup.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}