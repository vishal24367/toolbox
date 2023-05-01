package storage

import (
	"os"
)

func (s *LocalStorage) Delete(key string) error {
	err := os.Remove(key)
	if err != nil && os.IsNotExist(err) {
		return err
	}
	return nil
}
