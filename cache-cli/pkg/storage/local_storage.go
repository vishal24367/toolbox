package storage

type LocalStorage struct {
	Path          string
	StorageConfig StorageConfig
}

type LocalStorageOptions struct {
	Path   string
	Config StorageConfig
}

func NewLocalStorage(options LocalStorageOptions) (*LocalStorage, error) {

	storage := LocalStorage{
		Path:          options.Path,
		StorageConfig: options.Config,
	}

	return &storage, nil
}

func (s *LocalStorage) Config() StorageConfig {
	return s.StorageConfig
}
