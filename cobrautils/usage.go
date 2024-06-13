package cobrautils

import (
	_ "embed"

	"github.com/Hayao0819/nahi/tputils"
	"github.com/spf13/cobra"
)

//go:embed template.txt
var HelpTemplate string

type HelpLabels struct {
	Usage               string
	Aliases             string
	Examples            string
	AvailableCommands   string
	AdditionalCommands  string
	Flags               string
	GlobalFlags         string
	AddtionalHelpTopics string
	UseHelp             string
}

func (l *HelpLabels) SetUsage(usage string) *HelpLabels {
	l.Usage = usage
	return l
}
func (l *HelpLabels) SetAliases(aliases string) *HelpLabels {
	l.Aliases = aliases
	return l
}
func (l *HelpLabels) SetExamples(examples string) *HelpLabels {
	l.Examples = examples
	return l
}
func (l *HelpLabels) SetAvailableCommands(availableCommands string) *HelpLabels {
	l.AvailableCommands = availableCommands
	return l
}
func (l *HelpLabels) SetAdditionalCommands(additionalCommands string) *HelpLabels {
	l.AdditionalCommands = additionalCommands
	return l
}
func (l *HelpLabels) SetFlags(flags string) *HelpLabels {
	l.Flags = flags
	return l
}
func (l *HelpLabels) SetGlobalFlags(globalFlags string) *HelpLabels {
	l.GlobalFlags = globalFlags
	return l
}
func (l *HelpLabels) SetAddtionalHelpTopics(addtionalHelpTopics string) *HelpLabels {
	l.AddtionalHelpTopics = addtionalHelpTopics
	return l
}
func (l *HelpLabels) SetUseHelp(useHelp string) *HelpLabels {
	l.UseHelp = useHelp
	return l
}

var defaultHelpLabels *HelpLabels = &HelpLabels{
	Usage:               "Usage",
	Aliases:             "Aliases",
	Examples:            "Examples",
	AvailableCommands:   "Available Commands",
	AdditionalCommands:  "Additional Commands",
	Flags:               "Flags",
	GlobalFlags:         "Global Flags",
	AddtionalHelpTopics: "Additional help topics",
	UseHelp:             `Use "{{.CommandPath}} [command] --help" for more information about a command.`,
}

func DefaultHelpLabels() *HelpLabels {
	return defaultHelpLabels

}

func GenerateTemplate(label *HelpLabels) string {
	buf, err := tputils.ApplyToText(HelpTemplate, label)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func UseCustomizableHelpTemplate(cmd *cobra.Command, label *HelpLabels) {
	tmpl := GenerateTemplate(label)
	//fmt.Println(tmpl)
	cmd.SetUsageTemplate(tmpl)
}
