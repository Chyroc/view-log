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
	err:= sh.RunV(goexe, "build", "-o", "bin/vlog", addOSExecType("github.com/Chyroc/vlog/cmd/vlog"))
	if err != nil {
		return err
	}
	return sh.RunV(goexe, "build", "-o", "bin/vlog-cli", addOSExecType("github.com/Chyroc/vlog/cmd/vlog-cli"))
}
