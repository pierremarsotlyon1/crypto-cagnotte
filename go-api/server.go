package main

import (
	"crypto-cagnotte/go-api/app"
	"crypto-cagnotte/go-api/app/auth"
	"crypto-cagnotte/go-api/app/cagnotte"
	"crypto-cagnotte/go-api/app/coinbase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	app.Init()

	e := echo.New()
	e.POST("/login", auth.Login)
	e.POST("/register", auth.Register)
	e.POST("/coinbase/notification", coinbase.ReceiveNotification)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(auth.SecretJwt),
	}))

	cagnottesApi := apiGroup.Group("/cagnottes")
	cagnottesApi.POST("", cagnotte.Add)
	cagnottesApi.GET("/:id", cagnotte.Get)
	cagnottesApi.PUT("/:id/close", cagnotte.Close)
	cagnottesApi.POST("/:id/withdraw", cagnotte.Withdraw)

	e.Logger.Fatal(e.Start(":1323"))
}
