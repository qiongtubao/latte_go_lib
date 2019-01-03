package fs

import (
	"io/ioutil"
	"os"
	"path"
)

func WriteFile(fileName string, data string) error {
	error := os.MkdirAll(path.Dir(fileName), 0700)
	if error != nil {
		return error
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(data)
	return nil
}

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	return string(fd), err
}
