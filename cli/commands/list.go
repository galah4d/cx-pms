package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list installed cx packages",
	Long: `List installed cx packages.

Packages are listed in a case-insensitive sorted order.`,
	Args: cobra.NoArgs,
	Run:  list,
}

func init() {
	cxpmsCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	fmt.Println("Installed Packages:")
	for _, installation := range pms.GetInstallations("by-name") {
		// TODO add dynamic padding values
		fmt.Printf("| %-10s | %-30s | %-40s | %-30s |\n", installation.Package.Name, installation.Package.Source, installation.Date, installation.Path)
	}
}
