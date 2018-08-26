package commands

import (
	"fmt"

	"github.com/krhubert/core/container"
	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/commands/utils"
	"github.com/mesg-foundation/core/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type startCmd struct {
	baseCmd

	lfv logFormatValue
	llv logLevelValue

	e RootExecutor
}

func newStartCmd(e RootExecutor) *startCmd {
	c := &startCmd{
		lfv: logFormatValue("text"),
		llv: logLevelValue("info"),
		e:   e,
	}
	c.cmd = newCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the MESG Core",
		RunE:  c.runE,
	})

	c.cmd.Flags().Var(&c.lfv, "log-format", "log format [text|json]")
	c.cmd.Flags().Var(&c.llv, "log-level", "log level [debug|info|warn|error|fatal|panic]")
	return c
}

func (c *startCmd) runE(cmd *cobra.Command, args []string) error {
	// TODO: figure out how to move this to config
	viper.Set(config.LogFormat, string(c.lfv))
	viper.Set(config.LogLevel, string(c.llv))

	status, err := c.e.Status()
	if err != nil {
		return err
	}
	if status == container.RUNNING {
		fmt.Println(aurora.Green("MESG Core is running"))
		return nil
	}
	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Starting MESG Core..."}, func() {
		err = c.e.Start()
	})
	if err != nil {
		return err
	}

	fmt.Println(aurora.Green("MESG Core is running"))
	return nil
}
