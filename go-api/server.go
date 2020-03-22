package main

import (
	"crypto-cagnotte/go-api/app"
	"crypto-cagnotte/go-api/app/auth"
	"crypto-cagnotte/go-api/app/cagnotte"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	app.Init()

	e := echo.New()
	e.POST("/login", auth.Login)
	e.POST("/register", auth.Register)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(auth.SecretJwt),
	}))

	cagnottesApi := apiGroup.Group("/cagnottes")
	cagnottesApi.POST("", cagnotte.Add)
	cagnottesApi.GET("/:id", cagnotte.Get)

	e.Logger.Fatal(e.Start(":1323"))
}
