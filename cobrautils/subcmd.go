package cobrautils

import "github.com/spf13/cobra"

type Registory []*cobra.Command

// サブコマンドを追加する
func (r *Registory) Add(cmds ...*cobra.Command) {
	*r = append(*r, cmds...)
}

// サブコマンドの一覧を取得する
func (r *Registory) Get() []*cobra.Command {
	return *r
}

// サブコマンドをRootに追加する
func (r *Registory) Bind(root *cobra.Command) {
	root.AddCommand((*r)...)
}

var defaultRegistory Registory

func AddSubCmds(cmds ...*cobra.Command) {
	defaultRegistory.Add(cmds...)
}

func GetSubCmds() []*cobra.Command {
	return defaultRegistory.Get()
}

func BindSubCmds(root *cobra.Command) {
	defaultRegistory.Bind(root)
}
