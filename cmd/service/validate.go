package service

import (
	"fmt"

	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/service/importer"
	"github.com/spf13/cobra"
)

// Validate a service
var Validate = &cobra.Command{
	Use:   "validate",
	Short: "Validate a service file",
	Long: `Validate a service file. Check the yml format and rules.

All the definitions of the service file can be found in the page [Service File from the documentation](https://docs.mesg.com/guide/service/service-file.html).`,
	Example: `mesg-core service validate
mesg-core service validate ./SERVICE_FOLDER`,
	Run:               validateHandler,
	DisableAutoGenTag: true,
}

func validateHandler(cmd *cobra.Command, args []string) {
	validation, err := importer.Validate(defaultPath(args))
	utils.HandleError(err)

	validateServiceFile(validation)
	validateDockerfile(validation)
	validateSummary(validation)
}

func validateServiceFile(validation *importer.ValidationResult) {
	if validation.ServiceFileExist == false {
		fmt.Println("⨯ File 'mesg.yml' does not exist")
		return
	}

	if len(validation.ServiceFileWarnings) > 0 {
		fmt.Printf("%s File 'mesg.yml' is not valid. See documentation: %s\n", "⨯", "https://docs.mesg.com/guide/service/service-file.html")
		for _, warning := range validation.ServiceFileWarnings {
			fmt.Printf("  - %s\n", warning)
		}
	} else {
		fmt.Println("✔ File 'mesg.yml' is valid")
	}
}

func validateDockerfile(validation *importer.ValidationResult) {
	if validation.DockerfileExist {
		fmt.Println("✔ Dockerfile exists")
	} else {
		fmt.Println("⨯ Dockerfile does not exist")
	}
}

func validateSummary(validation *importer.ValidationResult) {
	if validation.IsValid() {
		fmt.Println("Service is valid")
	} else {
		fmt.Println("Service is not valid")
	}
}
