package widget

import (
	"cmdui/base"
	"fmt"
	"strings"
)

// 定义状态表的类型
// 表格中的数据可实时变化，且数据量必须确定
type StatusForm struct {
	*form

	y      int // 数据行坐标
	rowNum int // 表的行数
}

// 新建一个状态表
func NewStatusForm(rowNum int, fieldWidths ...int) *StatusForm {
	return &StatusForm{
		form:   newForm(fieldWidths...),
		rowNum: rowNum,
	}
}

// 添加一个表字段
func (own *StatusForm) AddField(field *base.Element) {
	own.addField(field)
}

// 添加一个表字段
func (own *StatusForm) Print() {
	if !own.isPrinting {
		_, own.y = base.GetCoordinate()
		fmt.Println(own.hSplitLine)
		fmt.Println(own.headString)
		fmt.Println(own.hSplitLine)
		fmt.Println(own.initFormString(own.rowNum))
		own.isPrinting = true
	}
}

// 更新表数据
func (own *StatusForm) Update(idx int, data []interface{}) {
	if !own.isPrinting {
		own.Print()
	}

	data = own.dealData(data)
	for i := 0; i < own.num; i++ {
		own.head[1].Get(i).Update(data[i])
	}
	own.head[1].PrintToCoordinate(0, own.getLineCoordinate(idx))
}

// 结束表的使用，用于保证光标到达表的最底部
func (own *StatusForm) End() {
	base.PrintToCoordinate(0, own.y+3+2*own.num, "\n")
}

func (own *StatusForm) getLineCoordinate(idx int) int {
	if 0 > idx || idx > own.rowNum {
		return own.y + 3
	}
	return own.y + 3 + 2*idx
}

func (own *StatusForm) initFormString(rowNum int) string {
	var builder strings.Builder
	for i := 0; i < rowNum; i++ {
		builder.WriteString(own.head[1].String())
		builder.WriteString("\n")
		builder.WriteString(own.hSplitLine)
		builder.WriteString("\n")
	}
	return builder.String()
}
