package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"honnef.co/go/tools/version"
)

// Execute executes root command.
func Execute() error {
	err := newRootCmd().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, err)
	}
	return nil
}

func newRootCmd() *cobra.Cobra {
	cmd := &cobra.Command{
		Use:               "mesg-core",
		Short:             "Core",
		DisableAutoGenTag: true,
		Version:           version.Version,
	}

	cmd.AddCommand(
		newLogsCmd(),
		newStartCmd(),
		newServiceCmd(),
		newStatusCmd(),
		newStopCmd(),
	)
	return cmd
}
