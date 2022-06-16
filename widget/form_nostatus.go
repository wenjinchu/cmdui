package widget

import (
	"cmdui/base"
	"fmt"
	"strings"
)

// 定义无状态表的类型
// 直接打印无任何状态变化的数据，无数据量限制
type NoStatusForm struct {
	*form
}

// 新建一个无状态表
func NewNoStatusForm(fieldWidths ...int) *NoStatusForm {
	return &NoStatusForm{newForm(fieldWidths...)}
}

// 添加一个表字段
func (own *NoStatusForm) AddField(field *base.Element) {
	own.addField(field)
}

// 实时打印表数据
func (own *NoStatusForm) Print(data []interface{}) {
	if !own.isPrinting {
		fmt.Println(own.hSplitLine)
		fmt.Println(own.headString)
		fmt.Println(own.hSplitLine)
	}
	fmt.Println(own.formatData(own.dealData(data)))
	own.isPrinting = true
}

// 格式化表数据 - 根据每个字段的隐藏字段处理
func (own *NoStatusForm) formatData(data []interface{}) string {
	var builder strings.Builder
	for i, elem := range own.head[1].elems {
		elem.Update(data[i])
	}
	builder.WriteString(own.head[1].String())
	builder.WriteString("\n")
	builder.WriteString(own.hSplitLine)

	return builder.String()
}
