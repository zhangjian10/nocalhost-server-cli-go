package main

import (
	"fmt"
	"nocalhost/server/cmd"
	"os"
)

func main() {
	// email := os.Getenv("NH_SERVER_EMAIL")
	// password := os.Getenv("NH_SERVER_PASSWORD")

	// assert.NotEmpty(email, "email")
	// assert.NotEmpty(password, "password")

	// api.Login(email, password)

	// devSpace := api.CreateVcluster(1)

	// api.DeleteDevSpace(devSpace.ID)

	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	rootCmd := cmd.NewCmdRoot("", "")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}
