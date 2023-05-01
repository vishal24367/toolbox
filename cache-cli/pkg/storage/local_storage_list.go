package storage

import (
	"os"
	"sort"
)

func (s *LocalStorage) List() ([]CacheKey, error) {

	// Open the directory
	dir, err := os.Open(s.Path)
	if err != nil {
		return []CacheKey{}, err
	}

	defer dir.Close()

	// Read the contents of the directory
	files, err := dir.Readdir(0)
	if err != nil {
		return []CacheKey{}, err
	}

	keys := []CacheKey{}
	for _, file := range files {
		storedAt := file.ModTime()
		keys = append(keys, CacheKey{
			Name:           file.Name(),
			Size:           file.Size(),
			StoredAt:       &storedAt,
			LastAccessedAt: findLastAccessedAt(file),
		})
	}
	return s.sortKeys(keys), nil
}

func (s *LocalStorage) sortKeys(keys []CacheKey) []CacheKey {
	switch s.Config().SortKeysBy {
	case SortBySize:
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i].Size > keys[j].Size
		})
	case SortByAccessTime:
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i].LastAccessedAt.After(*keys[j].LastAccessedAt)
		})
	default:
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i].StoredAt.After(*keys[j].StoredAt)
		})
	}

	return keys
}
