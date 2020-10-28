package file

import (
	"fmt"
	"io"
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

func Copy(source io.ReadSeeker, dest string) error {
	output, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		_ = output.Close()
	}()

	_, err = io.Copy(output, source)
	return err
}

func Move(source, dest string) error {
	return os.Rename(source, dest)
}
