package providers

import (
	"fmt"
	"os"
)

type DiskProvider struct{}

func (d *DiskProvider) SaveFile(filename string) (string, error) {
	path := fmt.Sprintf("uploads/%s", filename)
	_, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("file not exists")
	}
	return filename, nil
}

func (d *DiskProvider) DeleteFile(filename string) error {
	path := fmt.Sprintf("uploads/%s", filename)
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("file not exists")
	}
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("error deleting file")
	}
	return nil
}
