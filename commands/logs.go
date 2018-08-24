package commands

import (
	"fmt"
	"os"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type logsCmd struct {
	e RootExecutor
}

func newLogsCmd(e RootExecutor) *cobra.Command {
	c := &logsCmd{e: e}
	return newCommand(&cobra.Command{
		Use:   "logs",
		Short: "Show the MESG Core's logs",
		RunE:  c.runE,
	})
}

func (c *logsCmd) runE(cmd *cobra.Command, args []string) error {
	status, err := c.e.Status()
	if err != nil {
		return err
	}

	if status == STOPPED {
		fmt.Println(aurora.Brown("MESG Core is stopped"))
		return nil
	}

	reader, err := c.e.Logs()
	if err != nil {
		return err
	}
	defer reader.Close()

	stdcopy.StdCopy(os.Stdout, os.Stderr, reader)
	return nil
}
