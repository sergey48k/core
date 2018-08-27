package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/api/core"
	"github.com/mesg-foundation/core/commands/utils"
	"github.com/sirupsen/logrus"
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
	c = &serviceDevCmd{
		e:           e,
		eventFilter: "*",
	}

	c.cmd = newCommand(&cobra.Command{
		Use:     "dev",
		Short:   "Run your service in development mode",
		Example: "mesg-core service dev PATH",
		RunE:    c.runE,
	})
	c.cmd.Flags().String("event-filter", "e", "*", "Only log the data of the given event")
	c.cmd.Flags().StringP("task-filter", "t", "", "Only log the result of the given task")
	c.cmd.Flags().StringP("output-filter", "o", "", "Only log the data of the given output of a task result. If set, you also need to set the task in --task-filter")
	return c
}

func (c *serviceDevCmd) runE(cmd *cobra.Command, args []string) error {
	serviceID, isValid, err := createService(defaultPath(args))
	if !isValid {
		os.Exit(1)
	}
	utils.HandleError(err)
	fmt.Printf("%s Service started with success\n", aurora.Green("✔"))
	fmt.Printf("Service ID: %s\n", aurora.Bold(serviceID))

	go listenEvents(serviceID, cmd.Flag("event-filter").Value.String())
	go listenResults(serviceID, cmd.Flag("task-filter").Value.String(), cmd.Flag("output-filter").Value.String())

	closeReaders := showLogs(serviceID, "*")
	defer closeReaders()

	<-utils.WaitForCancel()

	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Deleting test service..."}, func() {
		// This will automatically stop the service too
		cli().DeleteService(context.Background(), &core.DeleteServiceRequest{
			ServiceID: serviceID,
		})
	})
}

func deployService(p string) (string, bool, error) { return "", false, nil }

func createService(path string) (serviceID string, isValid bool, err error) {
	serviceID, isValid, err = deployService(path)
	if !isValid || err != nil {
		return "", isValid, err
	}

	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Starting service..."}, func() {
		_, err = cli().StartService(context.Background(), &core.StartServiceRequest{
			ServiceID: serviceID,
		})
	})
	return serviceID, true, err
}

func listenEvents(serviceID string, filter string) {
	stream, err := cli().ListenEvent(context.Background(), &core.ListenEventRequest{
		ServiceID:   serviceID,
		EventFilter: filter,
	})
	utils.HandleError(err)
	fmt.Println(aurora.Cyan("Listening for events from the service..."))
	for {
		event, err := stream.Recv()
		if err != nil {
			logrus.Info(aurora.Red(err))
			return
		}
		logrus.Info("Receive event", aurora.Green(event.EventKey), ":", aurora.Bold(event.EventData))
	}
}

func listenResults(serviceID string, result string, output string) {
	stream, err := cli().ListenResult(context.Background(), &core.ListenResultRequest{
		ServiceID:    serviceID,
		TaskFilter:   result,
		OutputFilter: output,
	})
	utils.HandleError(err)
	fmt.Println(aurora.Cyan("Listening for results from the service..."))
	for {
		result, err := stream.Recv()
		if err != nil {
			logrus.Info(aurora.Red(err))
			return
		}
		logrus.Info("Receive result", aurora.Green(result.TaskKey), aurora.Cyan(result.OutputKey), "with data", aurora.Bold(result.OutputData))
	}
}
