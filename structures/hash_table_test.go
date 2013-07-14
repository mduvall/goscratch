package structures

import (
	"testing"
)

func Test_HashKeysReturnIndices(t *testing.T) {
	h := HashTable{}
	h.Insert("foo", "bar")
	if h.Get("foo") != "bar" {
		t.Errorf("This is not a hash table.")
	}
}
