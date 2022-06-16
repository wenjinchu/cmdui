package base

// 定义终端字体样式的类型
type Style int

var DefaultStyle Style = NoStyle

const (
	NoStyle   Style = iota // 没有样式
	HighLight              // 高亮
	HalfLight              // 亮度减半
	Italics                // 斜体
	Underline              // 下划线
	Blink                  // 闪烁
	Flash                  // 快闪
	Reverse                // 反显
	Blanking               // 隐藏
	Strikeout              // 删除线
)

var styleMap = map[Style][2]int{
	NoStyle:   [2]int{0, 0x0000},
	HighLight: [2]int{1, 0x0006}, // 高亮
	HalfLight: [2]int{2, 0x0000},
	Italics:   [2]int{3, 0x0000},
	Underline: [2]int{4, 0x8000},
	Blink:     [2]int{5, 0x0000},
	Flash:     [2]int{6, 0x0000},
	Reverse:   [2]int{7, 0x0000},
	Blanking:  [2]int{8, 0x0000},
	Strikeout: [2]int{9, 0x0000},
}

func getStyle(style Style) int {
	s, ok := styleMap[style]
	if !ok {
		return chooseByOS(styleMap[DefaultStyle])
	}
	return chooseByOS(s)
}
