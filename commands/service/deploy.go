package service

import (
	"errors"
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type deployCmd struct {
	e ServiceExecutor
}

func newDeployCmd(e ServiceExecutor) cobra.Command {
	c := &deployCmd{e: e}
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a service",
		Long: `Deploy a service.

To get more information, see the [deploy page from the documentation](https://docs.mesg.com/guide/service/deploy-a-service.html)`,
		Example:           `mesg-core service deploy PATH_TO_SERVICE`,
		RunE:              c.runE,
		Args:              cobra.MaximumNArgs(1),
		DisableAutoGenTag: true,
	}
}

func (c *deployCmd) runE(cmd *cobra.Command, args []string) error {
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

	fmt.Println("Service deployed with ID:", aurora.Green(serviceID))
	fmt.Printf("To start it, run the command: mesg-core service start %s\n", serviceID)
	return nil
}
