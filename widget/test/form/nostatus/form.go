package main

import (
	"cmdui/base"
	"cmdui/widget"
	"time"
)

func init() {
	base.HideCursor()
}

func main() {
	form := widget.NewNoStatusForm(10, 7, 13)
	form.AddField(base.NewElement("序号"))

	elem := base.NewElement("姓名")
	elem.SetDisplay(base.Left)
	form.AddField(elem)

	form.AddField(base.NewElement("性别"))

	elem = base.NewElement("年龄")
	elem.SetDisplay(base.Right)
	form.AddField(elem)

	datum := [][]interface{}{
		[]interface{}{1, "gao", "男", 35},
		[]interface{}{2, "wen", "男", 28},
		[]interface{}{3, "liu", "女", 25, 10000},
		[]interface{}{4, "cai", "女", 30},
		[]interface{}{5, "wei", "女"},
	}

	for _, data := range datum {
		form.Print(data)
		time.Sleep(1 * time.Second)
	}
}
