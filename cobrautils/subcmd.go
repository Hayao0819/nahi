package cobrautils

import "github.com/spf13/cobra"

type Registory []*cobra.Command

// サブコマンドを追加する
func (r *Registory) RegisterSubCmd(cmds ...*cobra.Command) {
	*r = append(*r, cmds...)
}

// サブコマンドの一覧を取得する
func (r *Registory) GetSubCmds() []*cobra.Command {
	return *r
}

// サブコマンドをRootに追加する
func (r *Registory) BindSubCmds(root *cobra.Command) {
	root.AddCommand((*r)...)
}

var defaultRegistory Registory

func RegisterSubCmd(cmds ...*cobra.Command) {
	defaultRegistory.RegisterSubCmd(cmds...)
}

func GetSubCmds() []*cobra.Command {
	return defaultRegistory.GetSubCmds()
}

func BindSubCmds(root *cobra.Command) {
	defaultRegistory.BindSubCmds(root)
}
