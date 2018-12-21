package fs

import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	error := WriteFile("test.txt", "test string")
	if error != nil {
		t.Error("write File error")
	} else {
		t.Log("write File ok")
	}
}
func TestReadFile(t *testing.T) {
	data, error := ReadFile("test.txt")
	if error != nil {
		t.Error(error)
	} else {
		t.Log("readFile", data)
	}

}
