package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/mesg-foundation/core/api/core"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/config"
	"github.com/mesg-foundation/core/daemon"
	"github.com/mesg-foundation/core/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Stop the Core
var Stop = &cobra.Command{
	Use:               "stop",
	Short:             "Stop the Core",
	Run:               stopHandler,
	DisableAutoGenTag: true,
}

func init() {
	RootCmd.AddCommand(Stop)
}

func stopHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Stopping Core...")
	err := stopServices()
	utils.HandleError(err)
	err = daemon.Stop()
	utils.HandleError(err)
	fmt.Println("Core stopped")
}

func getCli() (cli core.CoreClient, err error) {
	connection, err := grpc.Dial(viper.GetString(config.APIClientTarget), grpc.WithInsecure())
	if err != nil {
		return
	}
	cli = core.NewCoreClient(connection)
	return
}

func stopServices() (err error) {
	cli, err := getCli()
	if err != nil {
		return err
	}
	hashes, err := service.ListRunning()
	if err != nil {
		return err
	}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, hash := range hashes {
		wg.Add(1)
		go func(serviceID string) {
			defer wg.Done()
			_, errStop := cli.StopService(context.Background(), &core.StopServiceRequest{
				ServiceID: serviceID,
			})
			mutex.Lock()
			defer mutex.Unlock()
			if errStop != nil && err == nil {
				err = errStop
			}
		}(hash)
	}
	wg.Wait()
	return
}
