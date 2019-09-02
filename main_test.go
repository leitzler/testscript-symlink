package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(t *testing.T) {
	testscript.Run(t, testscript.Params{Dir: "scripts"})
}
