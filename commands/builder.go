package commands

import (
	"io"

	"github.com/spf13/cobra"
)

type StatusType uint

const (
	STOPPED StatusType = 0
	RUNNING StatusType = 1
)

type RootExecutor interface {
	Start() error
	Stop() error
	Status() (StatusType, error)
	Logs() (io.ReadCloser, error)
}

type Executor interface {
	RootExecutor
	// 	TODO(uncoment): service.ServiceExecutor
}

func Build(e Executor) *cobra.Command {
	return newRootCmd(e)
}

func newCommand(c *cobra.Command) *cobra.Command {
	c.DisableAutoGenTag = true
	return c
}
