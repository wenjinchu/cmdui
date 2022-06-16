package widget

import (
	"cmdui/base"
	"strings"
)

// 定义行组件的实现类型
type Line struct {
	coordinate [2]int // 坐标
	hasDelim   bool   // 是否有界定字符

	delim string          // 界定字符
	elems []*base.Element // 数据元素
}

func NewLine() *Line {
	return &Line{
		coordinate: [2]int{-1, -1},
		elems:      make([]*base.Element, 0),
	}
}

// 设置坐标
func (own *Line) SetCoordinate(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	own.coordinate = [2]int{x, y}
	for _, elem := range own.elems {
		elem.SetCoordinate(-1, -1)
	}
}

// 设置界定符
func (own *Line) SetDelim(delim string) {
	own.delim = delim
	own.hasDelim = true
}

// 添加元素
func (own *Line) Add(elem *base.Element) {
	own.elems = append(own.elems, elem)
}

// 根据索引获取元素
func (own *Line) Get(idx int) *base.Element {
	return own.elems[idx]
}

// 获取完整的字符串
func (own *Line) String() string {
	var builder strings.Builder
	for _, elem := range own.elems {
		if own.hasDelim {
			builder.WriteString(own.delim)
		}
		builder.WriteString(elem.String())
	}

	if own.hasDelim {
		builder.WriteString(own.delim)
	}
	return builder.String()
}

// 打印到终端指定坐标
func (own *Line) PrintToCoordinate(x, y int) {
	own.SetCoordinate(x, y)
	own.Print()
}

// 打印到终端
func (own *Line) Print() {
	// 设置终端坐标
	if own.coordinate[0] != -1 && own.coordinate[1] != -1 {
		base.SetCoordinate(own.coordinate[0], own.coordinate[1])
	}

	for _, elem := range own.elems {
		// 打印界定符到终端
		if own.hasDelim {
			base.PrintFormat(own.delim)
		}

		// 打印数据到终端
		elem.Print()
	}

	// 打印最后一个界定符到终端
	if own.hasDelim {
		base.PrintFormat(own.delim)
	}
}
