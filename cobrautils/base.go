package cobrautils

import "github.com/spf13/cobra"

type cmd cobra.Command

func CmdUtils(c *cobra.Command) *cmd {
	return (*cmd)(c)
}

func (c *cmd) DisableCompletion() *cmd {
	c.CompletionOptions.DisableDefaultCmd = true
	return c
}

func (c *cmd) SetUsageLabel(label string) *cmd {
	print((*cobra.Command)(c).UsageTemplate())
	(*cobra.Command)(c).SetUsageTemplate(label)
	return c
}
