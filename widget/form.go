package widget

import (
	"cmdui/base"
	"strings"
)

var (
	VDelim = "|" // 垂直分隔符
	HDelim = "-" // 水平分隔符
	CDelim = "+" // 交叉分隔符
)

// 根据元素的宽度生成水平分隔字符串
func GenDelimString(delim string, elem *base.Element) string {
	return strings.Repeat(delim, elem.GetWidth()) + CDelim
}

// 定义表的底层实现类型
type form struct {
	hSplitLine string // 水平分割行线
	headString string // 表头字符串
	isPrinting bool   // 数据是否正在打印

	num  int      // 字段数量
	head [2]*Line // 数据表头，第一个是表头本身，第二个用于处理数据

	// 各个字段的全局宽度
	// 该参数将覆盖字段本身的宽度参数
	fieldWidths []int
}

// 新建一个表
func newForm(fieldWidths ...int) *form {
	f := &form{
		hSplitLine:  CDelim,
		headString:  VDelim,
		isPrinting:  false,
		head:        [2]*Line{NewLine(), NewLine()},
		fieldWidths: fieldWidths,
	}
	f.head[1].SetDelim(VDelim)
	return f
}

// 自动获取新增字段的宽度
func (own *form) getWidth() int {
	num := len(own.fieldWidths)
	if own.num <= num {
		return own.fieldWidths[own.num-1]
	}
	return own.fieldWidths[len(own.fieldWidths)-1]
}

// 新增一个字段
func (own *form) addField(field *base.Element) {
	own.num++
	own.addDataField(field)
	own.addHeadField(field)
}

// 新增字段的隐藏字段 - 用于处理数据
func (own *form) addDataField(field *base.Element) {
	fore, back := field.GetColor()

	_filed := base.NewElement("")
	_filed.SetFore(fore)
	_filed.SetBack(back)
	_filed.SetStyle(field.GetStyle())
	_filed.SetWidth(own.getWidth())
	_filed.SetDisplay(field.GetDisplay())
	_filed.SetStatusFunc(field.GetStatusFunc())
	own.head[1].Add(_filed)
	own.hSplitLine += GenDelimString(HDelim, _filed) // 更新水平分割线
}

// 新增字段的表头字段
func (own *form) addHeadField(field *base.Element) {
	elem := base.NewElement(field.GetValue())
	elem.SetWidth(own.getWidth())
	own.head[0].Add(elem)
	own.headString += elem.String() + VDelim // 更新表头字符串
}

// 处理数据 - 标准化数据的尺寸
func (own *form) dealData(data []interface{}) []interface{} {
	size := len(data)
	if size < own.num {
		for i := 0; i < own.num-size; i++ {
			data = append(data, "")
		}
	} else if size > own.num {
		data = data[:own.num]
	}
	return data
}
