package main

import (
	"checkout-services/config"
	"checkout-services/routes"
	"fmt"
	// "log"
	// "os"
)

func main() {

	// tambahan
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }
	// akhir dari tambahan

	config.InitDB()
	config.InitMigrate()
	e := routes.New()
	// e.Logger.Fatal(e.Start(":" + port))
	e.Logger.Fatal(e.Start(":" + "8080"))
	fmt.Println("Toko Online")
}
