package main

import (
	"log"
	"os"

	mws "github.com/berrylradianh/go-jewelry/middlewares"
	rts "github.com/berrylradianh/go-jewelry/routes"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	log.Println("Starting the application...")
	e := rts.InitRoutes()
	mws.LogMiddleware(e)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
