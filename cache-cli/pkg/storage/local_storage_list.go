package storage

import (
	"io/fs"
	"os"
	"sort"
	"time"

	"github.com/pkg/sftp"
)

func (s *LocalStorage) List() ([]CacheKey, error) {
	files, err := os.ReadDir(s.Path)
	if err != nil {
		return nil, err
	}

	keys := []CacheKey{}
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}

		storedAt := fileInfo.ModTime()
		keys = append(keys, CacheKey{
			Name:           file.Name(),
			Size:           fileInfo.Size(),
			StoredAt:       &storedAt,
			LastAccessedAt: findLastAccessedTime(fileInfo),
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

// If we can't figure out the access time of the file,
// we fallback to the modification time.
func findLastAccessedTime(fileInfo fs.FileInfo) *time.Time {
	mtime := fileInfo.ModTime()

	// Try to get the underlying data source; if nil, fallback to mtime.
	ds := fileInfo.Sys()
	if ds == nil {
		return &mtime
	}

	// Try to cast the underlying data source to something we understand; if nil, fallback to mtime.
	stat, ok := ds.(*sftp.FileStat)
	if !ok {
		return &mtime
	}

	// atime can also be unset; fallback to mtime.
	if stat.Atime == 0 {
		mtime := fileInfo.ModTime()
		return &mtime
	}

	// Otherwise, we use atime.
	atime := time.Unix(int64(stat.Atime), 0)
	return &atime
}
