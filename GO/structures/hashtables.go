package structures

type HashTable struct {
	data map[string]string
}

func NewHashTable() *HashTable {
	return &HashTable{data: make(map[string]string)}
}

func (h *HashTable) Put(key, val string) {
	h.data[key] = val
}

// ИЗМЕНЕНИЕ: Возвращаем (string, bool)
func (h *HashTable) Get(key string) (string, bool) {
	val, ok := h.data[key]
	return val, ok
}