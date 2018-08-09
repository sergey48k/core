package cmd

import (
	"fmt"
	"os"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

// Logs of the core.
var Logs = &cobra.Command{
	Use:               "logs",
	Short:             "Show the Core's logs",
	Run:               logsHandler,
	DisableAutoGenTag: true,
}

func init() {
	RootCmd.AddCommand(Logs)
}

func logsHandler(cmd *cobra.Command, args []string) {
	status, err := daemon.Status()
	utils.HandleError(err)
	if status == container.STOPPED {
		fmt.Println("Core is stopped")
		return
	}
	reader, err := daemon.Logs()
	defer reader.Close()
	utils.HandleError(err)
	stdcopy.StdCopy(os.Stdout, os.Stderr, reader)
}
