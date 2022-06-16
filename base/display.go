package base

// 定义终端元素布局的类型
type Display int

var DefaultDisplay = Center // 默认布局

const (
	Center Display = iota // 中
	Left                  // 左
	Right                 // 右
	Upper                 // 上
	Lower                 // 下
)
