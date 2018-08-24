package commands

import (
	"github.com/spf13/cobra"
)

func newRootCmd(e Executor) *cobra.Command {
	cmd := newCommand(&cobra.Command{
		Use:   "mesg-core",
		Short: "MESG Core",
	})

	cmd.AddCommand(
		newStartCmd(e),
		newStatusCmd(e),
		newStopCmd(e),
		newLogsCmd(e),
		// TODO(uncomment): service.Build(e),
	)
	return cmd
}
