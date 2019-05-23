package game

type Info struct {
	forAll, forMe string
	err           error
}

func NewInfo(forAll, forMe string, err error) *Info {
	return &Info{forAll, forMe, err}
}

func NewErrorInfo(err error) *Info {
	return NewInfo("", "", err)
}

func (info *Info) ToAll() string {
	return info.forAll
}

func (info *Info) ToMe() string {
	return info.forMe
}

func (info *Info) Err() error {
	return info.err
}
