package storage

import "os"

func (s *SFTPStorage) Usage() (*UsageSummary, error) {
	files, err := s.SFTPClient.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var totalUsed int64
	for _, file := range files {
		if file.IsDir() || file.Mode()&os.ModeSymlink != 0 {
			continue
		}
		totalUsed += file.Size()
	}

	return &UsageSummary{
		Used: totalUsed,
		Free: s.Config().MaxSpace - totalUsed,
	}, nil
}
