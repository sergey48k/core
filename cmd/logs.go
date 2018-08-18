package cmd

import (
	"fmt"
	"os"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

type logsCmd struct{}

func newLogsCmd() *cobra.Command {
	c := &logsCmd{}

	cmd := &cobra.Command{
		Use:               "logs",
		Short:             "Show the Core's logs",
		RunE:              c.runE,
		DisableAutoGenTag: true,
	}

	return cmd
}

func (*logsCmd) runE(cmd *cobra.Command, args []string) error {
	status, err := daemon.Status()
	if err != nil {
		return err
	}

	if status == container.STOPPED {
		fmt.Println("Core is stopped")
		return nil
	}

	reader, err := daemon.Logs()
	if err != nil {
		return err
	}
	defer reader.Close()

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, reader)
	return err
}
