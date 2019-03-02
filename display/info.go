package display

// Info interface
type Info interface {
	Print() string
}

// InfoStruct struct
type InfoStruct struct {
	head, separator1, data, separator2 string
}

// PrintIt func
func (info InfoStruct) PrintIt() string {
	return info.head + info.separator1 + info.data + info.separator2
}

// NewInfo func
func NewInfo(info, sep1, field, sep2 string) InfoStruct {
	return InfoStruct{info, sep1, field, sep2}
}
