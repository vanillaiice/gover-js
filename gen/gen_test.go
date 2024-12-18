package gen_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/vanillaiice/gover-js/gen"
)

func TestUpdatePackageVersion(t *testing.T) {
	const filePath = "package_test.json"
	const want = "6.9.420"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	err = gen.UpdatePackageVersion(filePath, want)
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	var packageJsonData gen.PackageJsonData
	err = json.Unmarshal(content, &packageJsonData)
	if err != nil {
		t.Fatal(err)
	}

	if packageJsonData.Version != want {
		t.Errorf("got %q, want %q", packageJsonData.Version, want)
	}

	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
