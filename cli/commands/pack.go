package commands

// TODO clean this mess
/*import "github.com/spf13/cobra"


import (
	"bufio"
	"fmt"
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/models"
	"github.com/galah4d/cx-pms/src/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "generates a requirements file",
	Long: `generates a requirements file`,
	Run: pack,
}

func init() {
	cxpmsCmd.AddCommand(packCmd)
}

func pack(cmd *cobra.Command, args []string) {
	var installer models.Installer
	if err := installer.UnmarshalJSON(filepath.Join(os.Getenv("GOPATH"), config.InstallationFilePATH)); err != nil {
		fmt.Println("[!] Error: Unable to initialize installer")
		return
	}

	pkg_re := regexp.MustCompile(`package [a-zA-Z0-9_]+$`)
	import_re := regexp.MustCompile(`import "[a-zA-Z0-9_]+"$`) // FIXME
	ar := regexp.MustCompile("[^a-zA-Z0_9_]+")

	var packages []string
	var imports []string

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if pkg_re.Match([]byte(scanner.Text())) {
				packages = append(packages, strings.Split(scanner.Text(), " ")[1])
			} else if import_re.Match([]byte(scanner.Text())) {
				imports = append(imports, ar.ReplaceAllString(strings.Split(scanner.Text(), " ")[1], ""))
			}
		}
	}

	var requirements models.Requirements
	for _, r := range utils.Difference(imports, packages) {
		fmt.Println(r)
		if i, err := installer.GetInstallation(r); err == nil {
			requirements.AddRequirement(i.Package)
			fmt.Println(i.Package.Name)
		}
	}

	requirements.Save("r.json")
}*/
