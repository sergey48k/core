package service

import (
	"github.com/spf13/cobra"
)

type ServiceDetail struct {
	Name   string
	Events []string
	Tasks  []string
}

type ServiceExecutor interface {
	DeleteAll() error
	Delete(ids ...string) error
	Deploy(path string) (id string, valid bool, err error)
	Detail(id string) (*ServiceDetail, error)
}

func Build(e ServiceExecutor) *cobra.Command {
	return newRootCmd(e)
}
