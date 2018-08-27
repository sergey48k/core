package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type serviceListCmd struct {
	baseCmd

	e ServiceExecutor
}

func newServiceListCmd(e ServiceExecutor) *serviceListCmd {
	c := &serviceListCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "list",
		Short: "List all published services",
		Long: `This command returns all published services with basic information.
Optionally, you can filter the services published by a specific developer:
To have more details, see the [detail command](mesg-core_service_detail.md).`,
		Example: `mesg-core service list`,
		Args:    cobra.NoArgs,
		RunE:    c.runE,
	})
	return c
}

func (c *serviceListCmd) runE(cmd *cobra.Command, args []string) error {
	services, err := c.e.ServiceListWitsStatus()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "STATUS\tSERVICE\tNAME\n")
	for _, s := range services {
		fmt.Fprintf(w, "%s\t%s\t%s\n", "unknown" /* s.Status() */, s.Hash(), s.Name)
	}
	return w.Flush()
}
