package structures

type HashTable struct {
	Elements [1000]string
}

func (h *HashTable) Insert(key string, val string) {
	h.Elements[GetHashIndex(key, 1000)] = val
}

func (h *HashTable) Get(key string) string {
	return h.Elements[GetHashIndex(key, 1000)]
}

func GetHashIndex(key string, length int) int {
	keyAsBytes := []byte(key)
	sum := 0

	reset, i, mult := 0, 0, 1

	for _, b := range keyAsBytes {
		if reset == 2 {
			reset = 0
			sum += i * mult
			mult *= 256
			i = 0
		}

		i = (i << 8) | int(b)
		reset++
	}

	return sum % length
}
