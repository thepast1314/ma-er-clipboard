package entry

type EnumValue struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

func NewEnumValue(value int, label string) EnumValue {
	return EnumValue{Value: value, Label: label}
}
