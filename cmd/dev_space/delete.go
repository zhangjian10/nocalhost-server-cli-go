package devspace

import (
	"nocalhost/server/api"

	"github.com/spf13/cobra"
)

var id int64

// go run main.go devspace delete --id $ID
func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete devspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return api.DeleteDevSpace(id)
		},
	}
	cmd.Flags().Int64Var(&id, "id", 0, "ClusterId")

	cmd.MarkFlagRequired("id")
	return cmd
}
