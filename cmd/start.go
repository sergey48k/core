package cmd

import (
	"fmt"

	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

// Start the Core.
var Start = &cobra.Command{
	Use:               "start",
	Short:             "Start the Core",
	Run:               startHandler,
	DisableAutoGenTag: true,
}

func init() {
	RootCmd.AddCommand(Start)
}

func startHandler(cmd *cobra.Command, args []string) {
	status, err := daemon.Status()
	utils.HandleError(err)
	if status == container.RUNNING {
		fmt.Println("Core is running")
		return
	}
	fmt.Println("Starting Core...")
	_, err = daemon.Start()
	utils.HandleError(err)
	fmt.Println("Core is running")
}
