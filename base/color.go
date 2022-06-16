package base

// 定义终端颜色的类型
type Color int

var (
	DefaultForeColor = White // 默认前景色
	DefaultBackColor = Black // 默认背景色
)

const (
	Black  Color = iota + 1 // 黑色
	Blue                    // 蓝色
	Green                   // 绿色
	Red                     // 红色
	Purple                  // 紫色
	Yellow                  // 黄色
	White                   // 白色
	Grey                    // 灰色
)

var (
	foreColorMap = map[Color][2]int{
		Black:  [2]int{30, 0x0000},
		Blue:   [2]int{34, 0x0001},
		Green:  [2]int{32, 0x0002},
		Red:    [2]int{31, 0x0004},
		Purple: [2]int{35, 0x0005},
		Yellow: [2]int{33, 0x0006},
		White:  [2]int{37, 0x0007},
		Grey:   [2]int{30, 0x0008},
	}

	backColorMap = map[Color][2]int{
		Black:  [2]int{40, 0x0000},
		Blue:   [2]int{44, 0x0010},
		Green:  [2]int{42, 0x0020},
		Red:    [2]int{41, 0x0040},
		Purple: [2]int{45, 0x0050},
		Yellow: [2]int{43, 0x0060},
		White:  [2]int{47, 0x0070},
		Grey:   [2]int{43, 0x0080},
	}
)

const (
	posFore byte = 'f' // 前景色
	posBack byte = 'b' // 背景色
)

func getColor(pos byte, color Color) int {
	var opts [2]int
	var isExist bool

	if pos == posFore {
		opts, isExist = foreColorMap[color]
	} else {
		opts, isExist = backColorMap[color]
	}

	if !isExist {
		if pos == posFore {
			return chooseByOS(foreColorMap[DefaultForeColor])
		}
		return chooseByOS(backColorMap[DefaultBackColor])
	}
	return chooseByOS(opts)
}
