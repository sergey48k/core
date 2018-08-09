package service

import (
	"context"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/mesg-foundation/core/api/core"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/service"
	"github.com/spf13/cobra"
)

type serviceStatus struct {
	service *service.Service
	status  service.StatusType
}

type byStatus []serviceStatus

func (a byStatus) Len() int           { return len(a) }
func (a byStatus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStatus) Less(i, j int) bool { return a[j].status < a[i].status }

// List all the services
var List = &cobra.Command{
	Use:   "list",
	Short: "List all published services",
	Long: `This command returns all published services with basic information.
Optionally, you can filter the services published by a specific developer:
To have more details, see the [detail command](mesg-core_service_detail.md).`,
	Example:           `mesg-core service list`,
	Run:               listHandler,
	DisableAutoGenTag: true,
}

func listHandler(cmd *cobra.Command, args []string) {
	reply, err := cli().ListServices(context.Background(), &core.ListServicesRequest{})
	utils.HandleError(err)
	status, err := servicesWithStatus(reply.Services)
	utils.HandleError(err)
	sort.Sort(byStatus(status))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "STATUS\tSERVICE\tNAME\n")
	for _, s := range status {
		fmt.Fprintf(w, "%s\t%s\t%s\n", s.status.String(), s.service.Hash(), s.service.Name)
	}
	w.Flush()
}

func servicesWithStatus(services []*service.Service) (status []serviceStatus, err error) {
	for _, s := range services {
		st, err := s.Status()
		if err != nil {
			break
		}
		status = append(status, serviceStatus{
			service: s,
			status:  st,
		})
	}
	return
}
