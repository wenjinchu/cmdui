package base

type Language int

const (
	English Language = iota
	Chinese
)

// 获取字符串在终端的占位尺寸
func getStringSize(s string) int {
	runeNum := len([]rune(s))
	byteNum := len(s)
	switch {
	// 1个中文占3个字节，在终端占2个字符
	case byteNum == runeNum*3:
		return runeNum * 2
	default:
		return byteNum
	}
}
