package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

type DynamicArray struct {
	Data []string
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{Data: make([]string, 0, 10)}
}

func (da *DynamicArray) Size() int {
	return len(da.Data)
}

func (da *DynamicArray) PushBack(val string) {
	da.Data = append(da.Data, val)
}

func (da *DynamicArray) PushFront(val string) {
	da.Data = append([]string{val}, da.Data...)
}

func (da *DynamicArray) InsertAt(index int, val string) {
	if index < 0 || index > len(da.Data) {
		panic("Index out of range")
	}
	if index == len(da.Data) {
		da.PushBack(val)
		return
	}
	da.Data = append(da.Data[:index+1], da.Data[index:]...)
	da.Data[index] = val
}

func (da *DynamicArray) PopBack() {
	if len(da.Data) > 0 {
		da.Data = da.Data[:len(da.Data)-1]
	}
}

func (da *DynamicArray) PopFront() {
	if len(da.Data) > 0 {
		da.Data = da.Data[1:]
	}
}

func (da *DynamicArray) RemoveAt(index int) {
	if index < 0 || index >= len(da.Data) {
		panic("Index out of range")
	}
	da.Data = append(da.Data[:index], da.Data[index+1:]...)
}

func (da *DynamicArray) Get(index int) string {
	if index < 0 || index >= len(da.Data) {
		panic("Index out of range")
	}
	return da.Data[index]
}

func (da *DynamicArray) Set(index int, val string) {
	if index < 0 || index >= len(da.Data) {
		panic("Index out of range")
	}
	da.Data[index] = val
}

func (da *DynamicArray) Find(val string) int {
	for i, v := range da.Data {
		if v == val {
			return i
		}
	}
	return -1
}

// Сериализация (Текст)
func (da *DynamicArray) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, len(da.Data))
	for _, v := range da.Data {
		fmt.Fprintln(file, v)
	}
	return nil
}

func (da *DynamicArray) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var size int
	fmt.Fscanln(file, &size)
	da.Data = make([]string, 0, size)

	for i := 0; i < size; i++ {
		var val string
		fmt.Fscanln(file, &val)
		da.Data = append(da.Data, val)
	}
	return nil
}

// Сериализация (Бинарная - используем gob)
func (da *DynamicArray) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(da.Data)
}

func (da *DynamicArray) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&da.Data)
}