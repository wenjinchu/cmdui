package base

func HideCursor() {
	if osName == osWin {
		hideWindowsCursor()
		return
	} else {

	}
	panic("不支持系统：" + osName)
}

//******************************************************************************

func GetCoordinate() (int, int) {
	if osName == osWin {
		return getWindowsCursorCoordinate()
	} else {

	}
	panic("不支持系统：" + osName)
}

func SetCoordinate(x, y int) {
	if osName == osWin {
		setWindowsCursorCoordinate(x, y)
	} else {

	}
}

//******************************************************************************

// 格式化打印状态
func PrintFormat(msg string) {
	if osName == osWin {
		printToWindowsWithStatus(
			getStyle(DefaultStyle),
			getColor(posFore, DefaultForeColor),
			getColor(posBack, DefaultBackColor),
			msg)
	} else {

	}
}

// 打印带状态的字符串
func PrintToCoordinate(x, y int, msg string) {
	if osName == osWin {
		printToWindowsToCoordinate(x, y, msg)
	} else {

	}
}

// 打印无状态的字符串到指定坐标
func PrintWithStatus(style Style, fore, back Color, msg string) {
	if osName == osWin {
		printToWindowsWithStatus(
			getStyle(style),
			getColor(posFore, fore),
			getColor(posBack, back),
			msg)
	} else {

	}
}

// 打印带状态的字符串到指定坐标
func PrintWithStatusToCoordinate(x, y int, style Style, fore, back Color, msg string) {
	if osName == osWin {
		printToWindowsWithStatusToCoordinate(
			x, y,
			getStyle(style),
			getColor(posFore, fore),
			getColor(posBack, back),
			msg)
	} else {

	}
}
