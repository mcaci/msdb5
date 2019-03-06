package display

// Info interface
type Info interface {
	Display() string
}

// InfoStruct struct
type InfoStruct struct {
	head, separator1, data, separator2 string
}

// Display func
func (info InfoStruct) Display() string {
	return info.head + info.separator1 + info.data + info.separator2
}

// NewInfo func
func NewInfo(info, sep1, field, sep2 string) Info {
	return InfoStruct{info, sep1, field, sep2}
}

// Wrap func
func Wrap(head string, infos ...Info) []Info {
	infoApis := make([]Info, 0)
	infoApis = append(infoApis, NewInfo("", "", head, "("))
	for _, info := range infos {
		infoApis = append(infoApis, info)
	}
	infoApis = append(infoApis, NewInfo("", "", "", ")"))
	return infoApis
}

// All func
func All(infos ...Info) (str string) {
	for _, info := range infos {
		str += info.Display()
	}
	return
}
