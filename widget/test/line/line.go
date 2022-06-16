package main

import (
	"cmdui/base"
	"cmdui/widget"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	base.HideCursor()
}

func main() {
	line := widget.NewLine()
	line.SetDelim("|")

	elem1 := base.NewElement("速度")
	elem1.SetBack(base.Green)
	elem1.SetWidth(10)
	line.Add(elem1)

	elem2 := base.NewElement(0)
	elem2.SetBack(base.Blue)
	elem2.SetWidth(10)
	elem2.SetStatusFunc(func(val interface{}) (base.Style, base.Color, base.Color) {
		v := val.(int)
		if v < 60 {
			return base.NoStyle, base.Red, base.White
		} else if v < 80 {
			return base.NoStyle, base.Blue, base.Grey
		} else {
			return base.NoStyle, base.Green, base.Blue
		}
	})
	line.Add(elem2)

	elem3 := base.NewElement("安全")
	elem3.SetBack(base.Red)
	elem3.SetWidth(10)
	line.Add(elem3)
	line.Print()

	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go func() {
		for j := 0; j < 100; j++ {
			elem2.UpdateAndPrint(rand.Intn(100))
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()
	time.Sleep(2 * time.Second)

	_, y := base.GetCoordinate()
	line.SetCoordinate(3, 2+y)
	line.Print()
	time.Sleep(2 * time.Second)

	line.PrintToCoordinate(6, y+4)
	wg.Wait()
	base.PrintFormat("\n")
}
