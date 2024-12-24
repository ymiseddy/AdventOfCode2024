package set

import (
	"fmt"
	"iter"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func FromItem[T comparable](element T) *Set[T] {
	result := NewSet[T]()
	result.Add(element)
	return result
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: map[T]struct{}{}}
}

func (s *Set[T]) String() string {
	str := ""
	for element, _ := range s.elements {
		str += fmt.Sprintf("%v, ", element)

	}
	return str
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) AddAll(elements []T) {
	for _, element := range elements {
		s.Add(element)
	}
}

func (s *Set[T]) Contains(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for element, _ := range s.elements {
			if !yield(element) {
				return
			}
		}
	}
}

func Union[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element, _ := range a.elements {
		result.Add(element)
	}
	for element, _ := range b.elements {
		result.Add(element)
	}
	return result
}

func (s *Set[T]) Clone() *Set[T] {
	result := NewSet[T]()
	for element, _ := range s.elements {
		result.Add(element)
	}
	return result
}

func (s *Set[T]) ToSlice() []T {
	var result []T
	for element, _ := range s.elements {
		result = append(result, element)
	}
	return result
}

func (s *Set[T]) Without(element T) *Set[T] {
	result := NewSet[T]()
	for ele, _ := range s.elements {
		if ele != element {
			result.Add(ele)
		}
	}
	return result
}

func (s *Set[T]) With(element T) *Set[T] {
	result := FromItem[T](element)

	for ele, _ := range s.elements {
		result.Add(ele)
	}
	return result
}

func Intersection[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element, _ := range a.elements {
		if b.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func Difference[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element, _ := range a.elements {
		if !b.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func SymmetricDifference[T comparable](a, b *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element, _ := range a.elements {
		if !b.Contains(element) {
			result.Add(element)
		}
	}
	for element, _ := range b.elements {
		if !a.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
