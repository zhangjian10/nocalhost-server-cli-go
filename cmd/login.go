package cmd

import (
	"nocalhost/server/api"
	"nocalhost/server/cmd/constants"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

type loginOpts struct {
	email    string
	password string
	hostname string
}

func login(opts *loginOpts) error {
	api.SetParameters(opts.hostname, "")

	r, err := api.Login(opts.email, opts.password)

	if err != nil {
		return err
	}

	cfg := ini.Empty()
	sec, _ := cfg.GetSection(ini.DefaultSection)

	sec.NewKey("host", opts.hostname)
	sec.NewKey("token", r.Token)

	return cfg.SaveTo(constants.TokenPath)
}

func newCmdLogin() *cobra.Command {
	opts := &loginOpts{}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in to nocalhost server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return login(opts)
		},
	}

	cmd.PersistentFlags().BoolP("help", "", false, "help for this command")

	cmd.Flags().StringVarP(&opts.email, "email", "u", "", "Email address")
	cmd.MarkFlagRequired("email")

	cmd.Flags().StringVarP(&opts.password, "password", "p", "", "Email password")
	cmd.MarkFlagRequired("password")

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "h", "", "The hostname of the nocalhost server")
	cmd.MarkFlagRequired("hostname")

	return cmd
}
