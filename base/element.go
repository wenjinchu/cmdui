package base

import (
	"fmt"
	"strings"
)

// 定义元素的状态函数类型
type StatusFunc func(val interface{}) (style Style, fore, back Color)

type Widget interface {
	SetFore(Color)
	SetBack(Color)
	SetStyle(Style)
	SetWidth(int)
	SetDisplay(Display)
	SetStatusFunc(StatusFunc)

	GetColor() (fore, back Color)
	GetStyle() Style
	GetWidth() int
	GetDisplay() Display
	GetStatusFunc() StatusFunc

	Print()
	PrintToCoordinate(x, y int)
	UpdateAndPrint(interface{})
	UpdateAndPrintToCoordinate(x, y int, val interface{})
}

// 定义终端元素的实现类型
type Element struct {
	// 元素的状态
	fore            Color  // 前景色
	back            Color  // 后景色
	style           Style  // 字体样式
	coordinate      [2]int // 元素的起始坐标
	isUpdatedStatus bool   // 元素的状态是否更新

	// 元素的布局
	width           int         // 元素的宽度
	display         Display     // 内容的水平位置：左 - 中 - 右
	value           interface{} // 元素值
	statusFunc      StatusFunc  // 状态函数，根据元素值改变元素在终端的显示状态
	isUpdatedStruct bool        // 元素的结构状态是否更新

	cache string // 元素打印的缓存，元素的值、布局、宽度发生变化，缓存将更新
}

func NewElement(val interface{}) *Element {
	elem := &Element{
		fore:            DefaultForeColor,
		back:            DefaultBackColor,
		coordinate:      [2]int{-1, -1},
		value:           val,
		width:           -1,
		isUpdatedStatus: true,
		isUpdatedStruct: true,
	}
	elem.update(elem.formatString(elem.value))
	return elem
}

// 设置元素的前景色
func (own *Element) SetFore(color Color) {
	if own.fore != color {
		own.fore = color
		own.isUpdatedStatus = true
	}
}

// 设置元素的背景色
func (own *Element) SetBack(color Color) {
	if own.back != color {
		own.back = color
		own.isUpdatedStatus = true
	}
}

// 获取元素的前后景色
func (own *Element) GetColor() (fore, back Color) {
	return own.fore, own.back
}

// 设置元素的字体样式
func (own *Element) SetStyle(style Style) {
	if own.style != style {
		own.style = style
		own.isUpdatedStatus = true
	}
}

// 获取元素的字体样式
func (own *Element) GetStyle() Style {
	return own.style
}

// 设置元素的尺寸
func (own *Element) SetWidth(width int) {
	if width < 0 {
		return
	}

	if own.width != width {
		own.width = width
		own.update(own.formatString(own.value))
		own.isUpdatedStruct = true
	}

}

// 获取元素的最大水平空间
func (own *Element) GetWidth() int {
	return own.width
}

// 设置元素的(水平)布局格式：左 - 中 - 右
// 更改布局之后，需要更新元素值
func (own *Element) SetDisplay(display Display) {
	if own.display != display {
		if display != Left && display != Right {
			own.display = Center
		} else {
			own.display = display
		}

		own.update(own.formatString(own.value))
		own.isUpdatedStruct = true
	}

}

// 获取元素的布局格式
func (own *Element) GetDisplay() Display {
	return own.display
}

// 设置元素的起始坐标
func (own *Element) SetCoordinate(x, y int) {
	if x < -1 || y < -1 {
		return
	}

	if x == own.coordinate[0] && y == own.coordinate[1] {
		return
	}

	own.coordinate = [2]int{x, y}
	own.isUpdatedStatus = true
}

// 获取元素的起始坐标
func (own *Element) GetCoordinate() (x, y int) {
	return own.coordinate[0], own.coordinate[1]
}

// 设置元素的状态函数，根据元素值的变化而动态改变元素状态
func (own *Element) SetStatusFunc(statusFunc StatusFunc) {
	own.statusFunc = statusFunc
	own.isUpdatedStatus = true
}

// 获取元素的状态函数
func (own *Element) GetStatusFunc() StatusFunc {
	return own.statusFunc
}

// 获取元素的值
func (own *Element) GetValue() interface{} {
	return own.value
}

// 获取元素的输出字符串
func (own *Element) String() string {
	return own.cache
}

// 获取元素的输出字符串
func (own *Element) IsUpdated() bool {
	return own.isUpdatedStatus || own.isUpdatedStruct
}

//******************************************************************************

// 在光标当前位置打印
func (own *Element) Print() {
	// if !own.isUpdatedStatus && !own.isUpdatedStruct {
	// 	return
	// }

	// 检查是否指定坐标
	if own.coordinate[0] == -1 || own.coordinate[1] == -1 {
		own.SetCoordinate(GetCoordinate())
	}

	// 打印到终端
	PrintWithStatusToCoordinate(own.coordinate[0], own.coordinate[1], own.style, own.fore, own.back, own.cache)
	own.isUpdatedStatus = false
	own.isUpdatedStruct = false
}

// 打印到指定终端坐标
func (own *Element) PrintToCoordinate(x, y int) {
	own.SetCoordinate(x, y)

	if own.isUpdatedStatus || own.isUpdatedStruct {
		PrintWithStatusToCoordinate(x, y, own.style, own.fore, own.back, own.cache)
		own.isUpdatedStatus = false
		own.isUpdatedStruct = false
	}
}

// 在光标当前位置更新并打印
func (own *Element) UpdateAndPrint(val interface{}) {
	own.Update(val)
	own.Print()
}

// 更新并打印到指定终端坐标
func (own *Element) UpdateAndPrintToCoordinate(x, y int, val interface{}) {
	own.Update(val)
	own.PrintToCoordinate(x, y)
}

// 更新元素
func (own *Element) Update(val interface{}) string {
	if val != own.value {
		own.value = val
		own.update(own.formatString(own.value))
		own.isUpdatedStruct = true
	}

	return own.cache
}

func (own *Element) update(val string) {
	size := getStringSize(val)

	// 更新元素的输出长度
	switch {
	case own.width < 0: // 直接取内容
	case own.width < size: // 截取内容
		own.cache = string(val[:own.width])
	case own.width == size: // 刚好 ==> own.size<0
		own.cache = val

	// 更新元素的布局
	case own.width > size:
		switch own.display {
		// 居左布局
		case Left:
			own.cache = val + strings.Repeat(" ", own.width-size)
		// 居中布局
		case Center:
			idx := (own.width - size) / 2
			own.cache = fmt.Sprintf("%s%s%s", strings.Repeat(" ", idx), val, strings.Repeat(" ", own.width-size-idx))
		// 居右布局
		case Right:
			own.cache = strings.Repeat(" ", own.width-size) + val
		default:
			own.cache = val + strings.Repeat(" ", own.width-size)
		}
	}

	// 更新元素的状态
	if own.statusFunc != nil {
		own.style, own.fore, own.back = own.statusFunc(own.value)
	}
}

func (own *Element) formatString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}
