package cmd

import (
	"fmt"

	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

type startCmd struct{}

func newstartCmd() *cobra.Command {
	c := &startCmd{}

	cmd := &cobra.Command{
		Use:               "start",
		Short:             "Start the core",
		RunE:              c.runE,
		DisableAutoGenTag: true,
	}
	return cmd
}

func (*startCmd) runE(cmd *cobra.Command, args []string) error {
	status, err := daemon.Status()
	if err != nil {
		return err
	}

	if status == container.RUNNING {
		fmt.Println("Core is running")
		return nil
	}

	fmt.Println("Starting Core...")
	_, err = daemon.Start()
	if err != nil {
		return err
	}

	fmt.Println("Core is running")
	return nil
}
