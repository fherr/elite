package elite

import (
	"testing"
)

func TestNameGen(t *testing.T) {
	name := GetRandomName()
	t.Logf("name: %s", name)
	if name == "" {
		t.Errorf("got an empty string")
	}
}
