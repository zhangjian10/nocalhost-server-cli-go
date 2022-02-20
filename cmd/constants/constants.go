package constants

import (
	"nocalhost/server/utils"
	"os"
	"path/filepath"
)

var (
	dirname, _   = os.UserHomeDir()
	NHServerPath = filepath.Join(dirname, ".nh-server")
	TokenPath    = filepath.Join(NHServerPath, "token.ini")
)

func init() {
	utils.Mkdir(NHServerPath)
}
