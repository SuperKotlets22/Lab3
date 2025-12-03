package structures

import (
	"encoding/gob"
	"fmt"
	"os"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type CompleteBinaryTree struct {
	Root *TreeNode
	Size int
}

func (t *CompleteBinaryTree) Insert(val int) {
	newNode := &TreeNode{Data: val}
	if t.Root == nil {
		t.Root = newNode
		t.Size++
		return
	}
	// BFS
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Left == nil {
			curr.Left = newNode
			t.Size++
			return
		} else {
			queue = append(queue, curr.Left)
		}

		if curr.Right == nil {
			curr.Right = newNode
			t.Size++
			return
		} else {
			queue = append(queue, curr.Right)
		}
	}
}

func (t *CompleteBinaryTree) Find(val int) bool {
	if t.Root == nil {
		return false
	}
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Data == val {
			return true
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return false
}

func (t *CompleteBinaryTree) IsComplete() bool {
	if t.Root == nil {
		return true
	}
	queue := []*TreeNode{t.Root}
	flag := false
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Left != nil {
			if flag {
				return false
			}
			queue = append(queue, curr.Left)
		} else {
			flag = true
		}

		if curr.Right != nil {
			if flag {
				return false
			}
			queue = append(queue, curr.Right)
		} else {
			flag = true
		}
	}
	return true
}

// Сериализация (Level Order)
func (t *CompleteBinaryTree) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if t.Root == nil {
		return nil
	}
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Fprintln(file, curr.Data)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return nil
}

func (t *CompleteBinaryTree) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	t.Root = nil
	t.Size = 0
	var val int
	for {
		_, err := fmt.Fscanln(file, &val)
		if err != nil {
			break
		}
		t.Insert(val)
	}
	return nil
}

func (t *CompleteBinaryTree) SaveToBinary(filename string) error {
	// Собираем слайс значений
	var vals []int
	if t.Root != nil {
		queue := []*TreeNode{t.Root}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			vals = append(vals, curr.Data)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return gob.NewEncoder(file).Encode(vals)
}

func (t *CompleteBinaryTree) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var vals []int
	if err := gob.NewDecoder(file).Decode(&vals); err != nil {
		return err
	}
	t.Root = nil
	t.Size = 0
	for _, v := range vals {
		t.Insert(v)
	}
	return nil
}