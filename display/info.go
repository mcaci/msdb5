package display

import "fmt"

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

// PrintAll func
func PrintAll(infos ...InfoStruct) (str string) {
	for _, info := range infos {
		str += info.PrintIt()
	}
	return
}

// ToString func
func ToString(infos ...fmt.Stringer) string {
	var str string
	for _, info := range infos {
		str += info.String() + " "
	}
	return str
}
