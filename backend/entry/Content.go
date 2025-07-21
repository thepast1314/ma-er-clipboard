package entry

import "cat-clipboard/backend/enums"

type Content struct {
	Value string     `json:"value"`
	Type  enums.Code `json:"type"`
}

func NewContent(value string, t enums.Code) *Content {
	return &Content{Value: value, Type: t}
}
