package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Models"
)

// Get Users godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Models.User
// @Router /userapi/user [get]
func GetUsers(c *gin.Context)  {
	var user[]Models.User
	err := Models.GetAllUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, user)
		}else{
			c.XML(http.StatusOK, user)
		}
	}
}

// Create Users godoc
// @Summary Create an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} Models.User
// @Router /admin/user [POST]
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