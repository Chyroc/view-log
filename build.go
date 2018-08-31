// +build mage

package main

import (
	"os"
	"runtime"
	"strings"

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
	err := sh.RunV(goexe, "build", "-o", "bin/vlog", addOSExecType("github.com/Chyroc/vlog/cmd/vlog"))
	if err != nil {
		return err
	}
	return sh.RunV(goexe, "build", "-o", "bin/vlog-cli", addOSExecType("github.com/Chyroc/vlog/cmd/vlog-cli"))
}

func Fmt() error {
	goFiles, err := findGoFiles()
	if err != nil {
		return err
	}

	return sh.RunV("gofmt", append([]string{"-w"}, goFiles...)...)
}

func Import() error {
	goFiles, err := findGoFiles()
	if err != nil {
		return err
	}

	return sh.RunV("goimports", append([]string{"-w", "-local", "github.com/Chyroc/vlog"}, goFiles...)...)
}

func findGoFiles() ([]string, error) {
	output, err := sh.Output("find", ".", "-name", "*.go")
	if err != nil {
		return nil, err
	}
	var files []string
	for _, v := range strings.Split(output, "\n") {
		if !strings.HasPrefix(v, "./vendor") {
			files = append(files, v)
		}
	}

	return files, nil
}
