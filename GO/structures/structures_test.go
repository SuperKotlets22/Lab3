package structures

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper to cleanup files
func cleanup(filenames ...string) {
	for _, f := range filenames {
		os.Remove(f)
	}
}

// ==========================================
// 1. DYNAMIC ARRAY
// ==========================================
func TestDynamicArray_Full(t *testing.T) {
	da := NewDynamicArray()

	// Basic Ops
	da.PushBack("A")
	da.PushFront("B") // B, A
	assert.Equal(t, "B", da.Get(0))
	assert.Equal(t, 2, da.Size())

	da.InsertAt(1, "C") // B, C, A
	assert.Equal(t, "C", da.Get(1))

	da.Set(1, "D") // B, D, A
	assert.Equal(t, "D", da.Get(1))

	assert.Equal(t, 2, da.Find("A"))
	assert.Equal(t, -1, da.Find("Missing"))

	da.RemoveAt(1) // B, A
	assert.Equal(t, "A", da.Get(1))

	da.PopFront() // A
	assert.Equal(t, "A", da.Get(0))

	da.PopBack() // Empty
	assert.Equal(t, 0, da.Size())

	// Safe pop on empty
	da.PopBack()
	da.PopFront()
	assert.Equal(t, 0, da.Size())

	// Resize logic coverage
	for i := 0; i < 20; i++ {
		da.PushBack(fmt.Sprintf("%d", i))
	}
	assert.Equal(t, 20, da.Size())

	// InsertAt edge cases
	da.InsertAt(20, "End")
	assert.Equal(t, "End", da.Get(20))

	// Panics check
	assert.Panics(t, func() { da.Get(-1) })
	assert.Panics(t, func() { da.Get(100) })
	assert.Panics(t, func() { da.Set(-1, "X") })
	assert.Panics(t, func() { da.Set(100, "X") })
	assert.Panics(t, func() { da.InsertAt(-1, "X") })
	assert.Panics(t, func() { da.InsertAt(100, "X") })
	assert.Panics(t, func() { da.RemoveAt(-1) })
	assert.Panics(t, func() { da.RemoveAt(100) })

	// Serialization Text
	filename := "da_test.txt"
	defer cleanup(filename)
	err := da.SaveToText(filename)
	assert.NoError(t, err)

	da2 := NewDynamicArray()
	err = da2.LoadFromText(filename)
	assert.NoError(t, err)
	assert.Equal(t, da.Size(), da2.Size())
	assert.Equal(t, da.Get(0), da2.Get(0))

	// Serialization Binary
	binFile := "da_test.bin"
	defer cleanup(binFile)
	err = da.SaveToBinary(binFile)
	assert.NoError(t, err)

	da3 := NewDynamicArray()
	err = da3.LoadFromBinary(binFile)
	assert.NoError(t, err)
	assert.Equal(t, da.Size(), da3.Size())

	// Error handling (Bad file)
	assert.Error(t, da.LoadFromText("missing.txt"))
	assert.Error(t, da.LoadFromBinary("missing.bin"))
}

// ==========================================
// 2. LISTS (Singly & Doubly)
// ==========================================
func TestSinglyList_Full(t *testing.T) {
	sl := &SinglyList{}

	// Empty logic
	sl.PopFront()
	sl.PopBack()
	sl.RemoveByValue("X")
	sl.InsertBefore("A", "B") // Should do nothing on empty head

	// Single element logic
	sl.PushBack("A")
	sl.PopBack()
	assert.Equal(t, 0, sl.Size)

	sl.PushFront("B")
	sl.PopFront()
	assert.Equal(t, 0, sl.Size)

	// Complex Logic
	sl.PushBack("A")
	sl.InsertAfter("A", "B") // A -> B
	sl.InsertAfter("A", "C") // A -> C -> B
	assert.True(t, sl.Search("C"))

	sl.InsertBefore("A", "Start") // Start -> A ...
	assert.Equal(t, "Start", sl.Head.Data)
	sl.InsertBefore("B", "PreB") // ... C -> PreB -> B
	assert.True(t, sl.Search("PreB"))

	// Remove logic
	sl.RemoveByValue("Start") // Head
	assert.Equal(t, "A", sl.Head.Data)
	sl.RemoveByValue("B") // Tail
	sl.RemoveByValue("C") // Middle
	sl.RemoveByValue("Missing")

	// Serialization
	sl.PushBack("SaveMe")
	fname := "sl.txt"
	bname := "sl.bin"
	defer cleanup(fname, bname)

	sl.SaveToText(fname)
	sl2 := &SinglyList{}
	sl2.LoadFromText(fname)
	assert.Equal(t, sl.Size, sl2.Size)

	sl.SaveToBinary(bname)
	sl3 := &SinglyList{}
	sl3.LoadFromBinary(bname)
	assert.Equal(t, sl.Size, sl3.Size)
}

func TestDoublyList_Full(t *testing.T) {
	dl := &DoublyList{}

	// Empty
	dl.PopFront()
	dl.PopBack()

	// Push/Pop
	dl.PushBack("A")
	dl.PushFront("B") // B <-> A
	assert.Equal(t, "A", dl.Tail.Data)

	dl.PopFront() // A
	assert.Equal(t, "A", dl.Head.Data)
	dl.PopBack() // Empty
	assert.Equal(t, 0, dl.Size)

	// One element pop back
	dl.PushBack("X")
	dl.PopBack()
	assert.Equal(t, 0, dl.Size)

	// Search
	dl.PushBack("1")
	dl.PushBack("2")
	assert.True(t, dl.Search("2"))
	assert.False(t, dl.Search("3"))

	// Serialization
	fname := "dl.txt"
	bname := "dl.bin"
	defer cleanup(fname, bname)

	dl.SaveToText(fname)
	dl2 := &DoublyList{}
	dl2.LoadFromText(fname)
	assert.Equal(t, "1", dl2.Head.Data)

	dl.SaveToBinary(bname)
	dl3 := &DoublyList{}
	dl3.LoadFromBinary(bname)
	assert.Equal(t, "2", dl3.Tail.Data)
}

// ==========================================
// 3. STACK
// ==========================================
func TestStack_Full(t *testing.T) {
	st := &Stack{}
	assert.Equal(t, "", st.Pop())
	assert.Equal(t, "", st.Peek())

	st.Push("1")
	st.Push("2")
	assert.Equal(t, "2", st.Peek())
	assert.Equal(t, "2", st.Pop())
	assert.Equal(t, "1", st.Pop())

	// Serialization
	st.Push("A")
	st.Push("B")
	fname := "st.txt"
	bname := "st.bin"
	defer cleanup(fname, bname)

	st.SaveToText(fname)
	st2 := &Stack{}
	st2.LoadFromText(fname)
	assert.Equal(t, "B", st2.Peek())

	st.SaveToBinary(bname)
	st3 := &Stack{}
	st3.LoadFromBinary(bname)
	assert.Equal(t, "B", st3.Peek())
}

// ==========================================
// 4. QUEUE (Resize Coverage)
// ==========================================
func TestQueue_Full(t *testing.T) {
	q := NewQueue(2) // Small capacity
	assert.Equal(t, "", q.Pop())
	assert.Equal(t, "", q.Peek())

	q.Push("1")
	q.Push("2")
	q.Push("3") // Triggers resize
	assert.Equal(t, 3, q.Count)
	assert.Equal(t, 4, q.Capacity) // Doubled

	assert.Equal(t, "1", q.Pop())
	assert.Equal(t, "2", q.Peek())

	// Serialization
	fname := "q.txt"
	bname := "q.bin"
	defer cleanup(fname, bname)

	q.SaveToText(fname)
	q2 := NewQueue(10)
	q2.LoadFromText(fname)
	assert.Equal(t, "2", q2.Pop())

	q.SaveToBinary(bname)
	q3 := NewQueue(10)
	q3.LoadFromBinary(bname)
	assert.Equal(t, "2", q3.Peek())
	
	// Error handling
	assert.Error(t, q.LoadFromText("missing.txt"))
}

// ==========================================
// 5. HASH TABLE (Collisions Coverage)
// ==========================================
func TestHashTable_Full(t *testing.T) {
	ht := NewHashTable(2) // Small capacity -> Collisions guaranteed

	ht.Put("A", "1")
	ht.Put("C", "2") // Likely collision
	ht.Put("E", "3") // Collision

	assert.Equal(t, "1", ht.Get("A"))
	assert.Equal(t, "2", ht.Get("C"))
	assert.Equal(t, "3", ht.Get("E"))
	assert.Equal(t, "", ht.Get("Missing"))

	// Update
	ht.Put("A", "11")
	assert.Equal(t, "11", ht.Get("A"))

	// Remove
	ht.Remove("A") // Head of chain
	assert.Equal(t, "", ht.Get("A"))
	assert.Equal(t, "3", ht.Get("E")) // Others remain

	ht.Remove("Missing") // Should be safe

	// Serialization
	fname := "ht.txt"
	bname := "ht.bin"
	defer cleanup(fname, bname)

	ht.SaveToText(fname)
	ht2 := NewHashTable(5)
	ht2.LoadFromText(fname)
	assert.Equal(t, "3", ht2.Get("E"))

	ht.SaveToBinary(bname)
	ht3 := NewHashTable(5)
	ht3.LoadFromBinary(bname)
	assert.Equal(t, "3", ht3.Get("E"))
}

// ==========================================
// 6. TREE
// ==========================================
func TestTree_Full(t *testing.T) {
	tree := &CompleteBinaryTree{}
	assert.False(t, tree.Find(1))
	assert.True(t, tree.IsComplete())

	// Build tree: 1 -> L:2, R:3 -> L:4
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	assert.True(t, tree.Find(4))
	assert.False(t, tree.Find(99))
	assert.True(t, tree.IsComplete())

	// Serialization
	fname := "tree.txt"
	bname := "tree.bin"
	defer cleanup(fname, bname)

	tree.SaveToText(fname)
	t2 := &CompleteBinaryTree{}
	t2.LoadFromText(fname)
	assert.True(t, t2.Find(4))

	tree.SaveToBinary(bname)
	t3 := &CompleteBinaryTree{}
	t3.LoadFromBinary(bname)
	assert.True(t, t3.Find(3))
	
	// Test save empty
	empty := &CompleteBinaryTree{}
	assert.NoError(t, empty.SaveToText("empty.txt"))
	cleanup("empty.txt")
}

// ==========================================
// 7. BENCHMARKS
// ==========================================
func BenchmarkArray(b *testing.B) {
	da := NewDynamicArray()
	for i := 0; i < b.N; i++ {
		da.PushBack("data")
	}
}

func BenchmarkHashTable(b *testing.B) {
	ht := NewHashTable(1000)
	for i := 0; i < b.N; i++ {
		ht.Put(fmt.Sprintf("%d", i), "val")
	}
}