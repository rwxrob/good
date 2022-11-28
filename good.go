package good

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs"
	"github.com/rwxrob/fs/dir"
	"gopkg.in/yaml.v3"
)

var BuildDirName = `build`

type BuildTarget struct {
	OS   string   // go dist compatible os (GOOS)
	Arch []string // go dist compatible arch (GOARCH)
}

type BuildParams struct {

	// name of the binary to build
	Name string

	// relative directory to that containing build.yaml in which to build
	Dir string

	// target os and architectures to build
	Targets []BuildTarget

	// preserve anything extra in the yaml file when marshaling
	O map[string]any `yaml:",inline"`
}

func (bp BuildParams) YAML() string {
	buf, err := yaml.Marshal(bp)
	if err != nil {
		return "null"
	}
	return string(buf)
}

// ReadBuildParams looks in the path for build.yaml and returns
// a BuildParams from it.
func ReadBuildParams(path string) (*BuildParams, error) {

	if !strings.HasSuffix(path, `build.yaml`) {
		path = filepath.Join(path, `build.yaml`)
	}

	bp := new(BuildParams)
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf, bp); err != nil {
		return nil, err
	}
	return bp, nil
}

// Build looks for a build.yaml file and does a go build as specified
// normally producing a build directory within the same directory as the
// build.yaml file. If no Name is set the name of the directory
// containing the build.yaml file is be used. If Dir is not set then
// the package variable BuildDirName is used.
func Build(path string) error {

	p, err := ReadBuildParams(path)
	if err != nil {
		return err
	}

	if p.Dir == "" {
		p.Dir = BuildDirName
	}

	if p.Name == "" {
		p.Name = filepath.Base(path)
	}

	bdir := filepath.Join(path, p.Dir)

	pname, err := fs.Preserve(bdir)
	if err != nil {
		return err
	}
	defer fs.RevertIfMissing(bdir, pname)

	if err := dir.Create(bdir); err != nil {
		return err
	}

	odir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(odir)

	if err := os.Chdir(path); err != nil {
		return err
	}

	for _, target := range p.Targets {
		for _, arch := range target.Arch {
			log.Printf(`Building for %v/%v`, target.OS, arch)
			name := fmt.Sprintf(`%v_%v_%v`, p.Name, target.OS, arch)
			os.Setenv(`GOOS`, target.OS)
			os.Setenv(`GOARCH`, arch)
			if target.OS == `windows` {
				name += `.exe`
			}
			Z.Exec(`go`, `build`, `-o`, filepath.Join(p.Dir, name))
		}
	}

	return nil
}
