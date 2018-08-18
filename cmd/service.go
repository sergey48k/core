package cmd

import (
	"github.com/mesg-foundation/core/cmd/service"
	"github.com/spf13/cobra"
)

func newServiceCmd() *cobra.Command {
	return service.NewServiceCmd()
}
