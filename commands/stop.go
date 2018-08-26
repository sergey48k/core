package commands

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/commands/utils"
	"github.com/spf13/cobra"
)

type stopCmd struct {
	baseCmd
	e RootExecutor
}

func newStopCmd(e RootExecutor) *stopCmd {
	c := &stopCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stop the MESG Core",
		RunE:  c.runE,
	})
	return c
}

func (c *stopCmd) runE(cmd *cobra.Command, args []string) error {
	var err error
	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Stopping MESG Core..."}, func() {
		err = c.e.Stop()
	})

	if err != nil {
		return err
	}
	fmt.Println(aurora.Green("MESG Core stopped"))
	return nil
}
