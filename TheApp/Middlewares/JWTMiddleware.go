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
		const BEARER_SCHEMA = "Bearer"
		authHeader := context.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := Service.JwtAuthService().ValidateToken(tokenString)
		if token.Valid{
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		}else{
			fmt.Println(err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
