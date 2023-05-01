package storage

import (
	"fmt"
	"io"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

func (s *LocalStorage) Store(key, sourceFilePath string) error {

	log.Infof("Key %s and SourceFilePath %s.", key, sourceFilePath)

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	fileSize, _ := sourceFile.Stat()
	err = s.allocateSpace(fileSize.Size())
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", s.Path, path.Base(key))

	log.Infof("Destination FilePath %s.", filePath)
	destinationFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
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
