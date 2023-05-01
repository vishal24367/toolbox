package storage

import (
	"os"
	"path"
)

func (s *LocalStorage) Delete(key string) error {
	sourceFilePath := path.Join(s.Path, key)
	err := os.Remove(sourceFilePath)

	if err != nil && os.IsNotExist(err) {
		return err
	}

	return nil
}
