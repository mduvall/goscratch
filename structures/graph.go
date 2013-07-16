package structures

const MAXV = 10000

type Graph struct {
	Vertices [MAXV][]Edge
	Degree   [MAXV]int
}

type Edge struct {
	endVertex int
	weight    int
}

func (g *Graph) AddEdge(startVertex int, endVertex int, weight int) {
	edge := Edge{endVertex: endVertex, weight: weight}
	if g.Vertices[startVertex] == nil {
		g.Vertices[startVertex] = []Edge{edge}
	} else {
		g.Vertices[startVertex] = append(g.Vertices[startVertex], edge)
	}

	g.Degree[startVertex]++
}
