package structures

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- DYNAMIC ARRAY ---
func TestDynamicArray_Full(t *testing.T) {
	arr := NewDynamicArray()

	arr.PushBack("A")
	val, err := arr.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "A", val)

	// Ошибка выхода за границы
	_, err = arr.Get(99)
	assert.Error(t, err)

	// Сериализация JSON (успех)
	err = arr.SaveToText("arr.json")
	assert.NoError(t, err)

	loaded := NewDynamicArray()
	err = loaded.LoadFromText("arr.json")
	assert.NoError(t, err)
	val, _ = loaded.Get(0)
	assert.Equal(t, "A", val)

	// Сериализация Gob (успех)
	err = arr.SaveToBinary("arr.gob")
	assert.NoError(t, err)

	loadedBin := NewDynamicArray()
	err = loadedBin.LoadFromBinary("arr.gob")
	assert.NoError(t, err)

	// === ТЕСТЫ ОШИБОК ДЛЯ ПОВЫШЕНИЯ ПОКРЫТИЯ ===
	// Попытка загрузить несуществующий файл
	err = loaded.LoadFromText("bad_file.json")
	assert.Error(t, err)

	err = loaded.LoadFromBinary("bad_file.gob")
	assert.Error(t, err)

	// Очистка
	os.Remove("arr.json")
	os.Remove("arr.gob")
}

// --- LINKED LIST ---
func TestList(t *testing.T) {
	list := NewLinkedList()
	assert.Equal(t, "", list.GetHead())

	list.PushBack("head")
	list.PushBack("tail")
	assert.Equal(t, "head", list.GetHead())
}

// --- STACK ---
func TestStack(t *testing.T) {
	st := NewStack()
	assert.Equal(t, "", st.Pop()) // пустой

	st.Push("1")
	st.Push("2")
	assert.Equal(t, "2", st.Pop())
	assert.Equal(t, "1", st.Pop())
}

// --- QUEUE ---
func TestQueue(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, "", q.Pop()) // пустая

	q.Push("1")
	q.Push("2")
	assert.Equal(t, "1", q.Pop())
	assert.Equal(t, "2", q.Pop())
}

// --- HASH TABLE ---
func TestHashTable(t *testing.T) {
	ht := NewHashTable()
	ht.Put("k", "v")
	val, ok := ht.Get("k")
	assert.True(t, ok)
	assert.Equal(t, "v", val)
}

// --- TREE ---
func TestTree(t *testing.T) {
	tree := NewTree()
	
	// Вставляем достаточно элементов, чтобы заполнить несколько уровней
	// и зайти во все ветки else (append в очередь)
	tree.Insert(1) // Корень
	tree.Insert(2) // Левый
	tree.Insert(3) // Правый (сработает else для левого)
	tree.Insert(4) // Левый от 2 (сработает else для правого у корня и спуск вниз)
	
	assert.True(t, tree.Find(1))
	assert.True(t, tree.Find(3))
	assert.True(t, tree.Find(4))
	assert.False(t, tree.Find(99))

	// Проверка пустого дерева
	emptyTree := NewTree()
	assert.False(t, emptyTree.Find(1))
}

// --- BENCHMARK ---
func BenchmarkArrayPush(b *testing.B) {
	arr := NewDynamicArray()
	for i := 0; i < b.N; i++ {
		arr.PushBack("data")
	}
}
