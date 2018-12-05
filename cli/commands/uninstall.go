package commands

import (
	"github.com/galah4d/cx-pms/cli/messages"
	"github.com/galah4d/cx-pms/src/models"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	uninstallReq string
	yes bool // FIXME var name
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall cx packages",
	Long:  `Long description...`,
	Run: uninstall,
}

func init() {
	uninstallCmd.Flags().StringVarP(&uninstallReq, "requirement", "r", "requirements.json", "Uninstall all the packages listed in the given requirements file.")
	uninstallCmd.Flags().BoolVarP(&yes, "yes", "y",  false, "Don't ask for confirmation of uninstall deletions.")

	cxpmsCmd.AddCommand(uninstallCmd)
}

func uninstall(cmd *cobra.Command, args []string) {
	var installer models.Installer
	installer.UnmarshalJSON("installation.json") // TODO error handling

	// Uninstall packages from a requirements file
	if cmd.Flags().Changed("requirement") {
		reqs, err := models.LoadRequirements(cmd.Flag("requirement").Value.String())
		if err != nil {
			fmt.Println("[!] Error: Unable to load requirements file!")
			return
		}
		for _, pkg := range reqs.Packages {
			if installer.Installed(pkg) && (yes || messages.AskForConfirmation("Proceed (y/n)?")) {
				if err := installer.Uninstall(pkg); err != nil {
					fmt.Println(err)
				}
			}
		}
	} else {
		var pkg models.Package
		for _, arg := range args {
			if strings.Contains(arg, "/") {
				splitArgs := strings.Split(arg, "/")
				pkg = models.Package{Name: splitArgs[len(splitArgs)-1], Source:arg}

			} else {
				// TODO uninstall from pkg name
			}

			if installer.Installed(pkg) {
				if yes || messages.AskForConfirmation("Proceed (y/n)?") {
					if err := installer.Uninstall(pkg); err != nil {
						fmt.Println("[!] Error: Package uninstall failed!")
					}
				}
			} else {
				fmt.Printf("[!] Error: Unable to locate package %s!\n", pkg.Name)
			}
		}
	}
}