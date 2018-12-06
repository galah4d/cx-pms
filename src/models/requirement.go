package models

import (
	"encoding/json"
	"fmt"
	"github.com/galah4d/cx-pms/src/utils"
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

func (r *Requirements) Load(f string) error {
	data, err := utils.ReadJson(f)
	if err != nil {
		return err // File not found
	}
	return json.Unmarshal(data, r)
}

func (r *Requirements) Save(f string) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return utils.DumpJson(b, f)
}

func (r *Requirements) AddRequirement(pkg Package) {
	r.Packages = append(r.Packages, pkg)
}
