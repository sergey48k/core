package service

import (
	sv "github.com/mesg-foundation/core/service"
	"github.com/spf13/cobra"
)

type ServiceExecutor interface {
	DeleteAll() error
	Delete(ids ...string) error
	Deploy(path string) (id string, valid bool, err error)
	Detail(id string) (*sv.Service, error)
}

func Build(e ServiceExecutor) *cobra.Command {
	return newRootCmd(e)
}
