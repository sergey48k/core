package commands

import (
	"io"

	"github.com/krhubert/core/container"
	"github.com/spf13/cobra"
)

type RootExecutor interface {
	Start() error
	Stop() error
	Status() (container.StatusType, error)
	Logs() (io.ReadCloser, error)
}

type Executor interface {
	RootExecutor
	// 	TODO(uncoment): service.ServiceExecutor
}

func Build(e Executor) *cobra.Command {
	return newRootCmd(e)
}

type baseCmd struct {
	cmd *cobra.Command
}

// newCommand set default options for given command.
func newCommand(c *cobra.Command) *cobra.Command {
	c.DisableAutoGenTag = true
	return c
}
