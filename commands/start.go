package commands

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/commands/utils"
	"github.com/mesg-foundation/core/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type startCmd struct {
	lfv logFormatValue
	llv logLevelValue

	e RootExecutor
}

func newStartCmd(e RootExecutor) *cobra.Command {
	c := &startCmd{
		lfv: logFormatValue("text"),
		llv: logLevelValue("info"),
		e:   e,
	}
	cmd := newCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the MESG Core",
		RunE:  c.runE,
	})

	cmd.Flags().Var(&c.lfv, "log-format", "log format [text|json]")
	cmd.Flags().Var(&c.llv, "log-level", "log level [debug|info|warn|error|fatal|panic]")
	return cmd
}

func (c *startCmd) runE(cmd *cobra.Command, args []string) error {
	// TODO: figure out how to move this to config
	viper.Set(config.LogFormat, string(c.lfv))
	viper.Set(config.LogLevel, string(c.llv))

	status, err := c.e.Status()
	if err != nil {
		return err
	}
	if status == RUNNING {
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

// logFormatValue represents log format flag value.
type logFormatValue string

func (v *logFormatValue) Set(value string) error {
	if value != "text" && value != "json" {
		return fmt.Errorf("%s is not valid log format", value)
	}
	*v = logFormatValue(value)
	return nil
}
func (v *logFormatValue) Type() string   { return "string" }
func (v *logFormatValue) String() string { return string(*v) }

// logLevelValue represents log level flag value.
type logLevelValue string

func (v *logLevelValue) Set(value string) error {
	if _, err := logrus.ParseLevel(value); err != nil {
		return fmt.Errorf("%s is not valid log level", value)
	}
	*v = logLevelValue(value)
	return nil
}
func (v *logLevelValue) Type() string   { return "string" }
func (v *logLevelValue) String() string { return string(*v) }
