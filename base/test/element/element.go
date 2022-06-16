package main

import (
	"cmdui/base"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	base.DefaultBackColor = base.Yellow
	base.HideCursor()
}

func main() {
	fmt.Println("====================================================")
	elem := base.NewElement("Red")
	elem.SetWidth(20)
	elem.SetDisplay(base.Left)
	elem.SetFore(base.Red)
	elem.SetBack(base.Purple)
	elem.Print()
	base.PrintFormat("\n")

	elem = base.NewElement("")
	elem.SetWidth(20)
	elem.SetDisplay(base.Center)
	elem.SetFore(base.Blue)
	elem.SetBack(base.Green)
	elem.UpdateAndPrint("Blue")
	base.PrintFormat("\n")

	elem = base.NewElement("")
	elem.SetWidth(20)
	elem.SetDisplay(base.Right)
	elem.SetFore(base.Green)
	elem.SetBack(base.Blue)
	elem.UpdateAndPrint("Green")
	base.DefaultBackColor = base.Black
	base.PrintFormat("\n====================================================\n")

	elem = base.NewElement("")
	elem.SetWidth(15)
	elem.SetDisplay(base.Left)
	elem.SetFore(base.Red)
	elem.SetBack(base.Purple)
	elem.UpdateAndPrint(true)
	base.PrintFormat("\n")

	elem1 := base.NewElement(0)
	elem1.SetWidth(15)
	elem1.SetStatusFunc(func(val interface{}) (sytle base.Style, fore, back base.Color) {
		v := val.(int)
		if v < 60 {
			return base.NoStyle, base.Red, base.White
		} else if v < 80 {
			return base.NoStyle, base.Blue, base.Purple
		} else {
			return base.NoStyle, base.Green, base.Red
		}
	})
	elem1.Print()
	base.PrintFormat("\n")

	elem = base.NewElement("")
	elem.SetWidth(15)
	elem.SetDisplay(base.Right)
	elem.SetFore(base.Green)
	elem.SetBack(base.Blue)
	elem.UpdateAndPrint(false)

	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go func() {
		for j := 0; j < 100; j++ {
			elem1.UpdateAndPrint(rand.Intn(100))
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()
	wg.Wait()
	base.PrintFormat("\n\n")
}
