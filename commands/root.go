package commands

import (
	"github.com/spf13/cobra"
)

type rootCmd struct {
	baseCmd
}

func newRootCmd(e Executor) *rootCmd {
	c := &rootCmd{}
	c.cmd = newCommand(&cobra.Command{
		Use:   "mesg-core",
		Short: "MESG Core",
	})

	c.cmd.AddCommand(
		newStartCmd(e).cmd,
		newStatusCmd(e).cmd,
		newStopCmd(e).cmd,
		newLogsCmd(e).cmd,
		newRootServiceCmd(e).cmd,
	)
	return c
}
