package devspace

import (
	"github.com/spf13/cobra"
)

func NewDevSpaceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devspace",
		Short: "Management of devspace",
	}

	cmd.AddCommand(newCreateCmd())
	cmd.AddCommand(newDeleteCmd())

	return cmd
}
