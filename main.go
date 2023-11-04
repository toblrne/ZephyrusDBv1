package main

import (
	"fmt"

	"github.com/toblrne/ZephyrusDB/api"
	"github.com/toblrne/ZephyrusDB/db"
)

func main() {

	databaseDir := "./data"
	options := &db.Options{}

	// init db
	database, err := db.New(databaseDir, options)
	if err != nil {
		fmt.Println("Error", err)
	}

	handler := api.NewHandler(database)

	router := api.NewRouter(handler)

	router.Run(":8080")
}
