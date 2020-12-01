package Middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"theapp/Service"
)

func AuthorizeJWT() gin.HandlerFunc  {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == ""{
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := Service.JwtAuthService().ValidateToken(tokenString)
		if token.Valid{
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			context.Set("user", claims["name"])
			return
		}else{
			fmt.Println(err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
