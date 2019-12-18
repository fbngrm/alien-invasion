package invasion

import (
	"reflect"
	"testing"
)

var alienNameTest = struct {
	d string   // description of test case
	n int      // num aliens
	a []string // expected names
}{
	d: "expect 5 aliens names",
	n: 5,
	a: []string{"alien 1", "alien 2", "alien 3", "alien 4", "alien 5"},
}

func TestInitAlienNames(t *testing.T) {
	an := initAlienNames(alienNameTest.n)
	if !reflect.DeepEqual(alienNameTest.a, an) {
		t.Errorf("test case failed: %s", alienNameTest.d)
	}
}
