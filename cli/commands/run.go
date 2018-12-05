package commands

import (
	"fmt"
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run short",
	Long:  `run long`,
	Run:   run,
}

func init() {
	runCmd.Flags().StringVarP(&requirement, "requirement", "r", "requirements.json", "Include all packages listed in the given requirements file.")

	cxpmsCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	var installer models.Installer
	if err := installer.UnmarshalJSON(filepath.Join(os.Getenv("GOPATH"), config.InstallationFilePATH)); err != nil {
		fmt.Println("[!] Error: Unable to initialize installer")
		return
	}

	var cxFiles []string
	// Reads requirements if -r flag is set
	if cmd.Flags().Changed("requirement") {
		reqs, err := models.LoadRequirements(cmd.Flag("requirement").Value.String())
		if err != nil {
			fmt.Println("[!] Error: Unable to load requirements file!")
			return
		}

		for _, pkg := range reqs.Packages {
			path, err := installer.GetInstallationPath(pkg)
			if err != nil {
				// Todo package not installed, force-install flag
			} else {
				fs, err := pkg.GetFiles(path, ".cx")
				if err != nil {
					fmt.Println(err.Error())
					continue // TODO handle error
				}
				for _, cxFile := range fs {
					cxFiles = append(cxFiles, cxFile)
				}
			}
		}
	}

	args = append(cxFiles, args...)
	command := exec.Command("cx", args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
