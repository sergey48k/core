package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type serviceDocsCmd struct {
	baseCmd

	force bool

	e ServiceExecutor
}

func newServiceDocsCmd(e ServiceExecutor) *serviceDocsCmd {
	c := &serviceDocsCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:   "gen-doc",
		Short: "Generate the documentation for the service in a README.md file",
		Example: `mesg-core service gen-doc
mesg-core service gen-doc ./PATH_TO_SERVICE`,
		RunE: c.runE,
	})
	return c
}

func (c *serviceDocsCmd) runE(cmd *cobra.Command, args []string) error {
	path := "./"
	if len(args) > 0 {
		path = args[0]
	}

	readmePath := filepath.Join(path, "README.md")
	if _, err := os.Stat(readmePath); !c.force && err == nil {
		if survey.AskOne(&survey.Confirm{Message: "The file README.md already exists. Do you want to overwrite it?"}, &c.force, nil) != nil {
			return nil
		}
		if !c.force {
			return nil
		}
	}

	if err := c.e.ServiceGenerateDocs(path); err != nil {
		return err
	}

	fmt.Println(aurora.Green("File README.md generated with success"))
	return nil
}
