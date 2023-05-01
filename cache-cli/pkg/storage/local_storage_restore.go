package storage

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func (s *LocalStorage) Restore(key string) (*os.File, error) {

	log.Infof("Restore Key %s :", key)

	sourceFilePath := fmt.Sprintf("%s/%s", s.Path, path.Base(key))
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()

	// Creating file in /tmp folder
	destinationFile, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf("%s-*", key))
	if err != nil {
		return nil, err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return nil, err
	}

	return destinationFile, nil
}
