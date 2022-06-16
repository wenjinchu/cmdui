package main

import (
	"cmdui/base"
	"cmdui/widget"
	"math/rand"
	"time"
)

func init() {
	base.HideCursor()
}

func main() {
	bar := widget.NewProgressbar(1024*1024 + 10000)
	bar.SetDesc("Downing...")
	rand.Seed(time.Now().Unix())
	for i := int64(0); i < 1000; i++ {
		bar.Load(rand.Int63n(5000))
		if bar.IsEnd() {
			return
		}
		time.Sleep(time.Millisecond * 300)
	}
}
