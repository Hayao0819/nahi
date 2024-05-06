package cobrautils

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/spf13/cobra"
)

func GetInternalAllSubCmds(cmd *cobra.Command) []*cobra.Command {
	v := reflect.ValueOf(cmd)

	commandsField := v.Elem().FieldByName("commands")
	commandsPtr := (*[]*cobra.Command)(unsafe.Pointer(commandsField.UnsafeAddr()))

	// Now you can access the private "commands" field
	commands := *commandsPtr
	return commands
}

func GetInternalSubCmdByName(cmd *cobra.Command, name string) *cobra.Command {
	cmds := GetInternalAllSubCmds(cmd)

	for _, c := range cmds {
		cname := strings.Split(c.Use, " ")[0]
		if cname == name {
			return c
		}
	}
	return nil
}

func GetHelpCommand(cmd *cobra.Command) *cobra.Command {
	copied := *cmd

	copied.InitDefaultHelpCmd()
	subcmds := GetInternalSubCmdByName(&copied, "help")
	if subcmds == nil {
		return nil
	}
	return subcmds
}
