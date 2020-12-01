package Routes

import (
	"github.com/gin-gonic/gin"
	"jinapp/Controllers"
)

func SetupRouter() *gin.Engine  {
	r := gin.New()

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