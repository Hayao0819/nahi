package cobrautils

import "github.com/spf13/cobra"

// 指定されたサブコマンドを、現在のcmdと同じ環境で実行します
func CallCmd(me *cobra.Command, target cobra.Command, args ...string) error {
	target.SetOut(me.OutOrStdout())
	target.SetErr(me.OutOrStderr())
	target.SetIn(me.InOrStdin())
	target.SetArgs(args)
	return target.Execute()
}
