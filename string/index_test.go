package string

import (
	"fmt"
	"testing"
	"time"
)

func Test_TimeFormat(t *testing.T) {
	fmt.Println(TimeFormat("YYYY-MM-dd hh:mm:ss", time.Now()))
}
