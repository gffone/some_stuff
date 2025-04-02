package tasks

type Storage struct {
	data           map[int]int
	uniqueElements map[int]struct{}
}

func NewStorage() *Storage {
	return &Storage{
		data:           make(map[int]int),
		uniqueElements: make(map[int]struct{}),
	}
}

func (s *Storage) Add(element int) {
	if _, exists := s.data[element]; exists {
		s.data[element]++
		delete(s.uniqueElements, element)
	} else {
		s.data[element] = 1
		s.uniqueElements[element] = struct{}{}
	}
}

func (s *Storage) Delete(element int) {
	if _, exists := s.uniqueElements[element]; exists {
		delete(s.uniqueElements, element)
	}

	if nums, exists := s.data[element]; exists {
		if nums == 1 {
			delete(s.data, element)
		} else {
			s.data[element]--
		}
	}

	if s.data[element] == 1 {
		s.uniqueElements[element] = struct{}{}
	}
}

func (s *Storage) GetUnique() (int, bool) {

	for element, _ := range s.uniqueElements {
		return element, true
	}

	return 0, false
}
