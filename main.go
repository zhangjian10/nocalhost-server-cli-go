package main

import (
	"log"
	"nocalhost/server/api"
	"nocalhost/server/utils"
	"nocalhost/server/utils/assert"
	"os"
)

func main() {
	email := os.Getenv("NH_SERVER_EMAIL")
	password := os.Getenv("NH_SERVER_PASSWORD")

	assert.NotEmpty(email, "email")
	assert.NotEmpty(password, "password")

	api.Login(email, password)

	vcluster := api.CreateVcluster(1)

	log.Println(utils.ToJson(vcluster))
}
