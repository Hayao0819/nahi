package cobrautils

import "github.com/spf13/cobra"

var subCmds []*cobra.Command

func AddSubCmds(cmds ...*cobra.Command) {
	subCmds = append(subCmds, cmds...)
}

func GetSubCmds() []*cobra.Command {
	return subCmds
}

func AddSubCmdsToRoot(root *cobra.Command) {
	root.AddCommand(subCmds...)
}
