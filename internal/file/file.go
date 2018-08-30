package file

import (
	"github.com/hpcloud/tail"
)

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
