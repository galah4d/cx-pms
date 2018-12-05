package commands

import (
	"fmt"
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/spf13/cobra"
	"strings"
)

var (
	requirement string // FIXME move declaration
	no_deps bool
	upgrade bool
	target string
	force_reinstall bool
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs CX packages",
	Long:  `Install packages from:

    VCS project urls.
    Local project directories.
    Local or remote source archives.

cx-pms also supports installing from “requirements files”, which provide an easy way to specify a whole environment to be installed`,
	Run: install,
}

func init() {
	installCmd.Flags().StringVarP(&requirement, "requirement", "r", "requirements.json", "--requirement <file>")
	installCmd.Flags().BoolVar(&no_deps, "no-deps", false, "Don't install package dependencies")
	installCmd.Flags().StringVarP(&target, "target", "t", config.InstallPATH, "Install packages into <dir>. By default this will not replace existing files/folders in <dir>.")
	installCmd.Flags().BoolVarP(&upgrade, "upgrade", "U", false, "Upgrade all specified packages to the newest available version. The handling of dependencies depends on the upgrade-strategy used.")
	installCmd.Flags().String("upgrade-strategy", "only-if-needed", "Determines how dependency upgrading should be handled [default: only-if-needed]. “eager” - dependencies are upgraded regardless of whether the currently installed version satisfies the requirements of the upgraded package(s). “only-if-needed” - are upgraded only when they do not satisfy the requirements of the upgraded package(s).")
	installCmd.Flags().BoolVar(&force_reinstall, "force-reinstall", false, "Reinstall all packages even if they are already up-to-date.")
	installCmd.Flags().BoolP("ignore-installed", "I", false, "Ignore the installed packages (reinstalling instead).")

	cxpmsCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
	var installer models.Installer
	if err := installer.UnmarshalJSON(config.InstallationFilePATH); err != nil {
		fmt.Println("[!] Error: Unable to initialize installer")
		return
	}

	// Install a requirements file
	if cmd.Flags().Changed("requirement") {
		reqs, err := models.LoadRequirements(cmd.Flag("requirement").Value.String())
		if err != nil {
			fmt.Println("[!] Error: Unable to load requirements file!")
			return
		}
		for _, pkg := range reqs.Packages {
			if !installer.Installed(pkg) {
				if err := installer.Install(pkg, target); err != nil {
					fmt.Println(err)
				}
			} else {
				if force_reinstall {
					if err := installer.Reinstall(pkg, target); err != nil {
						fmt.Println(err)
					}
				} else if upgrade {
					// TODO
				}
			}
		}

	// Install from args
	} else {
		for _, arg := range args {
			splitArgs := strings.Split(arg, "/")
			pkg := models.Package{Name: splitArgs[len(splitArgs)-1], Source:arg}

			if err := installer.Install(pkg, target); err != nil {
				fmt.Println("[!] Error: Package installation failed!")
			}
		}
	}
}
