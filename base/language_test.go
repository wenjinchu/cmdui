package base

import (
	"testing"
)

func TestGetStringSize(t *testing.T) {
	t.Log(len([]rune("姓名")))
	t.Log(len("姓名"))

	t.Log(getStringSize("hello"))
	t.Log(getStringSize("姓名"))
}
