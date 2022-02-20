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
	// output   string
	// force    bool
}

func create(opts *createOpts) error {
	// var isExists bool

	// if opts.output != "" {
	// 	if isExists, _ = utils.PathExists(opts.output); isExists && !opts.force {
	// 		return fmt.Errorf("%s already exists, please add the -f option to override", opts.output)
	// 	}
	// }

	r, err := api.CreateVcluster(opts.id, opts.vcluster)

	if err != nil {
		return err
	}

	Kubeconfig, err := api.GetKubeconfig(r.ID)

	if err != nil {
		return err
	}

	// if opts.output != "" {
	// 	if err = ioutil.WriteFile(fmt.Sprintf("%s.kubeconfig", opts.output), []byte(*Kubeconfig), 0666); err != nil {
	// 		return err
	// 	}

	// 	err = ioutil.WriteFile(fmt.Sprintf("%s.id", opts.output), []byte(fmt.Sprint(r.ID)), 0666)

	// } else {

	fmt.Fprintf(os.Stdout, "ID=\"%v\"\n", r.ID)
	fmt.Fprintf(os.Stdout, "KUBECONFIG=\"%s\"\n", *Kubeconfig)
	// }

	return nil
}

// go run main.go devspace create --id 1 >out && source out
func newCreateCmd() *cobra.Command {
	opts := &createOpts{}

	c := &cobra.Command{
		Use:  "create",
		Long: "create devspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			return create(opts)
		},
	}

	// c.Flags().StringVarP(&opts.output, "output", "o", "", "output id and kubeconfig suffix file")
	c.Flags().BoolVarP(&opts.vcluster, "vcluster", "", true, "is vcluster")
	// c.Flags().BoolVarP(&opts.force, "force", "f", false, "if the file already exists, overwrite file")

	c.Flags().Int64VarP(&opts.id, "id", "", 0, "cluster_id")
	c.MarkFlagRequired("id")

	return c
}
