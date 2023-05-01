package storage

import "os"

func (s *LocalStorage) Usage() (*UsageSummary, error) {
	files, err := os.ReadDir(s.Path)
	if err != nil {
		return nil, err
	}

	var totalUsed int64
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}

		totalUsed = totalUsed + fileInfo.Size()
	}

	return &UsageSummary{
		Used: totalUsed,
		Free: s.Config().MaxSpace - totalUsed,
	}, nil
}
