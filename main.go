package main

import (
	"fmt"

	"github.com/kdd4/go_final_project/pkg/db"
	"github.com/kdd4/go_final_project/pkg/server"
)

const webDir = "web"

func main() {
	err := db.Init("scheduler.db")

	if err != nil {
		fmt.Printf("Error of database initializing: %s\n", err.Error())
		return
	}

	defer func(){
		err = db.Close()

		if err != nil {
			fmt.Printf("Error while closing database connection: %s", err.Error())
		}
	}()

	err = server.Run(webDir)

	if err != nil {
		fmt.Printf("Listen and serve server error: %s\n", err.Error())
		return
	}
}