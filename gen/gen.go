package gen

import (
	"encoding/json"
	"fmt"
	"os"

	oj "github.com/vanillaiice/gover-js/ordered-json"
)

// PackageJsonData is the package.json data.
type PackageJsonData struct {
	Version string                 `json:"version"`
	Other   map[string]interface{} `json:"-"`
}

// UpdatePackageVersion updates package.json with the new version.
func UpdatePackageVersion(filePath string, version string) (err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	var oj oj.OrderedJSON
	if err := json.Unmarshal(content, &oj); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	_, ok := oj.Data["version"]
	if !ok {
		return fmt.Errorf("version not found in package.json")
	}

	oj.Data["version"] = version

	updatedContent, err := json.MarshalIndent(oj, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	return os.WriteFile(filePath, updatedContent, 0644)
}
