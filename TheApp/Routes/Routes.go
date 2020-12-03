package Routes

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"theapp/Controllers"
	"theapp/Middlewares"
	"theapp/Service"
	swaggerFiles "github.com/swaggo/files"
	_ "theapp/docs"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()
	r.Use(Middlewares.SignRequest)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://9a1f5d57fb5545b1bec8fa373a7f3089@o485449.ingest.sentry.io/5541065",
	});  err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	var loginService Service.LoginService = Service.StaticLoginService()
	var jwtService Service.JWTService = Service.JwtAuthService()
	var loginController Controllers.LoginController = Controllers.LoginHandler(loginService, jwtService)

	r.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != ""{
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}else{
			context.JSON(http.StatusUnauthorized, nil)
		}
	})

	sentrygroup := r.Group("/sentry")
	{
		sentrygroup.GET("error", Controllers.ErrorSentry)
	}

	usergroup := r.Group("/userapi")
	{
		usergroup.GET("user", Controllers.GetUsers)

	}

	admingroup := r.Group("/admin",  Middlewares.AuthorizeJWT(), Middlewares.AuthorizeUser())
	{
		admingroup.GET("user/:id", Controllers.GetUserByID)
		admingroup.POST("user", Controllers.CreateUser)
		admingroup.PUT("user/:id", Controllers.UpdateUser)
		admingroup.DELETE("user/:id", Controllers.DeleteUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}