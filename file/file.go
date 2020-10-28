package file

import (
	"fmt"
	"os"
)

func CheckFile(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", path)
	}

	return nil
}

func Move(source, dest string) error {
	return os.Rename(source, dest)
}
