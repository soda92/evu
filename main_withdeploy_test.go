//go:build withdeploy

package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestWithdeploy(t *testing.T) {
	p := commonTestScriptsParam
	p.Dir = "testscripts/withdeploy"
	testscript.Run(t, p)
}
