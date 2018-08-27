package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type serviceDeleteCmd struct {
	baseCmd

	all   bool
	force bool

	e ServiceExecutor
}

func newServiceDeleteCmd(e ServiceExecutor) *serviceDeleteCmd {
	c := &serviceDeleteCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete one or many services",
		Example: `mesg-core service delete SERVICE_ID [SERVICE_ID...]
mesg-core service delete --all`,
		RunE: c.runE,
	})

	c.cmd.Flags().BoolVar(&c.all, "all", false, "Delete all services")
	c.cmd.Flags().BoolVarP(&c.force, "force", "f", false, "Force delete all services")
	return c
}

func (c *serviceDeleteCmd) runE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 && !c.all {
		return errors.New("at least one service id must be provided (or run with --all flag)")
	}

	if c.all && !c.force {
		if survey.AskOne(&survey.Confirm{Message: "Are you sure to delete all services?"}, &c.force, nil) != nil {
			return nil
		}
		// is still no confirm return.
		if !c.force {
			return nil
		}
	}

	if c.all {
		if err := c.e.ServiceDeleteAll(); err != nil {
			return err
		}
		fmt.Println("All services are deleted")
		return nil
	}

	for _, arg := range args {
		if err := c.e.ServiceDelete(arg); err != nil {
			fmt.Fprintln(os.Stderr, "Can't delete %s service: %s", arg, err)
		} else {
			fmt.Println("Service", arg, "deleted")
		}
	}
	return nil
}
