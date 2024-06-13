package cobrautils

import (
	_ "embed"

	"github.com/Hayao0819/nahi/tputils"
	"github.com/spf13/cobra"
)

//go:embed template.txt
var HelpTemplate string

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
