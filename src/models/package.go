package models

import (
	"github.com/galah4d/cx-pms/src/utils"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
	"strings"
)

type Package struct {
	Name string `json:"name"`
	Source string `json:"source"`
	Requirements Requirements `json:"requirements,omitempty"`
}

func (p *Package) install (target string) error {
	fmt.Printf("[+] Installing '%s' from %s... to %s\n", p.Name, p.Source, target)
	// Initializes the installation directory
	if err := utils.DirInit(filepath.Join(target, p.Source)); err != nil {
		return err // Unable to initialize directory
	}

	_, err := git.PlainClone(filepath.Join(target, p.Source), false, &git.CloneOptions{
		URL: "http://" + p.Source,
		Progress: nil,//os.Stdout,
	})
	if err != nil {
		return err // Failed to install package
	}
	return nil
}

func (p Package) IsInstalled () bool {
	return utils.DirExists(filepath.Join("./pkg", p.Source))
}

func (p *Package) hasDependencies () bool {
	return len(p.Requirements.Packages) > 0
}

func (p *Package) installRequirements (target string) {
	for _, r := range p.Requirements.Packages {
		err := r.install(target)
		if err != nil {
			fmt.Printf("[!] Unable to fulfil requirement '%s'!", r.Name)
		}
	}
}

func (pkg *Package) uninstall (p string) error {
	fmt.Printf("[+] Uninstalling '%s'\n", pkg.Name)
	return os.RemoveAll(filepath.Join(p + pkg.Source))
}

func (pkg *Package) reinstall () error {
	/*
	if err := pkg.uninstall(); err != nil {
		return err
	}
	return pkg.install()
	/*/
	return nil
}

func (p Package) upgrade () error {
	return nil
}

func (pkg *Package) GetFiles (installationPath string, suffix string) ([]string, error) {
	var files []string
	root := filepath.Join(installationPath, pkg.Source)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if utils.FileExists(path) && strings.HasSuffix(path, suffix) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}