/*******************************************************************************
Method: (并发不安全的)最小进度条
Author: Lemine
Langua: Golang 1.16
Modify：2021/04/22
*******************************************************************************/
package widget

import (
	"cmdui/base"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	defaultBarSize  = 50
	defaultBarStyle = [2]string{"█", " "}

	defaultFormatFunc = func(cursor, total int64) string {
		if total < cursor {
			cursor = total
		}
		cursorVal, cursorUnit := FormatBytes(cursor)
		totalVal, totalUnit := FormatBytes(total)
		return fmt.Sprintf("%0.2f%s/%0.2f%s", cursorVal, cursorUnit, totalVal, totalUnit)
	}

	defaultSpeedFunc = func(add, seconds int64) string {
		addVal, addUnit := FormatBytes(add)
		return fmt.Sprintf("%0.2f%s/s", addVal/float64(seconds), addUnit)
	}
)

type (
	FormatFunc func(cursor, total int64) string
	SpeedFunc  func(add, seconds int64) string
)

//最小进度条
type Progressbar struct {
	y int // 进度条在终端的Y坐标

	desc    string  // 进度条任务描述
	percent float64 // 当前进度百分比

	size   int       // 进度条总长度
	style  [2]string // style 进度条显示的样式：style[0]-填充样式；style[1]-留白样式
	length int       // 当前进度条长度

	total  int64 // 数据总量
	cursor int64 // 当前数据量

	startTime int64  // 开始时间
	nowTime   int64  // 当前时间
	cost      int64  // 耗时
	speed     string // 速度
	isStart   bool   // 开始
	isEnd     bool   // 结束

	formatFunc FormatFunc
	speedFunc  SpeedFunc
}

func NewProgressbar(total int64) *Progressbar {
	return &Progressbar{
		y:          -1,
		total:      total,
		size:       defaultBarSize,
		style:      defaultBarStyle,
		formatFunc: defaultFormatFunc,
		speedFunc:  defaultSpeedFunc,
	}
}

func (own *Progressbar) SetSize(size int) {
	own.size = size
}

func (own *Progressbar) GetSize() int {
	return own.size
}

func (own *Progressbar) SetStyle(style [2]string) {
	own.style = style
}

func (own *Progressbar) GetStyle() [2]string {
	return own.style
}

func (own *Progressbar) SetDesc(desc string) {
	own.desc = desc
}

func (own *Progressbar) SetYCoordinate(y int) {
	if y < 0 {
		return
	}
	own.y = y
}

func (own *Progressbar) GetYCoordinate() int {
	return own.y
}

func (own *Progressbar) IsEnd() bool {
	return own.isEnd
}

func (own *Progressbar) Load(add int64) {
	// 检查是否开始
	if !own.isStart {
		own.isStart = true
		own.startTime = time.Now().UnixNano()
		own.nowTime = own.startTime
	}

	// 检查是否结束
	if own.isEnd {
		return
	}

	// 计算相关参数
	own.calculate(add)

	// 检查起始坐标
	if own.y < 0 {
		_, own.y = base.GetCoordinate()
	}

	h, m, s := FormatSeconds(own.cost)
	bar := fmt.Sprintf("%s %0.2f%% |%s%s|(%s, %s)[%s:%s:%s]%s",
		own.desc, own.percent,
		strings.Repeat(own.style[0], own.length),
		strings.Repeat(own.style[1], own.size-own.length+1),
		own.formatFunc(own.cursor, own.total),
		own.speed,
		fillZero(h), fillZero(m), fillZero(s),
		strings.Repeat(" ", 30),
	)
	base.PrintWithStatusToCoordinate(0, own.y, base.NoStyle, base.Green, base.DefaultBackColor, bar)

	if own.cursor == own.total {
		own.isEnd = true
		base.PrintFormat("\n")
	}
}

func (own *Progressbar) calculate(add int64) {
	own.cursor += add
	if own.cursor > own.total {
		own.length = own.size
		own.cursor = own.total
		own.isEnd = true
	}
	// 获取进度百分比
	own.percent = float64(own.cursor) / float64(own.total) * 100
	// 获取进度长度
	own.length = int(math.Ceil((own.percent * (float64(own.size) / float64(100)))))
	// 计算速度
	now := time.Now().UnixNano()
	own.speed = own.speedFunc(add, (now-own.nowTime)/1000000000+1) // +1防止无穷小
	own.nowTime = now

	// 计算耗时
	own.cost = (own.nowTime - own.startTime) / 1000000000
}

func fillZero(num int64) string {
	switch {
	case num < 10:
		return "0" + strconv.Itoa(int(num))
	case num < 60:
		return strconv.Itoa(int(num))
	default:
		return "60"
	}
}

//******************************************************************************

func FormatBytes(size int64) (float64, string) {
	switch {
	case size < 1024:
		return float64(size), "B"
	case size < (1024 * 1024):
		return float64(size) / float64(1024), "KB"
	case size < (1024 * 1024 * 1024):
		return float64(size) / float64(1024*1024), "MB"
	case size < (1024 * 1024 * 1024 * 1024):
		return float64(size) / float64(1024*1024*1024), "GB"
	case size < (1024 * 1024 * 1024 * 1024 * 1024):
		return float64(size) / float64(1024*1024*1024*1024), "TB"
	default:
		return float64(-1), "UNKOWN"
	}
}

func FormatSeconds(seconds int64) (int64, int64, int64) {
	return seconds / 3600, (seconds / 60) % 60, seconds % 60
}
