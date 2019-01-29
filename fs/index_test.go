package fs

import (
	"testing"
)

func Test_WriteFile(t *testing.T) {
	error := WriteFile("test.txt", "test string")
	if error != nil {
		t.Error("write File error")
	} else {
		t.Log("write File ok")
	}
}
func Test_ReadFile(t *testing.T) {
	data, error := ReadFile("test.txt")
	if error != nil {
		t.Error(error)
	} else {
		t.Log("readFile", data)
	}

}
