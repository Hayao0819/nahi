package mobra

import "github.com/Hayao0819/nahi/cobrautils"

func (c *cmd) CustomHelp(label *cobrautils.HelpLabels) *cmd {
	cobrautils.UseCustomizableHelpTemplate(c.cc, label)
	return c
}
