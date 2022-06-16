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
	form := widget.NewStatusForm(5, 10)

	elem := base.NewElement("公司")
	elem.SetBack(base.Blue)
	form.AddField(elem)

	elem = base.NewElement("姓名")
	elem.SetFore(base.Green)
	elem.SetDisplay(base.Left)
	form.AddField(elem)

	elem = base.NewElement("性别")
	elem.SetStatusFunc(func(val interface{}) (base.Style, base.Color, base.Color) {
		if v, ok := val.(string); ok {
			if v == "男" {
				return base.NoStyle, base.Blue, base.Red
			} else if v == "女" {
				return base.NoStyle, base.Red, base.Green
			}
		}
		return base.NoStyle, base.DefaultForeColor, base.DefaultBackColor
	})
	form.AddField(elem)

	elem = base.NewElement("年龄")
	elem.SetDisplay(base.Right)
	form.AddField(elem)
	form.Print()

	sexMap := map[int]interface{}{
		0: "男",
		1: "女",
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		c := rand.Intn(5)
		switch {
		case c == 0:
			form.Update(0, []interface{}{0, rand.Intn(5), sexMap[rand.Intn(2)], rand.Intn(100)})
		case c == 1:
			form.Update(1, []interface{}{1, rand.Intn(5), sexMap[rand.Intn(2)], rand.Intn(100)})
		case c == 2:
			form.Update(2, []interface{}{2, rand.Intn(5), sexMap[rand.Intn(2)], rand.Intn(100)})
		case c == 3:
			form.Update(3, []interface{}{3, rand.Intn(5), sexMap[rand.Intn(2)], rand.Intn(100)})
		case c == 4:
			form.Update(4, []interface{}{4, rand.Intn(5), sexMap[rand.Intn(2)], rand.Intn(100)})
		}
		time.Sleep(300 * time.Millisecond)
	}
	form.End()
}
