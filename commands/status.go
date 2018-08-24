package commands

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type statusCmd struct {
	e RootExecutor
}

func newStatusCmd(e RootExecutor) *cobra.Command {
	c := &statusCmd{e: e}
	return newCommand(&cobra.Command{
		Use:   "status",
		Short: "Status of the MESG Core",
		RunE:  c.runE,
	})
}

func (c *statusCmd) runE(cmd *cobra.Command, args []string) error {
	// TODO: should improve this function with a waitFor
	status, err := c.e.Status()
	if err != nil {
		return err
	}

	if status == RUNNING {
		fmt.Println(aurora.Green("MESG Core is running"))
	} else {
		fmt.Println(aurora.Brown("MESG Core is stopped"))
	}
	return nil
}
