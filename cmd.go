package good

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `good`,
	Aliases:   []string{`go`},
	Version:   `v0.1.0`,
	Summary:   `go helper commands and tasks`,
	Copyright: `Copyright 2022 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands:  []*Z.Cmd{help.Cmd, buildCmd},
}

var buildCmd = &Z.Cmd{
	Name:     `build`,
	MaxArgs:  1,
	Summary:  `build using build.yaml params file`,
	Usage:    `(help|PATH)`,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The {{aka}} command builds from the build.yaml params file contained
		within the target path. The {{pre "name"}} is the main name of the
		executable to be built. It is derived from the name of the directory
		containing the build.yaml file if not set. The {{pre "targets"}}
		array contains one entry for each target operating system (GOOS)
		each with an array of {{pre "arch"}} supported architectures (GOARCH).

         targets:
           - os: windows
             arch:
               - amd64
           - os: linux
             arch:
               - amd64
           - os: darwin
             arch:
               - amd64
               - arm64

		Would produce a {{pre "build"}} directory (the default name) with
		the following files inside of it:

           build/foo_linux_amd64
           build/foo_darwin_arm64
           build/foo_windows_amd64.exe
           build/foo_darwin_amd64

		Note that the Windows executable has the {{pre ".exe"}} added. This is the only exception.

	`,

	Call: func(x *Z.Cmd, args ...string) error {
		if len(args) == 0 {
			args = append(args, `.`)
		}
		return Build(args[0])
	},
}
