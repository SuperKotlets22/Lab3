package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

// --- SINGLY LINKED LIST ---

type SNode struct {
	Data string
	Next *SNode
}

type SinglyList struct {
	Head *SNode
	Tail *SNode
	Size int
}

func (l *SinglyList) PushBack(val string) {
	newNode := &SNode{Data: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *SinglyList) PushFront(val string) {
	newNode := &SNode{Data: val, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
	l.Size++
}

func (l *SinglyList) InsertAfter(target, val string) {
	curr := l.Head
	for curr != nil {
		if curr.Data == target {
			newNode := &SNode{Data: val, Next: curr.Next}
			curr.Next = newNode
			if curr == l.Tail {
				l.Tail = newNode
			}
			l.Size++
			return
		}
		curr = curr.Next
	}
}

func (l *SinglyList) InsertBefore(target, val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == target {
		l.PushFront(val)
		return
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Data == target {
			newNode := &SNode{Data: val, Next: curr.Next}
			curr.Next = newNode
			l.Size++
			return
		}
		curr = curr.Next
	}
}

func (l *SinglyList) PopFront() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	l.Size--
}

func (l *SinglyList) PopBack() {
	if l.Head == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		curr := l.Head
		for curr.Next != l.Tail {
			curr = curr.Next
		}
		curr.Next = nil
		l.Tail = curr
	}
	l.Size--
}

func (l *SinglyList) RemoveByValue(val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == val {
		l.PopFront()
		return
	}
	curr := l.Head
	for curr.Next != nil {
		if curr.Next.Data == val {
			if curr.Next == l.Tail {
				l.Tail = curr
			}
			curr.Next = curr.Next.Next
			l.Size--
			return
		}
		curr = curr.Next
	}
}

func (l *SinglyList) Search(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Сериализация для SinglyList
func (l *SinglyList) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	curr := l.Head
	for curr != nil {
		fmt.Fprintln(file, curr.Data)
		curr = curr.Next
	}
	return nil
}

func (l *SinglyList) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	l.Head = nil
	l.Tail = nil
	l.Size = 0
	var val string
	for {
		_, err := fmt.Fscanln(file, &val)
		if err != nil {
			break
		}
		l.PushBack(val)
	}
	return nil
}

func (l *SinglyList) SaveToBinary(filename string) error {
	var data []string
	curr := l.Head
	for curr != nil {
		data = append(data, curr.Data)
		curr = curr.Next
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(data)
}

func (l *SinglyList) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	if err := gob.NewDecoder(file).Decode(&data); err != nil {
		return err
	}
	l.Head, l.Tail, l.Size = nil, nil, 0
	for _, v := range data {
		l.PushBack(v)
	}
	return nil
}

// --- DOUBLY LINKED LIST ---

type DNode struct {
	Data string
	Next *DNode
	Prev *DNode
}

type DoublyList struct {
	Head *DNode
	Tail *DNode
	Size int
}

func (l *DoublyList) PushBack(val string) {
	newNode := &DNode{Data: val}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
	l.Size++
}

func (l *DoublyList) PushFront(val string) {
	newNode := &DNode{Data: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Size++
}

func (l *DoublyList) PopFront() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Size--
}

func (l *DoublyList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Size--
}

func (l *DoublyList) Search(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Сериализация (аналогично SinglyList, используем промежуточный slice для бинарки)
func (l *DoublyList) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	curr := l.Head
	for curr != nil {
		fmt.Fprintln(file, curr.Data)
		curr = curr.Next
	}
	return nil
}

func (l *DoublyList) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	l.Head, l.Tail, l.Size = nil, nil, 0
	var val string
	for {
		_, err := fmt.Fscanln(file, &val)
		if err != nil {
			break
		}
		l.PushBack(val)
	}
	return nil
}

func (l *DoublyList) SaveToBinary(filename string) error {
	var data []string
	curr := l.Head
	for curr != nil {
		data = append(data, curr.Data)
		curr = curr.Next
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(data)
}

func (l *DoublyList) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	if err := gob.NewDecoder(file).Decode(&data); err != nil {
		return err
	}
	l.Head, l.Tail, l.Size = nil, nil, 0
	for _, v := range data {
		l.PushBack(v)
	}
	return nil
}