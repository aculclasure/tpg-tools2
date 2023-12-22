package kv

type store struct {
	data map[string]string
}

func OpenStore(path string) (*store, error) {
	return &store{
		data: map[string]string{},
	}, nil
}

func (s store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}

func (s *store) Set(key, value string) {
	s.data[key] = value
}
