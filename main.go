package main

import (
	"checkout-services/config"
	"checkout-services/routes"
	"fmt"
	"log"
	"os"
)

func main() {

	// tambahan
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// end

	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":" + port))
	fmt.Println("Toko Online")
}
