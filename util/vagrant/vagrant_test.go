//
package vagrant

import (
	"fmt"
	"os/exec"
	"testing"
)

// test if Exists works as intended
func TestExists(t *testing.T) {

	exists := false
	if err := exec.Command("vagrant", "-v").Run(); err == nil {
		exists = true
	}

	//
	testExists := Exists()

	if testExists != exists {
		t.Error("Results don't match!")
	}
}

// tests if appNameToPort works as intended
func TestAppNameToPort(t *testing.T) {
	name := "test-app"
	target := "11816" // pre-calculated port for "test-app"
	actual := appNameToPort(name)

	if actual != target {
		t.Error(fmt.Sprintf("Expected port '%s' got '%s'", target, actual))
	}
}

// tests to ensure that ports generated by appNameToPort are unique for similar
// names
func TestAppNameToPortUnique(t *testing.T) {
	var name string

	name = "test-app"
	first := appNameToPort(name)

	name = "app-test"
	second := appNameToPort(name)

	if first == second {
		t.Error("Ports arent unique")
	}
}
