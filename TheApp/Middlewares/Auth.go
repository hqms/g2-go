package Middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Service"
)

func AuthorizeUser() gin.HandlerFunc{
	return func(context *gin.Context) {
		//set when token from user validated
		var user,_ = context.Get("user")
		if user != ""{
			method := context.Request.Method
			allowed := false
			for _, auth := range Service.StaticAuthService(){
				if user == auth.Email{
					for _, perm := range auth.Permission{
						if method == perm{
							allowed = true
						}
					}
				}
			}

			if !allowed{
				context.AbortWithStatus(http.StatusUnauthorized)
			}
		}
		return
	}
}
