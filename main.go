package main

import (
	"fmt"
	"nocalhost/server/api"
	"os"
)

func main() {
	email := os.Getenv("NH_SERVER_EMAIL")
	password := os.Getenv("NH_SERVER_PASSWORD")

	fmt.Println(email)
	fmt.Println(password)

	api.Login(email, password)
}
