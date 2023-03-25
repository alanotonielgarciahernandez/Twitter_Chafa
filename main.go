package main

import (
	"log"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/handlers"
)

func main() {
	if db.CheckDBConnection() == 0 {
		log.Fatal("Without connection with DB")
		return
	}
	handlers.Managers()
}
