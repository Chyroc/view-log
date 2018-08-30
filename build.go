// +build mage

package main

import (
	"os"
	"runtime"

	"github.com/magefile/mage/sh"
)

var goexe = "go"

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}
}

func addOSExecType(str string) string {
	if runtime.GOOS == "windows" {
		return str + ".exe"
	}

	return str
}

func Build() error {
	return sh.RunV(goexe, "build", "-o", "bin/view-log", addOSExecType("github.com/Chyroc/view-log/cmd/view-log"))
}
