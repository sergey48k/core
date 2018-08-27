package commands

import (
	"io"

	"github.com/krhubert/core/container"
	"github.com/mesg-foundation/core/service"
	"github.com/spf13/cobra"
)

type RootExecutor interface {
	Start() error
	Stop() error
	Status() (container.StatusType, error)
	Logs() (io.ReadCloser, error)
}

type ServiceExecutor interface {
	DeleteAll() error
	Delete(ids ...string) error
	Deploy(path string) (id string, valid bool, err error)
	Detail(id string) (*service.Service, error)
}

type Executor interface {
	RootExecutor
	ServiceExecutor
}

func Build(e Executor) *cobra.Command {
	return newRootCmd(e).cmd
}

type baseCmd struct {
	cmd *cobra.Command
}

// newCommand set default options for given command.
func newCommand(c *cobra.Command) *cobra.Command {
	c.DisableAutoGenTag = true
	return c
}
