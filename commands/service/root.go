package service

import (
	"github.com/spf13/cobra"
)

func newRootCmd(e ServiceExecutor) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Manage your services",
	}

	cmd.AddCommand(
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
	return cmd
}
