package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var (
	OsArch    = ""
	Version   = ""
	GitCommit = ""
	BuildTime = ""
	Branch    = ""
)

func reformatDate(buildTime string) string {
	t, errTime := time.Parse(time.RFC3339Nano, buildTime)
	if errTime == nil {
		return t.Format(time.ANSIC)
	}
	return buildTime
}

func newCmdVersion() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Nocalhost Server CLI\n")
			fmt.Printf("    Version: %s\n", Version)
			fmt.Printf("    Branch: %s\n", Branch)
			fmt.Printf("    Git commit: %s\n", GitCommit)
			fmt.Printf("    Built time: %s\n", reformatDate(BuildTime))
			fmt.Printf("    Built OS/Arch: %s\n", OsArch)
			fmt.Printf("    Built Go version: %s\n", runtime.Version())
		},
	}

	return cmd
}
