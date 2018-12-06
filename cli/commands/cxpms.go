package commands

import (
	"fmt"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/spf13/cobra"
	"os"
)

var pms models.PackageManagementSystem

var verbose bool

var cxpmsCmd = &cobra.Command{
	Use:   "cxpms",
	Short: "CX-PMS is a package management tool for CX",
	Long:  `CX-PMS is a package management tool for CX long version...`,
	Run: func(cmd *cobra.Command, args []string) {
		// print help and exit
		if err := cmd.Help(); err != nil {
			os.Exit(0)
		}
	},
}

func init() {
	cxpmsCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	err := pms.InitPMS()
	if err != nil {
		fmt.Println("Error: Failed to initialize PMS...")
		os.Exit(0)
	}
}

func Execute() {
	if err := cxpmsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
