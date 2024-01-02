package main

//import routes
import (
	"belajar_golang/routes"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func main() {

	/**
	 * Load .env file
	 */

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	})

	port := os.Getenv("PORT")

	r := routes.Routes()
	r.Run(":" + port)
}
