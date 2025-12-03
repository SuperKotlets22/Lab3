package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

type HashNode struct {
	Key   string
	Value string
	Next  *HashNode
}

type HashTable struct {
	Table    []*HashNode
	Capacity int
	Size     int
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{
		Table:    make([]*HashNode, cap),
		Capacity: cap,
		Size:     0,
	}
}

func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	for _, c := range key {
		hash = (hash*31 + int(c)) % ht.Capacity
	}
	return hash
}

func (ht *HashTable) Put(key, value string) {
	idx := ht.hashFunction(key)
	curr := ht.Table[idx]
	for curr != nil {
		if curr.Key == key {
			curr.Value = value
			return
		}
		curr = curr.Next
	}
	newNode := &HashNode{Key: key, Value: value, Next: ht.Table[idx]}
	ht.Table[idx] = newNode
	ht.Size++
}

func (ht *HashTable) Get(key string) string {
	idx := ht.hashFunction(key)
	curr := ht.Table[idx]
	for curr != nil {
		if curr.Key == key {
			return curr.Value
		}
		curr = curr.Next
	}
	return ""
}

func (ht *HashTable) Remove(key string) {
	idx := ht.hashFunction(key)
	curr := ht.Table[idx]
	var prev *HashNode = nil

	for curr != nil {
		if curr.Key == key {
			if prev == nil {
				ht.Table[idx] = curr.Next
			} else {
				prev.Next = curr.Next
			}
			ht.Size--
			return
		}
		prev = curr
		curr = curr.Next
	}
}

// Сериализация (Пары Ключ-Значение)
func (ht *HashTable) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, ht.Size)
	for _, head := range ht.Table {
		curr := head
		for curr != nil {
			fmt.Fprintln(file, curr.Key, curr.Value)
			curr = curr.Next
		}
	}
	return nil
}

func (ht *HashTable) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var size int
	fmt.Fscanln(file, &size)
	// Clear
	ht.Table = make([]*HashNode, ht.Capacity)
	ht.Size = 0
	for i := 0; i < size; i++ {
		var k, v string
		fmt.Fscanln(file, &k, &v)
		ht.Put(k, v)
	}
	return nil
}

// Для бинарной сериализации используем map для удобства gob
func (ht *HashTable) SaveToBinary(filename string) error {
	data := make(map[string]string)
	for _, head := range ht.Table {
		curr := head
		for curr != nil {
			data[curr.Key] = curr.Value
			curr = curr.Next
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(data)
}

func (ht *HashTable) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data map[string]string
	if err := gob.NewDecoder(file).Decode(&data); err != nil {
		return err
	}
	ht.Table = make([]*HashNode, ht.Capacity)
	ht.Size = 0
	for k, v := range data {
		ht.Put(k, v)
	}
	return nil
}