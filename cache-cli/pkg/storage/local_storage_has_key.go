package storage

import (
	"os"
)

func (s *LocalStorage) HasKey(key string) (bool, error) {
	_, err := os.Stat(key)
	if os.IsNotExist(err) || err != nil {
		return false, err
	}
	return true, nil
}
