package main

import (
	"fmt"
	"os"
)

func getToken() string {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("No token provided")
		os.Exit(1)
	}
	return token
}
