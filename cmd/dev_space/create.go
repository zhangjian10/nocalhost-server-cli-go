package devspace

import (
	"fmt"
	"nocalhost/server/api"
	"os"

	"github.com/spf13/cobra"
)

type createOpts struct {
	vcluster bool
	id       int64
}

func create(opts *createOpts) error {

	r, err := api.CreateVcluster(opts.id, opts.vcluster)

	if err != nil {
		return err
	}

	Kubeconfig, err := api.GetKubeconfig(r.ID)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "ID=\"%v\"\n", r.ID)
	fmt.Fprintf(os.Stdout, "KUBECONFIG=\"%s\"\n", *Kubeconfig)

	return nil
}

// go run main.go devspace create --id 1 >out && source out
func newCreateCmd() *cobra.Command {
	opts := &createOpts{}

	c := &cobra.Command{
		Use:   "create",
		Short: "Create Devspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return create(opts)
		},
	}

	c.Flags().BoolVarP(&opts.vcluster, "vcluster", "", true, "is vcluster")

	c.Flags().Int64VarP(&opts.id, "id", "", 0, "cluster_id")
	c.MarkFlagRequired("id")

	return c
}
