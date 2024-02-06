package cobrautils

import "github.com/spf13/cobra"

var subCmds []*cobra.Command

// サブコマンドを追加する
func AddSubCmds(cmds ...*cobra.Command) {
	subCmds = append(subCmds, cmds...)
}

// サブコマンドの一覧を取得する
func GetSubCmds() []*cobra.Command {
	return subCmds
}

// サブコマンドをRootに追加する
func AddSubCmdsToRoot(root *cobra.Command) {
	root.AddCommand(subCmds...)
}
