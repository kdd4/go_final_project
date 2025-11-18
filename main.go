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

	err = server.Run(webDir)

	if err != nil {
		fmt.Printf("Listen and serve server error: %s\n", err.Error())
	}
}