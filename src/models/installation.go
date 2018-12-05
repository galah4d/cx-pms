package models

import (
	"encoding/json"
	"errors"
	"github.com/galah4d/cx-pms/config"
	"github.com/galah4d/cx-pms/src/utils"
	"os"
	"path/filepath"
	"time"
)

type Installer struct {
	Installations []*Installation `json:"installations"`
}

func (i *Installer) UnmarshalJSON(f string) error {
	data, err := utils.ReadJson(f)
	if err != nil {
		return err // File not found
	}
	return json.Unmarshal(data, i)
}

func (i *Installer) MarshalJSON(f string) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return utils.DumpJson(b, f)
}

func (i *Installer) Install(pkg Package, path string) error {
	if i.Installed(pkg) {
		return errors.New("package already installed")
	} else {
		if err := pkg.install(path); err != nil {
			return errors.New("package installation failed")
		}
		i.Installations = append(i.Installations, newInstallation(pkg, path))
		return i.MarshalJSON(filepath.Join(os.Getenv("GOPATH"), config.InstallationFilePATH))
	}
}

func (i *Installer) Uninstall(pkg Package) error {
	for index, installation := range i.Installations {
		if installation.Package.Source == pkg.Source {
			if err := pkg.uninstall(installation.Path); err != nil {
				return errors.New("package uninstall failed")
			}
			// Removes the installation
			copy(i.Installations[index:], i.Installations[index+1:])
			i.Installations[len(i.Installations)-1] = nil
			i.Installations = i.Installations[:len(i.Installations)-1]
			return i.MarshalJSON(filepath.Join(os.Getenv("GOPATH"), config.InstallationFilePATH))
		}
	}
	return errors.New("package not installed")
}

func (i *Installer) Reinstall(pkg Package, path string) error {
	if err := i.Uninstall(pkg); err != nil {
		return err
	}
	return i.Install(pkg, path)
}

func (i Installer) Installed(pkg Package) bool {
	for _, installation := range i.Installations {
		if installation.Package.Source == pkg.Source { // TODO implement comparator
			return true
		}
	}
	return false
}

func (i Installer) GetInstallationPath(pkg Package) (string, error) {
	for _, installation := range i.Installations {
		if installation.Package.Source == pkg.Source {
			return installation.Path, nil
		}
	}
	return "", errors.New("package not installed installed")
}

type Installation struct {
	Date    time.Time `json:"date"`
	Path    string    `json:"path"`
	Package Package   `json:"package"`
}

func newInstallation(pkg Package, path string) *Installation {
	i := &Installation{
		Date:    time.Now(),
		Path:    path,
		Package: pkg,
	}
	return i
}
