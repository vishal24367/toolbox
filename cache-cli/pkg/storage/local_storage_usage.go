package storage

import "os"

func (s *LocalStorage) Usage() (*UsageSummary, error) {

	// Open the directory
	dir, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}

	defer dir.Close()

	// Read the contents of the directory
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var totalUsed int64
	for _, file := range files {
		totalUsed = totalUsed + file.Size()
	}

	return &UsageSummary{
		Used: totalUsed,
		Free: s.Config().MaxSpace - totalUsed,
	}, nil
}
