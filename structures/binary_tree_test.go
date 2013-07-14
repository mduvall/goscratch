package structures

import (
	"testing"
)

func Test_CreateTreeWithIntSlice(t *testing.T) {
	tree := CreateTree([]int{5, 4, 2, 6, 8, 3, 54, 53, 2})

	if tree.Search(3) == nil {
		t.Errorf("You messed up here man")
	}

	if tree.Search(54) == nil {
		t.Errorf("You messed up here man")
	}

	if tree.Search(100) != nil {
		t.Errorf("Thats not in the tree")
	}

	if tree.Min() != 2 {
		t.Errorf("Wrong min")
	}

	if tree.Max() != 54 {
		t.Errorf("Wrong max")
	}

}
