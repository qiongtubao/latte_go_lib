package xhr

import (
	"testing"
)

func Test_NoQueryGet(t *testing.T) {
	xhr, err := Get("http://www.baidu.com")
	if err != nil {
		t.Error("Get Init Error")
	}
	_, err = xhr.Send()
	if err != nil {
		t.Error("Get Init Error")
	}
	//fmt.Println(str)
}
