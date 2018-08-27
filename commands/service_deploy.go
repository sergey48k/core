package commands

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type serviceDeployCmd struct {
	baseCmd

	e ServiceExecutor
}

func newServiceDeployCmd(e ServiceExecutor) *serviceDeployCmd {
	c := &serviceDeployCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "deploy",
		Short: "Deploy a service",
		Long: `Deploy a service.

To get more information, see the [deploy page from the documentation](https://docs.mesg.com/guide/service/deploy-a-service.html)`,
		Example: `mesg-core service deploy PATH_TO_SERVICE`,
		RunE:    c.runE,
		Args:    cobra.MaximumNArgs(1),
	})
	return c
}

func (c *serviceDeployCmd) runE(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) > 1 {
		path = args[0]
	}
	id, valid, err := c.e.Deploy(path)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("service is not valid")
	}

	fmt.Println("Service deployed with ID:", aurora.Green(id))
	fmt.Printf("To start it, run the command: mesg-core service start %s\n", id)
	return nil
}
