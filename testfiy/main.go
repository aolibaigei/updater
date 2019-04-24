package main

import (
	"fmt"

	"github.com/Masterminds/semver"
)

func main() {
	c, err := semver.NewConstraint(">=4.8.0-56.61-16.04.1")
	if err != nil {
		// Handle constraint not being parseable.
	}

	v, _ := semver.NewVersion("4.8.0-56.69-16.04.1")
	if err != nil {
		// Handle version not being parseable.
	}
	// Check if the version meets the constraints. The a variable will be true.
	a := c.Check(v)
	//true
	fmt.Println(a)
}
