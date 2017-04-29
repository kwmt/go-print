package print

import(
	"testing"
)

func Test(t *testing.T) {
	expect := "hell"
	actual := Print("hello")
	if expect != actual {
		t.Errorf("actual: %s, expect: %s", actual, expect)
	}
}