package commands

import (
	"github.com/spf13/cobra"
)

type rootServiceCmd struct {
	baseCmd
}

func newRootServiceCmd(e ServiceExecutor) *rootServiceCmd {
	c := &rootServiceCmd{}
	c.cmd = newCommand(&cobra.Command{
		Use:   "service",
		Short: "Manage your services",
	})

	c.cmd.AddCommand(
	// newDeployCmd(e),
	// newValidateCmd(e),
	// newTestCmd(e),
	// newStartCmd(e),
	// newStopCmd(e),
	// newDetailCmd(e),
	// newListCmd(e),
	// newInitCmd(e),
	// newDeleteCmd(e),
	// newLogsCmd(e),
	// newDocsCmd(e),
	// newDevCmd(e),
	// newExecuteCmd(e),
	)
	return c
}
