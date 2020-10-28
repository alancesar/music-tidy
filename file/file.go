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

func Copy(dest string, source io.ReadSeeker) (int64, error) {
	output, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = output.Close()
	}()

	return io.Copy(output, source)
}
