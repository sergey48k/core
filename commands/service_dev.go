package commands

import (
	"fmt"
	"os"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/commands/utils"
	"github.com/spf13/cobra"
)

type serviceDevCmd struct {
	baseCmd

	eventFilter  string
	taskFilter   string
	outputFilter string

	e ServiceExecutor
}

func newServiceDevCmd(e ServiceExecutor) *serviceDevCmd {
	c := &serviceDevCmd{
		e:           e,
		eventFilter: "*",
	}

	c.cmd = newCommand(&cobra.Command{
		Use:     "dev",
		Short:   "Run your service in development mode",
		Example: "mesg-core service dev PATH",
		Args:    cobra.MaximumNArgs(1),
		RunE:    c.runE,
	})
	c.cmd.Flags().StringVarP(&c.eventFilter, "event-filter", "e", "*", "Only log the data of the given event")
	c.cmd.Flags().StringVarP(&c.taskFilter, "task-filter", "t", "", "Only log the result of the given task")
	c.cmd.Flags().StringVarP(&c.outputFilter, "output-filter", "o", "", "Only log the data of the given output of a task result. If set, you also need to set the task in --task-filter")
	return c
}

func (c *serviceDevCmd) runE(cmd *cobra.Command, args []string) error {
	path := "./"
	if len(args) > 0 {
		path = args[0]
	}

	id, listeEvents, listeResults, err := c.e.ServiceDev(path, c.eventFilter, c.taskFilter, c.outputFilter)
	if err != nil {
		return err
	}
	defer c.e.ServiceDelete(id)

	fmt.Printf("%s Service started with success\n", aurora.Green("✔"))
	fmt.Printf("Service ID: %s\n", aurora.Bold(id))

	reader, err := c.e.ServiceLogs(id)
	if err != nil {
		return err
	}
	defer reader.Close()

	go stdcopy.StdCopy(os.Stdout, os.Stderr, reader)

	abort := utils.WaitForCancel()

loop:
	for {
		select {
		case e := <-listeEvents:
			fmt.Println("Receive event", aurora.Green(e.EventKey), ":", aurora.Bold(e.EventData))
		case r := <-listeResults:
			fmt.Println("Receive result", aurora.Green(r.TaskKey), aurora.Cyan(r.OutputKey), "with data", aurora.Bold(r.OutputData))
		case <-abort:
			break loop
		}
	}
	return nil
}
