package player

import "github.com/nikiforosFreespirit/msdb5/display"

type info struct {
	name, host string
}

// IsSameHost func
func (info *info) IsSameHost(origin string) bool {
	return info.host == origin
}

// IsName func
func (info *info) IsName(name string) bool {
	return info.name == name
}

// IsNameEmpty func
func (info *info) IsNameEmpty() bool { return info.IsName("") }

// Info func
func (info *info) Info() []display.Info {
	name := display.NewInfo("Name", ":", info.name, ";")
	return display.Wrap("Info", name)
}

func (info *info) String() string {
	host := display.NewInfo("Host", ":", info.host, ";")
	name := display.NewInfo("Name", ":", info.name, ";")
	return display.All(display.Wrap("Info", host, name)...)
}
