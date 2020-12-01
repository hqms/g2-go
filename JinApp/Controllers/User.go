package Controllers

import (
	"github.com/gin-gonic/gin"
	"jinapp/Models"
	"net/http"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /accounts/{id} [get]
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