package cmd

import (
	"fmt"

	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

// Status command returns started services
var Status = &cobra.Command{
	Use:               "status",
	Short:             "Status of the Core",
	Run:               statusHandler,
	DisableAutoGenTag: true,
}

func init() {
	RootCmd.AddCommand(Status)
}

func statusHandler(cmd *cobra.Command, args []string) {
	// TODO: should improve this function with a waitFor
	status, err := daemon.Status()
	utils.HandleError(err)
	if status == container.RUNNING {
		fmt.Println("Core is running")
	} else {
		fmt.Println("Core is stopped")
	}
}
