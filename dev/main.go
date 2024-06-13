package main

import (
	"os"

	cu "github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/nahi/mobra"
	"github.com/spf13/cobra"
)

var reg = cu.Registory{}

func root() *cobra.Command {
	root := mobra.New("nahi-dev").
		Short("dev tool for nahi").
		BindSubCmds(&reg).
		DisableDefaultCmd().
		HideUsage().
		Cobra()

	return root
}

func main() {
	if err := root().Execute(); err != nil {
		os.Exit(1)
	}
}
