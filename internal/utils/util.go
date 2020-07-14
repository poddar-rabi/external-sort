package utils

import (
	"io"
	"io/ioutil"
	"os"
)

func DeleteFiles(files []string) error {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return err
		}
	}
	return nil
}

func IsDirExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func CreateTempFile(dir string) (*os.File, error) {
	return ioutil.TempFile(dir, "sorted_*_file.log")
}

func IsEmptyFile(filepath string) bool {
	fi, err := os.Stat(filepath)
	if err != nil {
		return true
	}
	return fi.Size() == 0
}
