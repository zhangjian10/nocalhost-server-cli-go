package cmd

import (
	"errors"
	"io/ioutil"
	"nocalhost/server/api"
	"nocalhost/server/cmd/constants"
	devspace "nocalhost/server/cmd/dev_space"
	"nocalhost/server/utils"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

func preRunE(cmd *cobra.Command) error {

	if cmd.Name() == "login" {
		return nil
	}

	if exists, _ := utils.PathExists(constants.TokenPath); !exists {
		return errors.New("you need to log in and use, you can execute: nh-server login")
	}

	data, _ := ioutil.ReadFile(constants.TokenPath)

	cfg, _ := ini.LoadSources(ini.LoadOptions{
		AllowNonUniqueSections: true,
	}, data)

	sec, _ := cfg.GetSection(ini.DefaultSection)

	host, _ := sec.GetKey("host")
	token, _ := sec.GetKey("token")

	api.SetParameters(host.Value(), token.Value())

	return nil
}

func NewCmdRoot(version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "nh-server",
		Long: "Nocalhost server cli",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return preRunE(cmd)
		},
	}

	cmd.InitDefaultHelpCmd()

	cmd.AddCommand(devspace.NewDevSpaceCmd())
	cmd.AddCommand(newCmdLogin())
	cmd.AddCommand(newCmdVersion())

	return cmd
}
