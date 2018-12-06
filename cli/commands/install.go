package commands

import (
	"fmt"
	"github.com/galah4d/cx-pms/cli/messages"
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	requirement    string // FIXME move declaration
	no_deps        bool
	upgrade        bool
	target         string
	forceReinstall bool
)

var installCmd = &cobra.Command{
	Use:   messages.UsageInstall,
	Short: messages.Shortinstall,
	Long:  messages.LongInstall,
	Run:   install,
}

func init() {
	installCmd.Flags().StringVarP(&requirement, "requirement", "r", "requirements.json", "--requirement <file>")
	installCmd.Flags().BoolVar(&no_deps, "no-deps", false, "Don't install package dependencies")
	installCmd.Flags().StringVarP(&target, "target", "t", filepath.Join(os.Getenv("GOPATH"), config.InstallPATH), "Install packages into <dir>. By default this will not replace existing files/folders in <dir>.")
	installCmd.Flags().BoolVarP(&upgrade, "upgrade", "U", false, "Upgrade all specified packages to the newest available version. The handling of dependencies depends on the upgrade-strategy used.")
	installCmd.Flags().String("upgrade-strategy", "only-if-needed", "Determines how dependency upgrading should be handled [default: only-if-needed]. 'eager'' - dependencies are upgraded regardless of whether the currently installed version satisfies the requirements of the upgraded package(s). 'only-if-needed' - are upgraded only when they do not satisfy the requirements of the upgraded package(s).")
	installCmd.Flags().BoolVar(&forceReinstall, "force-reinstall", false, "Reinstall all packages even if they are already up-to-date.")
	installCmd.Flags().BoolP("ignore-installed", "I", false, "Ignore the installed packages (reinstalling instead).")

	cxpmsCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
	// Install from a requirements file
	if cmd.Flags().Changed("requirement") {
		reqs, err := models.LoadRequirements(requirement)
		if err != nil {
			fmt.Println("[!] Error: Unable to load requirements file!")
			return
		}
		if err := pms.InstallRequirements(reqs, target, forceReinstall); err != nil {
			fmt.Println(err.Error())
			return
		}

		// Install from args
	} else {
		for _, arg := range args {
			splitArgs := strings.Split(arg, "/")
			pkg := models.Package{Name: splitArgs[len(splitArgs)-1], Source: arg}

			if err := pms.InstallPkg(pkg, target, forceReinstall); err != nil {
				fmt.Println("[!] Error: Package installation failed!")
			}
		}
	}
}
