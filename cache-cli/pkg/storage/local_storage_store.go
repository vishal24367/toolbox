package storage

import (
	"fmt"
	"time"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

func (s *LocalStorage) Store(key, localFilePath string) error {
	epochNanos := time.Now().UnixNano()
	tmpKey := fmt.Sprintf("%s-%d", os.Getenv("NEETO_CI_JOB_ID"), epochNanos)

	localFileInfo, err := os.Stat(localFilePath)
	if err != nil {
		return err
	}

	err = s.allocateSpace(localFileInfo.Size())
	if err != nil {
		return err
	}

	localFile, err := os.Open(localFilePath)
	if err != nil {
		return err
	}

	// create tmp file in s.Path directory
	remoteTmpFilePath := path.Join(s.Path, tmpKey)
	remoteTmpFile, err := os.Create(remoteTmpFilePath)
	if err != nil {
		_ = localFile.Close()
		return err
	}

	_, err = remoteTmpFile.ReadFrom(localFile)

	if err != nil {
		if rmErr := os.Remove(remoteTmpFilePath); rmErr != nil {
			log.Errorf("Error removing temporary file %s: %v", remoteTmpFilePath, rmErr)
		}

		_ = localFile.Close()
		_ = remoteTmpFile.Close()
		return err
	}

	err = os.Rename(remoteTmpFilePath, path.Join(s.Path, key))
	if err != nil {
		if rmErr := os.Remove(remoteTmpFilePath); rmErr != nil {
			log.Errorf("Error removing temporary file %s: %v", remoteTmpFilePath, rmErr)
		}

		_ = localFile.Close()
		_ = remoteTmpFile.Close()
		return err
	}

	err = remoteTmpFile.Close()
	if err != nil {
		_ = localFile.Close()
		return err
	}

	return localFile.Close()
}

func (s *LocalStorage) allocateSpace(space int64) error {
	usage, err := s.Usage()
	if err != nil {
		return err
	}

	freeSpace := usage.Free
	if freeSpace < space {
		fmt.Printf("Not enough space, deleting keys based on %s...\n", s.Config().SortKeysBy)
		keys, err := s.List()
		if err != nil {
			return err
		}

		for freeSpace < space {
			lastKey := keys[len(keys)-1]
			err = s.Delete(lastKey.Name)
			if err != nil {
				return err
			}

			log.Infof("Key '%s' is deleted.", lastKey.Name)
			freeSpace = freeSpace + lastKey.Size
			keys = keys[:len(keys)-1]
		}
	}

	return nil
}
