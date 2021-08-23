package lexer

type set struct {
	data map[string]struct{}
}

func newSet() *set {
	return &set{data: map[string]struct{}{}}
}

func (s *set) contains(in string) bool {
	_, ok := s.data[in]
	return ok
}

func (s *set) add(in string) {
	s.data[in] = struct{}{}
}