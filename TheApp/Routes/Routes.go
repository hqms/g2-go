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
	}
	return r
}