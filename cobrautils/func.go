package cobrautils

import "github.com/spf13/cobra"

// PreRun

// func WithParentPreRun(f func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
// 	return func(cmd *cobra.Command, args []string) {
// 		parent := cmd.Parent()
// 		if parent != nil {
// 			if parent.PreRun != nil {
// 				parent.PreRun(parent, args)
// 			}
// 		}

// 		f(cmd, args)
// 	}
// }

// func WithParentPreRunE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
// 	return func(cmd *cobra.Command, args []string) error {
// 		parent := cmd.Parent()
// 		if parent != nil {
// 			if parent.PreRunE != nil {
// 				err := parent.PreRunE(parent, args)
// 				if err != nil {
// 					return err
// 				}
// 			} else if parent.PreRun != nil {
// 				parent.PreRun(parent, args)
// 			}
// 		}

// 		return f(cmd, args)
// 	}
// }

func WithParentPersistentPreRun(f func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		parent := cmd.Parent()
		if parent != nil {
			if parent.PersistentPreRun != nil {
				parent.PersistentPreRun(parent, args)
			}
		}

		f(cmd, args)
	}
}

func WithParentPersistentPreRunE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		parent := cmd.Parent()
		if parent != nil {
			if parent.PersistentPreRunE != nil {
				err := parent.PersistentPreRunE(parent, args)
				if err != nil {
					return err
				}
			} else if parent.PersistentPreRun != nil {
				parent.PersistentPreRun(parent, args)
			}
		}

		return f(cmd, args)
	}
}

// PostRun

// func WithParentPostRun(f func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
// 	return func(cmd *cobra.Command, args []string) {
// 		f(cmd, args)

// 		parent := cmd.Parent()
// 		if parent != nil {
// 			if parent.PostRun != nil {
// 				parent.PostRun(parent, args)
// 			}
// 		}
// 	}
// }

// func WithParentPostRunE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
// 	return func(cmd *cobra.Command, args []string) error {
// 		err := f(cmd, args)
// 		if err != nil {
// 			return err
// 		}

// 		parent := cmd.Parent()
// 		if parent != nil {
// 			if parent.PostRunE != nil {
// 				err = parent.PostRunE(parent, args)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}

// 		return nil
// 	}
// }

func WithParentPersistentPostRun(f func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		f(cmd, args)

		parent := cmd.Parent()
		if parent != nil {
			if parent.PersistentPostRun != nil {
				parent.PersistentPostRun(parent, args)
			}
		}
	}
}

func WithParentPersistentPostRunE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := f(cmd, args)
		if err != nil {
			return err
		}

		parent := cmd.Parent()
		if parent != nil {
			if parent.PersistentPostRunE != nil {
				err = parent.PersistentPostRunE(parent, args)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}
}
