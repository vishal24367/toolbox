package storage

import (
	"os"
	"path"
)

func (s *LocalStorage) HasKey(key string) (bool, error) {
	sourceFilePath := path.Join(s.Path, key)
	_, err := os.Stat(sourceFilePath)

	if os.IsNotExist(err) || err != nil {
		return false, err
	}

	return true, nil
}
