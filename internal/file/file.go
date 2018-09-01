package file

import (
	"io/ioutil"
	"os"
	"sort"

	"github.com/hpcloud/tail"
)

func ListFileInfos(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool { return files[i].ModTime().Sub(files[j].ModTime()) < 0 })

	return files, err
}

func Tail(file string, callback func(line string) error) error {
	t, err := tail.TailFile(file, tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for line := range t.Lines {
		if line.Err != nil {
			return line.Err
		}
		if err = callback(line.Text); err != nil {
			return err
		}
	}
	return nil
}
