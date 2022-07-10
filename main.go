package main

import (
	"checkout-services/config"
	"checkout-services/routes"
	"fmt"
	"log"
)

func main() {

	// tambahan
	port := "8090" // os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// akhir dari tambahan

	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	e.Logger.Fatal(e.Start(":" + port))
	fmt.Println("Toko Online")
}
