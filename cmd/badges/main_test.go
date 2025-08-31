package main

import (
	"os"
	"testing"
)

func TestBadgesMain(t *testing.T) {
	// a simple test case that runs the command with some default values
	// and checks if the output file is created.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"badges", "-readme=README.md", "-test-status=Passing", "-coverage=90", "-race-status=Clean", "-vet-status=OK", "-license=MIT"}

	main()

	// check if the file exists
	if _, err := os.Stat("docs/img/badges.svg"); os.IsNotExist(err) {
		t.Fatal("Expected file to be created, but it was not")
	}

	// remove the file
	os.Remove("docs/img/badges.svg")
	os.Remove("README.md")
}
