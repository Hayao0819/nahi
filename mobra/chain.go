package mobra

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

type cmd struct {
	cc *cobra.Command
}

func New(use string) *cmd {
	c := cmd{cc: &cobra.Command{}}
	return c.Use(use)
}

func (c *cmd) Use(use string) *cmd {
	c.cc.Use = use
	return c
}

func (c *cmd) Aliases(aliases []string) *cmd {
	c.cc.Aliases = aliases
	return c
}

// SetSuggestFor
func (c *cmd) SuggestFor(suggestFor []string) *cmd {
	c.cc.SuggestFor = suggestFor
	return c
}

func (c *cmd) Short(short string) *cmd {
	c.cc.Short = short
	return c
}

// SetGroupID
func (c *cmd) GroupID(groupID string) *cmd {
	c.cc.GroupID = groupID
	return c
}

func (c *cmd) Long(long string) *cmd {
	c.cc.Long = long
	return c
}

func (c *cmd) Example(example string) *cmd {
	c.cc.Example = example
	return c
}

func (c *cmd) ValidArgs(validArgs []string) *cmd {
	c.cc.ValidArgs = validArgs
	return c
}

func (c *cmd) ValidArgsFunction(validArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) *cmd {
	c.cc.ValidArgsFunction = validArgsFunction
	return c
}

func (c *cmd) Args(args cobra.PositionalArgs) *cmd {
	c.cc.Args = args
	return c
}

func (c *cmd) ArgAliases(argAliases []string) *cmd {
	c.cc.ArgAliases = argAliases
	return c
}

func (c *cmd) Deprecated(deprecated string) *cmd {
	c.cc.Deprecated = deprecated
	return c
}

func (c *cmd) Annotation(key, value string) *cmd {
	c.cc.Annotations[key] = value
	return c
}

func (c *cmd) Version(version string) *cmd {
	c.cc.Version = version
	return c
}

// Funcs

func (c *cmd) Run(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.Run = f
	return c
}

func (c *cmd) RunE(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.RunE = f
	return c
}

func (c *cmd) PreRun(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PreRun = f
	return c
}

func (c *cmd) PreRunE(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PreRunE = f
	return c
}

func (c *cmd) PersistentPreRun(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PersistentPreRun = f
	return c
}

func (c *cmd) PersistentPreRunE(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PersistentPreRunE = f
	return c
}

func (c *cmd) PostRun(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PostRun = f
	return c
}

func (c *cmd) PostRunE(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PostRunE = f
	return c
}

func (c *cmd) PersistentPostRun(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PersistentPostRun = f
	return c
}

func (c *cmd) PersistentPostRunE(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PersistentPostRunE = f
	return c
}

func (c *cmd) PersistentPreRunWithParent(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PersistentPreRun = cobrautils.WithParentPersistentPreRun(f)
	return c
}

func (c *cmd) PersistentPreRunEWithParent(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PersistentPreRunE = cobrautils.WithParentPersistentPreRunE(f)
	return c
}

func (c *cmd) PersistentPostRunWithParent(f func(cmd *cobra.Command, args []string)) *cmd {
	c.cc.PersistentPostRun = cobrautils.WithParentPersistentPostRun(f)
	return c
}

func (c *cmd) PersistentPostRunEWithParent(f func(cmd *cobra.Command, args []string) error) *cmd {
	c.cc.PersistentPostRunE = cobrautils.WithParentPersistentPostRunE(f)
	return c
}

// Hidden

func (c *cmd) Hidden(hidden bool) *cmd {
	c.cc.Hidden = hidden
	return c
}

func (c *cmd) Hide() *cmd {
	c.Hidden(true)
	return c
}

func (c *cmd) Show() *cmd {
	c.Hidden(false)
	return c
}

// Errors

func (c *cmd) SetSilenceErrors(silenceErrors bool) *cmd {
	c.cc.SilenceErrors = silenceErrors
	return c
}

func (c *cmd) PrintErrors() *cmd {
	c.SetSilenceErrors(false)
	return c
}

func (c *cmd) HideErrors() *cmd {
	c.SetSilenceErrors(true)
	return c
}

// Usage

func (c *cmd) SetSilenceUsage(silenceUsage bool) *cmd {
	c.cc.SilenceUsage = silenceUsage
	return c
}

func (c *cmd) PrintUsage() *cmd {
	c.SetSilenceUsage(false)
	return c
}

func (c *cmd) HideUsage() *cmd {
	c.SetSilenceUsage(true)
	return c
}

func (c *cmd) SetDisableFlagParsing(disableFlagParsing bool) *cmd {
	c.Cobra().DisableFlagParsing = disableFlagParsing
	return c
}

func (c *cmd) SetDisableAutoGenTag(disableAutoGenTag bool) *cmd {
	c.cc.DisableAutoGenTag = disableAutoGenTag
	return c
}

func (c *cmd) SetDisableFlagsInUseLine(disableFlagsInUseLine bool) *cmd {
	c.cc.DisableFlagsInUseLine = disableFlagsInUseLine
	return c
}

func (c *cmd) SetDisableSuggestions(disableSuggestions bool) *cmd {
	c.cc.DisableSuggestions = disableSuggestions
	return c
}

func (c *cmd) SetSuggestionsMinimumDistance(suggestionsMinimumDistance int) *cmd {
	c.cc.SuggestionsMinimumDistance = suggestionsMinimumDistance
	return c
}

func (c *cmd) DisableCompletion() *cmd {
	c.cc.CompletionOptions.DisableDefaultCmd = true
	return c
}

func (c *cmd) DisableHelpCommand() *cmd {
	c.cc.SetHelpCommand(&cobra.Command{})
	return c
}

// utils

func (c *cmd) DisableDefaultCmd() *cmd {
	return c.DisableCompletion().DisableHelpCommand()
}

func (c *cmd) Cobra() *cobra.Command {
	return c.cc
}

// subcmd
func (c *cmd) BindSubCmds(r *cobrautils.Registory) *cmd {
	r.Bind(c.Cobra())
	return c
}

// func (c *cmd) RegisterTo(r *cobrautils.Registory) *cmd {
// 	r.RegisterSubCmd(c.Cobra())
// 	return c
// }
