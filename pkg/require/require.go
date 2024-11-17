package require

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%q accepts no arguments\n\nUsage:  %s", cmd.CommandPath(), cmd.UseLine())
	}
	return nil
}

func NoMoreArgsCompFunc(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return noMoreArgsComp()
}

func noMoreArgsComp() ([]string, cobra.ShellCompDirective) {
	activeHelpMessage := "This command does not take any more arguments (but may accept flags)"
	return cobra.AppendActiveHelp(nil, activeHelpMessage), cobra.ShellCompDirectiveNoFileComp
}
