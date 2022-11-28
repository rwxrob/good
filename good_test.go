package good_test

import (
	"fmt"

	"github.com/rwxrob/good"
)

func ExampleReadBuildParams() {
	bp, err := good.ReadBuildParams(`testdata/foo`)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bp.Targets[0])

	// Output:
	// {windows [amd64]}

}

/*
func ExampleBuild() {

	// as file path
	if err := good.Build(`testdata/foo`); err != nil {
		fmt.Println(err)
	}

	// Output:
	// ignored

}
*/
