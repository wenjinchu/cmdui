package main

import (
	"cmdui/base"
)

func main() {
	// 前景色
	base.PrintWithStatus(base.NoStyle, base.Red, base.Black, "red")
	base.PrintWithStatus(base.NoStyle, base.Green, base.Black, "green")
	base.PrintWithStatus(base.NoStyle, base.Yellow, base.Black, "yellow")
	base.PrintWithStatus(base.NoStyle, base.Blue, base.Black, "blue")
	base.PrintWithStatus(base.NoStyle, base.Grey, base.Black, "grey")
	base.PrintWithStatus(base.NoStyle, base.White, base.Black, "black")
	base.PrintWithStatus(base.NoStyle, base.Purple, base.Black, "purple")
	base.PrintFrmat("\n")

	// 背景色
	base.PrintWithStatus(base.NoStyle, base.Black, base.Red, "red")
	base.PrintWithStatus(base.NoStyle, base.Black, base.Green, "green")
	base.PrintWithStatus(base.NoStyle, base.Black, base.Yellow, "yellow")
	base.PrintWithStatus(base.NoStyle, base.Black, base.Blue, "blue")
	base.PrintWithStatus(base.NoStyle, base.Black, base.Grey, "grey")
	base.PrintWithStatus(base.NoStyle, base.Black, base.White, "white")
	base.PrintWithStatus(base.NoStyle, base.Black, base.Purple, "purple")
	base.PrintFrmat("\n")

	// 样式
	base.PrintWithStatus(base.HalfLight, base.White, base.Black, "half light")
	base.PrintWithStatus(base.Underline, base.White, base.Black, "underline")
	base.PrintFrmat("\n")
}
