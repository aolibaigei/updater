package main

import (
	"fmt"

	"github.com/Masterminds/semver"
)

func main() {
	c, err := semver.NewConstraint(">=1.2.3-beta.12+build345")
	if err != nil {
		// Handle constraint not being parseable.
	}

	v, _ := semver.NewVersion("1.2.3-beta.1+build777")
	if err != nil {
		// Handle version not being parseable.
	}
	// Check if the version meets the constraints. The a variable will be true.
	a := c.Check(v)
	//true
	fmt.Println(a)
}
