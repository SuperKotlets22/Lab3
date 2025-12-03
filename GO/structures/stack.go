package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Stack struct {
	Data []string
}

func (s *Stack) Push(val string) {
	s.Data = append(s.Data, val)
}

func (s *Stack) Pop() string {
	if len(s.Data) == 0 {
		return ""
	}
	val := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return val
}

func (s *Stack) Peek() string {
	if len(s.Data) == 0 {
		return ""
	}
	return s.Data[len(s.Data)-1]
}

func (s *Stack) Size() int {
	return len(s.Data)
}

// Сериализация
func (s *Stack) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, len(s.Data))
	for _, v := range s.Data {
		fmt.Fprintln(file, v)
	}
	return nil
}

func (s *Stack) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var size int
	fmt.Fscanln(file, &size)
	s.Data = make([]string, 0, size)
	for i := 0; i < size; i++ {
		var val string
		fmt.Fscanln(file, &val)
		s.Data = append(s.Data, val)
	}
	return nil
}

func (s *Stack) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(s.Data)
}

func (s *Stack) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewDecoder(file).Decode(&s.Data)
}