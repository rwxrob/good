package good_test

import (
	"fmt"

	"github.com/rwxrob/good"
)

func ExampleReadBuildParams() {
	bp, err := good.ReadBuildParams(`testdata/build.yaml`)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bp.Targets[0])

	// Output:
	// {windows [amd64]}

}
