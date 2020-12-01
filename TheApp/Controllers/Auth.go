package Controllers

import (
	"github.com/gin-gonic/gin"
	"theapp/Dto"
	"theapp/Service"
)

type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService Service.LoginService
	jwtService Service.JWTService
}

func LoginHandler(loginService Service.LoginService, jwtservice Service.JWTService)  LoginController{
	return &loginController{
		loginService: loginService,
		jwtService:   jwtservice,
	}
}

func (controller *loginController)Login(c *gin.Context)  string{
	var credential Dto.LoginCredentials
	err := c.ShouldBind(&credential)

	if err != nil{
		return "No Data Found"
	}

	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated{
		return controller.jwtService.GenerateToken(credential.Email, true)
	}

	return ""
}