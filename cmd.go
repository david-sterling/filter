package filter

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "filter",
	Aliases:  []string{`filt`},
	Summary:  "Filter text for URLs and other patterns",
	Commands: []*Z.Cmd{help.Cmd},
}
