package enums

type Code int

const (
	UNKNOWN = iota
	TEXT
	PICTURE
	FILE
	LINK
)

// 获取描述字符串
func (c Code) String() string {
	if label, ok := codeLabels[c]; ok {
		return label
	}
	return "Unknown"
}

var codeLabels = map[Code]string{
	TEXT:    "文本",
	PICTURE: "图片",
	FILE:    "文件",
	LINK:    "链接",
}

// 将 int 转换为 Code（带校验）
func ConvertEnum(i int) (Code, bool) {
	c := Code(i)
	if _, ok := codeLabels[c]; ok {
		return c, ok
	}
	return UNKNOWN, false // 或者新建一个 UNKNOWN
}

// 获取所有枚举值
func AllCodes() []Code {
	return []Code{TEXT, PICTURE, FILE, LINK}
}
