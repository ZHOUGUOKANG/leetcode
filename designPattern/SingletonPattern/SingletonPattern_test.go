package SingletonPattern

import (
	"testing"
)

func Test_SingletonPattern_GetInstance_Echo(t *testing.T) {
	result := SingleObject{}.GetInstance().Echo()
	if result != "SingleObject Echo" {
		t.Error()
	}
}
