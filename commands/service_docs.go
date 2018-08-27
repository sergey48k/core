package commands

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/commands/assets"
	"github.com/mesg-foundation/core/service/importer"
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

	// TODO : think if this shold be handle by ServiceExecutor interface

	readmePath := filepath.Join(path, "README.md")
	if _, err := os.Stat(readmePath); !c.force && err == nil {
		if survey.AskOne(&survey.Confirm{Message: "The file README.md already exists. Do you want to overwrite it?"}, &c.force, nil) != nil {
			return nil
		}
		if !c.force {
			return nil
		}
	}

	service, err := importer.From(path)
	if err != nil {
		if _, ok := err.(*importer.ValidationError); ok {
			fmt.Println("Run the command 'service validate' for more details")
		}
		return err
	}

	readmeTemplate, err := assets.Asset("cmd/service/assets/readmeTemplate.md")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(readmePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("doc").Parse(string(readmeTemplate)))
	if err := tmpl.Execute(f, service); err != nil {
		return err
	}

	fmt.Println(aurora.Green("File README.md generated with success"))
	return nil
}
