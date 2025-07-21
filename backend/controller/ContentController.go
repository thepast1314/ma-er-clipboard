package controller

import (
	"cat-clipboard/backend/entry"
	"cat-clipboard/backend/enums"
	"strings"
)

type ContentController struct{}

func NewContentController() *ContentController {
	return &ContentController{}
}

// 水果名称列表
var fruits = []string{
	"苹果", "香蕉", "橙子", "葡萄", "西瓜", "菠萝", "芒果", "草莓", "蓝莓", "樱桃",
	"哈密瓜", "柚子", "猕猴桃", "火龙果", "榴莲", "石榴", "杏子", "李子", "荔枝", "龙眼",
	"梨", "柿子", "椰子", "山竹", "桑葚",
}
var contents = []*entry.Content{
	entry.NewContent("苹果", enums.TEXT),
	entry.NewContent("香蕉", enums.PICTURE),
	entry.NewContent("橙子", enums.PICTURE),
	entry.NewContent("葡萄", enums.PICTURE),
	entry.NewContent("西瓜", enums.PICTURE),
	entry.NewContent("菠萝", enums.TEXT),
	entry.NewContent("芒果", enums.TEXT),
	entry.NewContent("草莓", enums.TEXT),
	entry.NewContent("蓝莓", enums.TEXT),
	entry.NewContent("樱桃", enums.LINK),
	entry.NewContent("哈密瓜", enums.LINK),
	entry.NewContent("柚子", enums.LINK),
	entry.NewContent("猕猴桃", enums.LINK),
	entry.NewContent("火龙果", enums.LINK),
	entry.NewContent("榴莲", enums.FILE),
	entry.NewContent("石榴", enums.FILE),
	entry.NewContent("杏子", enums.FILE),
	entry.NewContent("李子", enums.FILE),
	entry.NewContent("荔枝", enums.FILE),
	entry.NewContent("龙眼", enums.FILE),
	entry.NewContent("梨", enums.LINK),
	entry.NewContent("柿子", enums.PICTURE),
	entry.NewContent("椰子", enums.TEXT),
	entry.NewContent("山竹", enums.PICTURE),
	entry.NewContent("桑葚", enums.FILE),
}

func (c *ContentController) QueryAllContent() []*entry.Content {
	var size = len(fruits)

	// 创建长度为 50 的数组
	//var arr = make([]entry.Content, size)

	// 填充内容，循环使用水果名
	for i := 0; i < size; i++ {
		//arr[i] = *entry.NewContent(fruits[i], i) // 超过就循环使用水果名
	}
	return contents[:]
}

func (c *ContentController) QueryContentByValue(value string) []*entry.Content {
	//println(value)

	var result []*entry.Content

	for _, fruit := range contents {
		if strings.Contains(fruit.Value, value) {
			result = append(result, fruit)
		}
	}

	//for _, content := range result {
	//	fmt.Println(content)
	//}

	return result
}

func (c *ContentController) QueryContentByType(t int) []*entry.Content {
	var result []*entry.Content
	enum, ok := enums.ConvertEnum(t)
	if !ok {
		return result
	}

	for _, fruit := range contents {
		if fruit.Type == enum {
			result = append(result, fruit)
		}
	}

	return result
}

func (c *ContentController) QueryContentByCondition(request entry.ContentController) []*entry.Content {
	t := request.Type
	value := request.Value
	enum, ok := enums.ConvertEnum(t)

	var result []*entry.Content
	for _, fruit := range contents {
		if !ok && value == "" { // 都没值
			continue
		} else if ok && value != "" { // 都有值
			if fruit.Type == enum && strings.Contains(fruit.Value, value) {
				result = append(result, fruit)
			}
		} else if ok && value == "" { // 类型有值
			if fruit.Type == enum {
				result = append(result, fruit)
			}
		} else { // Value有值
			if strings.Contains(fruit.Value, value) {
				result = append(result, fruit)
			}
		}
	}
	return result
}

func (c ContentController) DeleteContentByPrimary(key string) {
	for i, fruit := range contents {
		if strings.Contains(key, fruit.Value) {
			contents = append(contents[:i], contents[i+1:]...)
			break
		}
	}
}
