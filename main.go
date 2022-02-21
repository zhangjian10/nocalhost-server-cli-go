package main

import (
	"fmt"
	"nocalhost/server/cmd"
	"os"
)

func main() {
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
