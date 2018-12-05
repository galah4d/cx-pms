package models

import (
	"github.com/galah4d/cx-pms/src/utils"
	"encoding/json"
	"fmt"
)

type Requirements struct {
	Packages []Package `json:"requirements"`
}

func LoadRequirements(f string) (*Requirements, error) {
	byteValue, err := utils.ReadJson(f)
	if err != nil {
		return nil, err // File not found
	}

	var requirements Requirements
	err = json.Unmarshal(byteValue, &requirements)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &requirements, nil
}
/*
func (reqs Requirements) Install (target string, noDeps bool, forceReinstall bool) error {
	fmt.Println("[+] Installing requirements...")
	for _, pkg := range reqs.Packages {
		// If package isn´t installed install it from source
		if !pkg.IsInstalled() {
			if err := pkg.install(target); err != nil {
				fmt.Println(err.Error())
			}
		// Otherwise
		} else {
			if forceReinstall {
				pkg.reinstall()
				pkg.uninstall()
			}
		}


	}
	fmt.Println("[+] Done")
	return nil
}*/


