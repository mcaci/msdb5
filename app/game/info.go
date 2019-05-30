package game

type Info struct {
	destination string
	message     string
	err         error
}

func NewInfo(destination, message string, err error) *Info {
	return &Info{destination, message, err}
}

func NewErrorInfo(err error) []*Info {
	return []*Info{&Info{"", "", err}}
}

func (info *Info) Dest() string {
	return info.destination
}

func (info *Info) Msg() string {
	return info.message
}

func (info *Info) Err() error {
	return info.err
}
