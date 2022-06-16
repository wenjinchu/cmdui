package widget

import (
	"testing"
	"time"
)

func TestProgressbar(t *testing.T) {
	bar := NewProgressbar(100)
	bar.Load(10)
	time.Sleep(1 * time.Second)
	t.Log(bar.speed)
	bar.Load(89)
	t.Log(bar.speed)
	time.Sleep(1 * time.Second)
	bar.Load(1)
	t.Log(bar.speed)

	FormatBytes(1655275485987380400)
}
