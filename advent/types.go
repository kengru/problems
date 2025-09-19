package advent

import "fmt"

type Stack[V any] struct {
	Length int
	Items  []V
}

type Queue[T any] struct {
	Length int
	Items  []T
}

func (s Stack[V]) Peek() V {
	return s.Items[s.Length-1]
}

func (q Queue[T]) Peek() T {
	return q.Items[q.Length-1]
}

func (s *Stack[V]) Pop() V {
	item := s.Items[s.Length-1]
	s.Items = s.Items[:s.Length-1]
	s.Length--
	return item
}

func (s *Stack[V]) Add(item V) {
	s.Items = append(s.Items, item)
	s.Length++
}

func (s *Stack[V]) Print() {
	if s.Length == 0 {
		fmt.Println("empty")
		return
	}
	idx := s.Length - 1
	fmt.Println("Stack")
	for idx >= 0 {
		fmt.Println("[", s.Items[idx], "]")
		idx--
	}
}

type Directory struct {
	Parent *Directory
	Name   string
	Size   int
}
