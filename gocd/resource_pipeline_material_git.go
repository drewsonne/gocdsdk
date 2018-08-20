package gocd

import (
	"errors"
)

func (mag MaterialAttributesGit) equal(a2i MaterialAttribute) (bool, error) {
	var ok bool
	a2, ok := a2i.(MaterialAttributesGit)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}
	urlsEqual := mag.URL == a2.URL
	branchesEqual := mag.Branch == a2.Branch ||
		mag.Branch == "" && a2.Branch == "master" ||
		mag.Branch == "master" && a2.Branch == ""

	if !urlsEqual {
		return false, nil
	}
	return branchesEqual, nil
}

// GenerateGeneric form (map[string]interface) of the material filter
func (mag MaterialAttributesGit) GenerateGeneric() (ma map[string]interface{}) {
	ma = map[string]interface{}{}

	for _, pair := range []struct {
		key   string
		value string
	}{
		{key: "name", value: mag.Name},
		{key: "url", value: mag.URL},
		{key: "branch", value: mag.Branch},
		{key: "submodule_folder", value: mag.SubmoduleFolder},
		{key: "destination", value: mag.Destination},
	} {
		if pair.value != "" {
			ma[pair.key] = pair.value
		}
	}

	for _, pair := range []struct {
		key   string
		value bool
	}{
		{key: "auto_update", value: mag.AutoUpdate},
		{key: "shallow_clone", value: mag.ShallowClone},
		{key: "invert_filter", value: mag.InvertFilter},
	} {
		if pair.value {
			ma[pair.key] = pair.value
		}
	}

	if f := mag.Filter.GenerateGeneric(); f != nil {
		ma["filter"] = f
	}
	return
}

// HasFilter in this material attribute
func (mag MaterialAttributesGit) HasFilter() bool {
	return true
}

// GetFilter from material attribute
func (mag MaterialAttributesGit) GetFilter() *MaterialFilter {
	return mag.Filter
}

// UnmarshallInterface from a JSON string to a MaterialAttributesGit struct
func unmarshallMaterialAttributesGit(mag *MaterialAttributesGit, i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mag.Name = value.(string)
		case "url":
			mag.URL = value.(string)
		case "auto_update":
			mag.AutoUpdate = value.(bool)
		case "branch":
			mag.Branch = value.(string)
		case "submodule_folder":
			mag.SubmoduleFolder = value.(string)
		case "destination":
			mag.Destination = value.(string)
		case "shallow_clone":
			mag.ShallowClone = value.(bool)
		case "invert_filter":
			mag.InvertFilter = value.(bool)
		case "filter":
			if v1, ok1 := value.(map[string]interface{}); ok1 {
				mag.Filter = unmarshallMaterialFilter(v1)
			}
		}
	}
}
