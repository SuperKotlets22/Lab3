package structures

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"os"
)

// динамический массив
type DynamicArray struct {
	Data []string
}

// конструктор
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{Data: []string{}}
}

func (d *DynamicArray) PushBack(val string) {
	d.Data = append(d.Data, val)
}

// ИЗМЕНЕНИЕ: Возвращаем (string, error)
func (d *DynamicArray) Get(i int) (string, error) {
	if i < 0 || i >= len(d.Data) {
		return "", errors.New("index out of bounds")
	}
	return d.Data[i], nil
}

// json сериализация
func (d *DynamicArray) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(d.Data)
}

func (d *DynamicArray) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&d.Data)
}

// gob сериализация
func (d *DynamicArray) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(d.Data)
}

func (d *DynamicArray) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewDecoder(file).Decode(&d.Data)
}