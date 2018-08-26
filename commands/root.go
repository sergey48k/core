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
		newStartCmd(e).cmd,
		newStatusCmd(e).cmd,
		newStopCmd(e).cmd,
		newLogsCmd(e).cmd,
		// TODO(uncomment): service.Build(e),
	)
	return cmd
}
