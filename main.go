package main

import (
	"fmt"
	"nocalhost/server/api"
	"nocalhost/server/utils/assert"
	"os"
)

func main() {
	email := os.Getenv("NH_SERVER_EMAIL")
	password := os.Getenv("NH_SERVER_PASSWORD")

	assert.NotEmpty(email)
	assert.NotEmpty(password)

	api.Login(email, password)

	id := api.CreateVcluster(1)

	fmt.Sprintf("id: %v", id)
}
