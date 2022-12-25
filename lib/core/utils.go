package core

import "os"

type StrArray []string

func (i StrArray) Has(val string) bool {
	for _, v := range i {
		if v == val {
			return true
		}
	}
	return false
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}
