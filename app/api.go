package app

// Action interface
type Action interface {
	Process(request, origin string) *Info
}

type Info struct {
	forAll, forMe string
	err           error
}

func NewInfo(forAll, forMe string, err error) *Info {
	return &Info{forAll, forMe, err}
}

func (info *Info) ForAll() string {
	return info.forAll
}

func (info *Info) ForMe() string {
	return info.forMe
}

func (info *Info) Err() error {
	return info.err
}
