package commands

import (
	"io"

	"github.com/krhubert/core/container"
	"github.com/mesg-foundation/core/api/core"
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
	ServiceByID(id string) (*service.Service, error)
	ServiceDeleteAll() error
	ServiceDelete(ids ...string) error
	ServiceDeploy(path string) (id string, valid bool, err error)
	ServiceDetail(id string) (*service.Service, error)
	ServiceDev(path, eventFilter, taskFilter, outputFilter string) (id string, listeEvents chan core.EventData, listenResults chan core.ResultData, err error)
	ServiceLogs(id string) (io.ReadCloser, error)
	ServiceDependencyLogs(id string, dependency string) ([]io.ReadCloser, error)
	ServiceExeucteTask(id, taskKey, inputData string, tags []string) (listenResults chan core.ResultData, err error)

	ServiceStart(id string) error
	ServiceStop(id string) error
	ServiceValidate(path string) error
	ServiceGenerateDocs(path string) error
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
