package controller

import (
	"cat-clipboard/backend/entry"
	"cat-clipboard/backend/enums"
)

type EnumController struct{}

func NewEnumController() *EnumController {
	return &EnumController{}
}

func (e *EnumController) QueryAllValueTypeEnum() []entry.EnumValue {
	valueTypeEnums := enums.AllCodes()

	var result = make([]entry.EnumValue, 0)

	for _, e := range valueTypeEnums {
		result = append(result, entry.NewEnumValue(int(e), e.String()))
	}

	return result
}
