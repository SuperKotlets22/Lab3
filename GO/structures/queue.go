package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Queue struct {
	Data     []string
	Front    int
	Rear     int
	Count    int
	Capacity int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		Data:     make([]string, cap),
		Capacity: cap,
		Front:    0,
		Rear:     -1,
		Count:    0,
	}
}

func (q *Queue) resize() {
	newCap := q.Capacity * 2
	newData := make([]string, newCap)
	for i := 0; i < q.Count; i++ {
		newData[i] = q.Data[(q.Front+i)%q.Capacity]
	}
	q.Data = newData
	q.Capacity = newCap
	q.Front = 0
	q.Rear = q.Count - 1
}

func (q *Queue) Push(val string) {
	if q.Count == q.Capacity {
		q.resize()
	}
	q.Rear = (q.Rear + 1) % q.Capacity
	q.Data[q.Rear] = val
	q.Count++
}

func (q *Queue) Pop() string {
	if q.Count == 0 {
		return ""
	}
	val := q.Data[q.Front]
	q.Front = (q.Front + 1) % q.Capacity
	q.Count--
	return val
}

func (q *Queue) Peek() string {
	if q.Count == 0 {
		return ""
	}
	return q.Data[q.Front]
}

// Сериализация (сохраняем логический порядок)
func (q *Queue) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, q.Count)
	for i := 0; i < q.Count; i++ {
		fmt.Fprintln(file, q.Data[(q.Front+i)%q.Capacity])
	}
	return nil
}

func (q *Queue) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var size int
	fmt.Fscanln(file, &size)
	// Reset
	q.Capacity = size * 2
	if q.Capacity == 0 {
		q.Capacity = 10
	}
	q.Data = make([]string, q.Capacity)
	q.Front, q.Rear, q.Count = 0, -1, 0
	for i := 0; i < size; i++ {
		var val string
		fmt.Fscanln(file, &val)
		q.Push(val)
	}
	return nil
}

func (q *Queue) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Собираем в слайс для простоты сохранения
	temp := make([]string, 0, q.Count)
	for i := 0; i < q.Count; i++ {
		temp = append(temp, q.Data[(q.Front+i)%q.Capacity])
	}
	return gob.NewEncoder(file).Encode(temp)
}

func (q *Queue) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var temp []string
	if err := gob.NewDecoder(file).Decode(&temp); err != nil {
		return err
	}
	q.Capacity = len(temp) * 2
	if q.Capacity == 0 {
		q.Capacity = 10
	}
	q.Data = make([]string, q.Capacity)
	q.Front, q.Rear, q.Count = 0, -1, 0
	for _, v := range temp {
		q.Push(v)
	}
	return nil
}