package storage

import (
	"os"
	"path"
)

func (s *LocalStorage) Clear() error {
	keys, err := s.List()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	for _, key := range keys {
		sourceFilePath := path.Join(s.Path, key.Name)
		err := os.Remove(sourceFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}
