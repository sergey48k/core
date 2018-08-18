package cmd

import (
	"fmt"

	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

type statusCmd struct{}

func newstatusCmd() *cobra.Command {
	c := &statusCmd{}

	cmd := &cobra.Command{
		Use:               "status",
		Short:             "Print status of the Core",
		RunE:              c.runE,
		DisableAutoGenTag: true,
	}
	return cmd
}

func (*statusCmd) runE(cmd *cobra.Command, args []string) error {
	status, err := daemon.Status()
	if err != nil {
		return err
	}

	switch status {
	case container.RUNNING:
		fmt.Println("Core status: running")
	case container.STOPPED:
		fmt.Println("Core status: stopped")
	default:
		fmt.Println("Core status:", status)
	}
	return nil
}
