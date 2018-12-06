package commands

import (
	"fmt"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/spf13/cobra"
	"strings"
)

var (
	uninstallReqs string
	yes           bool // FIXME var name
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall cx packages",
	Long:  `Long description...`,
	Run:   uninstall,
}

func init() {
	uninstallCmd.Flags().StringVarP(&uninstallReqs, "requirements", "r", "requirements.json", "Uninstall all the packages listed in the given requirements file.")
	uninstallCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Don't ask for confirmation of uninstall deletions.")

	cxpmsCmd.AddCommand(uninstallCmd)
}

func uninstall(cmd *cobra.Command, args []string) {
	// Uninstall packages from a requirements file
	if cmd.Flags().Changed("requirement") {
		reqs, err := models.LoadRequirements(uninstallReqs)
		if err != nil {
			fmt.Println("[!] Error: Unable to load requirements file!")
			return
		}
		if err := pms.UninstallRequirements(reqs, yes); err != nil {
			fmt.Println(err.Error())
			return
		}

		// Uninstall from ars
	} else {
		for _, arg := range args {
			if strings.Contains(arg, "/") {
				if pkg, ok := pms.GetPackageBySource(arg); ok {
					if err := pms.UninstallPkg(*pkg, yes); err != nil {
						fmt.Println("[!] Error: Package uninstall failed!")
					}
				}
			} else {
				// FIXME code repetition ^
				if pkg, ok := pms.GetPackageByName(arg); ok {
					if err := pms.UninstallPkg(*pkg, yes); err != nil {
						fmt.Println("[!] Error: Package uninstall failed!")
					}
				}
			}
		}
	}
}
