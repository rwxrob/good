package good

import (
	"os"

	"gopkg.in/yaml.v3"
)

type BuildTarget struct {
	OS   string
	Arch []string
}

type BuildParams struct {
	Targets []BuildTarget
	O       map[string]any `yaml:",inline"`
}

func (bp BuildParams) YAML() string {
	buf, err := yaml.Marshal(bp)
	if err != nil {
		return "null"
	}
	return string(buf)
}

func ReadBuildParams(file string) (*BuildParams, error) {
	bp := new(BuildParams)
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf, bp); err != nil {
		return nil, err
	}
	return bp, nil
}
