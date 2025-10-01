package file_manager

import (
	"fmt"
	"os"
)

func MakeDir(name string) (err error) {
	err = os.MkdirAll(name, os.ModePerm)
	if err != nil {
		return fmt.Errorf("make (%s) dir failed: %w", name, err)
	}
	return
}

func MakeFile(name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("make (%s) file failed: %w", name, err)
	}
	err = file.Close()
	if err != nil {
		return fmt.Errorf("close (%s) file failed: %w", name, err)
	}
	return
}

func Remove(name string) (err error) {
	err = os.RemoveAll(name)
	if err != nil {
		return fmt.Errorf("remove (%s) failed: %w", name, err)
	}
	return
}
