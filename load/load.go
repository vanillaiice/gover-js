package load

import (
	"encoding/json"
	"os"
)

// PackageData represents the package.json data.
type PackageData struct {
	Version string `json:"version"`
}

// FromFile loads the version data from file.
func FromFile(file string) (*PackageData, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var versionData PackageData
	if err = json.Unmarshal(data, &versionData); err != nil {
		return nil, err
	}

	return &versionData, nil
}

// FromFilePanic is the same as FromFile but panics on error.
func FromFilePanic(file string) (vd *PackageData) {
	var err error
	vd, err = FromFile(file)
	if err != nil {
		panic(err)
	}

	return
}
