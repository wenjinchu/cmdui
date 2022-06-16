package base

import "runtime"

var osName string

const (
	osWin = "windows"
	osLin = "linux"
)

func init() {
	osName = runtime.GOOS
}

// 根本系统选择参数：0-Linux；1-Windows
func chooseByOS(opts [2]int) int {
	switch osName {
	case osLin:
		return opts[0]
	case osWin:
		return opts[1]
	}
	return 0
}
