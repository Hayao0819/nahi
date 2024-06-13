package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func helpTemplateCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "help-template output",
		Short: "Escape all '{{' and '}}' for golang template",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			re, err := regexp.Compile(`\{\{([^\}\}]*)\}\}`)
			if err != nil {
				return err
			}
			tmpl := ""
			{
				blankCmd := cobra.Command{}
				tmpl = blankCmd.UsageTemplate()
			}
			escaped := re.ReplaceAllString(tmpl, "{{\"{{\"}}$1{{\"}}\"}}")

			// カスタム用の加工
			escaped = strings.Replace(escaped, "Usage:", "{{.Usage}}:", 1)
			escaped = strings.Replace(escaped, "Aliases:", "{{.Aliases}}:", 1)
			escaped = strings.Replace(escaped, "Examples:", "{{.Examples}}:", 1)
			escaped = strings.Replace(escaped, "Available Commands:", "{{.AvailableCommands}}:", 1)
			escaped = strings.Replace(escaped, "Additional Commands:", "{{.AdditionalCommands}}:", 1)
			escaped = strings.Replace(escaped, "Flags:", "{{.Flags}}:", 1)
			escaped = strings.Replace(escaped, "Global Flags:", "{{.GlobalFlags}}:", 1)
			escaped = strings.Replace(escaped, "Additional help topics:", "{{.AddtionalHelpTopics}}:", 1)
			escaped = strings.Replace(escaped, re.ReplaceAllString(`Use "{{.CommandPath}} [command] --help" for more information about a command.`,"{{\"{{\"}}$1{{\"}}\"}}" ), "{{.UseHelp}}", 1)

			fmt.Println(escaped)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return &cmd
}

func init() {
	reg.RegisterSubCmd(helpTemplateCmd())
}
