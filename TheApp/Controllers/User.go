package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Models"
)

func GetUsers(c *gin.Context)  {
	var user[]Models.User
	err := Models.GetAllUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.JSON(http.StatusNotFound, user)
	}

	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}