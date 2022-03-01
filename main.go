package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/JamsMendez/starwars/routers"
)

func main() {
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api := server.Group("api")
	rsDev := routers.GQLDev{API: api}
	rs := routers.GQL{API: api}
	rsDev.New()
	rs.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Println("ServerHTTP.ERROR: ", err)
	}
}
