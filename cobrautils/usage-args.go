package cobrautils

type labelOptArg func(*HelpLabels)

func Usage(usage string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetUsage(usage)
	}
}
func Aliases(aliases string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetAliases(aliases)
	}
}
func Examples(examples string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetExamples(examples)
	}
}
func AvailableCommands(availableCommands string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetAvailableCommands(availableCommands)
	}
}
func AdditionalCommands(additionalCommands string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetAdditionalCommands(additionalCommands)
	}
}
func Flags(flags string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetFlags(flags)
	}
}
func GlobalFlags(globalFlags string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetGlobalFlags(globalFlags)
	}
}
func AddtionalHelpTopics(addtionalHelpTopics string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetAddtionalHelpTopics(addtionalHelpTopics)
	}
}
func UseHelp(useHelp string) labelOptArg {
	return func(l *HelpLabels) {
		l.SetUseHelp(useHelp)
	}
}

func NewLabel(opts ...labelOptArg) *HelpLabels {
	l := DefaultHelpLabels()
	for _, opt := range opts {
		opt(l)
	}
	return l
}
