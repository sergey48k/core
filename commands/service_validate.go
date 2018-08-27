package commands

import (
	"github.com/spf13/cobra"
)

type serviceValidateCmd struct {
	baseCmd

	e ServiceExecutor
}

func newServiceValidateCmd(e ServiceExecutor) *serviceValidateCmd {
	c := &serviceValidateCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "validate",
		Short: "Validate a service file",
		Long: `Validate a service file. Check the yml format and rules.

All the definitions of the service file can be found in the page [Service File from the documentation](https://docs.mesg.com/guide/service/service-file.html).`,
		Example: `mesg-core service validate
mesg-core service validate ./SERVICE_FOLDER`,
		Args: cobra.MaximumNArgs(1),
		RunE: c.runE,
	})
	return c
}

func (c *serviceValidateCmd) runE(cmd *cobra.Command, args []string) error {
	path := "./"
	if len(args) > 0 {
		path = args[0]
	}
	return c.e.ServiceValidate(path)
}
