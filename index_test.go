package lib

import (
	"testing"
)

func Test_IntToStr(t *testing.T) {
	str := IntToStr(1)
	if str != "1" {
		t.Error("IntToStr error")
	}
}

func Test_StrToInt(t *testing.T) {
	interger, err := StrToInt("1")
	if err != nil {
		t.Error("err")
	}
	if interger != 1 {
		t.Error("IntToStr error")
	}
}
