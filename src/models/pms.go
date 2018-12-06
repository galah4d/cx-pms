package models

import (
	"fmt"
	"github.com/galah4d/cx-pms/cli/messages"
	"github.com/galah4d/cx-pms/config"
	"os"
	"path/filepath"
)

type PackageManagementSystem struct {
	Installer Installer
}

func (pms *PackageManagementSystem) InitPMS() error {
	return pms.Installer.UnmarshalJSON(filepath.Join(os.Getenv("GOPATH"), config.InstallationFilePATH))
}

func (pms *PackageManagementSystem) GetPackageByName(n string) (*Package, bool) {
	for _, i := range pms.Installer.Installations {
		if i.Package.Name == n {
			return &i.Package, true
		}
	}
	return nil, false
}

func (pms *PackageManagementSystem) GetPackageBySource(s string) (*Package, bool) {
	for _, i := range pms.Installer.Installations {
		if i.Package.Source == s {
			return &i.Package, true
		}
	}
	return nil, false
}

func (pms *PackageManagementSystem) GetInstallations(s string) Installations {
	installations := pms.Installer.Installations
	if s == "by-name" {
		SortInstallationsByName(installations)
	}
	return installations
}

func (pms *PackageManagementSystem) InstallRequirements(r *Requirements, p string, fr bool) error {
	for _, pkg := range r.Packages {
		if err := pms.InstallPkg(pkg, p, fr); err != nil {
			return err
		}
	}
	return nil
}

func (pms *PackageManagementSystem) InstallPkg(pkg Package, p string, fr bool) error {
	if !pms.Installer.Installed(pkg) {
		return pms.Installer.Install(pkg, p)
	} else {
		if fr {
			return pms.Installer.Reinstall(pkg, p)
		}
		// TODO implement (upgrade_flag, error returns)
	}
	return nil
}

func (pms *PackageManagementSystem) UninstallRequirements(r *Requirements, y bool) error {
	for _, pkg := range r.Packages {
		if err := pms.UninstallPkg(pkg, y); err != nil {
			return err
		}
	}
	return nil
}

func (pms *PackageManagementSystem) UninstallPkg(pkg Package, y bool) error {
	if y || messages.AskForConfirmation("Proceed (y/n)?") {
		return pms.Installer.Uninstall(pkg)
	} else {
		fmt.Printf("Canceled uninstall of %s", pkg.Name)
	}
	return nil
}
