package commands

import (
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/models"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list installed cx packages",
	Long:  `List installed cx packages.

Packages are listed in a case-insensitive sorted order.`,
	Run: list,
}

func init() {
	cxpmsCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	var installer models.Installer
	if err := installer.UnmarshalJSON(config.InstallationFilePATH); err != nil {
		fmt.Println("[!] Error: Unable to initialize installer")
		return
	}

	fmt.Println("Installed Packages:")
	for _, installation := range installer.Installations { // TODO implement sorting
		fmt.Printf("%s\t| %s\t| %s\t| %s\t\n", installation.Package.Name, installation.Package.Source, installation.Date, installation.Path)
	}
}
