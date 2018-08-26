package commands

import (
	"fmt"

	"github.com/krhubert/core/container"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type statusCmd struct {
	baseCmd

	e RootExecutor
}

func newStatusCmd(e RootExecutor) *statusCmd {
	c := &statusCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "status",
		Short: "Status of the MESG Core",
		RunE:  c.runE,
	})
	return c
}

func (c *statusCmd) runE(cmd *cobra.Command, args []string) error {
	// TODO: should improve this function with a waitFor
	status, err := c.e.Status()
	if err != nil {
		return err
	}

	if status == container.RUNNING {
		fmt.Println(aurora.Green("MESG Core is running"))
	} else {
		fmt.Println(aurora.Brown("MESG Core is stopped"))
	}
	return nil
}
