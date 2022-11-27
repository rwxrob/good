package good

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `good`,
	Aliases:  []string{`go`},
	Summary:  `go helper commands and tasks`,
	Commands: []*Z.Cmd{help.Cmd},
}
