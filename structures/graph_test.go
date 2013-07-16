package structures

import (
	"fmt"
	"testing"
)

func Test_InitializeGraphWorks(t *testing.T) {
	g := Graph{}
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 0, 3)
	g.AddEdge(2, 1, 3)

	if g.Degree[2] != 2 {
		t.Errorf("Wrong degree on vertex")
	}

}
