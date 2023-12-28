package main

//import routes
import (
	"belajar_golang/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	/**
	 * Load .env file
	 */

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	r := routes.Routes()
	r.Run(":" + port)
}
