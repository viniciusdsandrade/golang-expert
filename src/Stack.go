package main

import (
	"fmt"
	"reflect"
)

type Stack struct {
	elemento       []interface{}
	tamanhoInicial int
	ultimo         int
}

func Constructor(tamanhoInicial int) *Stack {
	return &Stack{
		elemento:       make([]interface{}, tamanhoInicial),
		tamanhoInicial: tamanhoInicial,
		ultimo:         -1,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.ultimo == -1
}

func (s *Stack) IsFull() bool {
	return s.ultimo+1 == len(s.elemento)
}

func (s *Stack) Push(e interface{}) {
	if e == nil {
		return
	}

	if s.IsFull() {
		s.RedimencionarParaCima()
	}

	s.ultimo++
	s.elemento[s.ultimo] = e
}

func (s *Stack) Peek() (interface{}, error) {
	if s.ultimo < 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.elemento[s.ultimo], nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.ultimo < 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	e := s.elemento[s.ultimo]
	s.ultimo--

	if len(s.elemento) > s.tamanhoInicial &&
		s.ultimo+1 <= len(s.elemento)/4 {
		s.RedimencionarParaBaixo()
	}

	return e, nil
}

func (s *Stack) RedimencionarParaCima() {
	novo := make([]interface{}, len(s.elemento)*2)
	copy(novo, s.elemento)
	s.elemento = novo
}

func (s *Stack) RedimencionarParaBaixo() {
	novo := make([]interface{}, len(s.elemento)/2)
	copy(novo, s.elemento)
	s.elemento = novo
}

func CopyConstructor(modelo *Stack) *Stack {
	if modelo == nil {
		panic("Modelo ausente")
	}

	novo := &Stack{
		tamanhoInicial: modelo.tamanhoInicial,
		ultimo:         modelo.ultimo,
		elemento:       make([]interface{}, modelo.tamanhoInicial),
	}

	for i := 0; i < modelo.ultimo; i++ {
		novo.elemento[i] = modelo.elemento[i]
	}

	return novo
}

func (s *Stack) Clone() *Stack {
	clone := CopyConstructor(s)
	return clone
}

func (s *Stack) HashCode() int {
	const prime = 31
	hash := 1

	hash = hash*prime + s.tamanhoInicial
	hash = hash*prime + s.ultimo

	for i := 0; i < s.ultimo; i++ {
		hash = hash*prime + int(reflect.ValueOf(s.elemento[i]).Elem().Int())
	}

	if hash < 0 {
		hash = -hash
	}

	return hash
}

func (s *Stack) Equals(other *Stack) bool {
	if s == other {
		return true
	}
	if other == nil {
		return false
	}

	if s.tamanhoInicial != other.tamanhoInicial {
		return false
	}
	if s.ultimo != other.ultimo {
		return false
	}

	for i := 0; i < s.ultimo; i++ {
		if s.elemento[i] != other.elemento[i] {
			return false
		}
	}
	return true
}

func (s *Stack) ToString() string {
	if s.IsEmpty() {
		return "[]"
	}
	return "[" + fmt.Sprintf("%v", s.elemento[s.ultimo]) + "]"
}

func (s *Stack) ToArray() string {
	if s.ultimo < 0 {
		return "[]"
	}
	ret := "["
	for i := s.ultimo; i > 0; i-- {
		ret += fmt.Sprintf("%v, ", s.elemento[i])
	}
	ret += fmt.Sprintf("%v]", s.elemento[0])
	return ret
}
