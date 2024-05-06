package cobrautils_test

import (
	"testing"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func TestGetInternalSubCmds(t *testing.T) {

	var root = &cobra.Command{
		Use:   "root",
		Short: "root command",
	}

	var sub1 = &cobra.Command{
		Use:   "sub1",
		Short: "sub1 command",
	}

	root.AddCommand(sub1)

	cmds := cobrautils.GetInternalAllSubCmds(root)
	if len(cmds) != 1 {
		t.Errorf("want 1, got %d", len(cmds))
	}
	for _, cmd := range cmds {
		t.Log(cmd.Use)
	}
}

func TestGetHelpCommand(t *testing.T) {

	var root = &cobra.Command{
		Use:   "root",
		Short: "root command",
	}

	var sub1 = &cobra.Command{
		Use:   "sub1",
		Short: "sub1 command",
	}

	root.AddCommand(sub1)

	help := cobrautils.GetHelpCommand(root)
	if help == nil {
		t.Error("help command not found")
		return
	}
	t.Log(help.Use)
}
