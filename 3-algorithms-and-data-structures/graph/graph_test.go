package graph_test

import (
	"ads/graph"
	"ads/queue"
	"testing"
)

func TestGraphAddVertex(t *testing.T) {
	g := graph.New[int]()

	g.AddVertex(&graph.Vertex[int]{Id: "A", Data: 1})
	g.AddVertex(&graph.Vertex[int]{Id: "B", Data: 2})
	g.AddVertex(&graph.Vertex[int]{Id: "C", Data: 3})
	g.AddVertex(&graph.Vertex[int]{Id: "D", Data: 4})
	g.AddVertex(&graph.Vertex[int]{Id: "E", Data: 5})

	if (!(
		g.Vertices != nil      ||
		g.Vertices["A"] != nil ||
		g.Vertices["B"] != nil ||
		g.Vertices["C"] != nil ||
		g.Vertices["D"] != nil ||
		g.Vertices["E"] != nil)) {
		t.Fatalf("Graph.AddVertex() should be correctly structure")
	}
}

func TestGraphDFSAddVertex(t *testing.T) {
	g := graph.New[string]()

	g.AddVertex(&graph.Vertex[string]{Id: "A", Data: "A"})
	g.AddVertex(&graph.Vertex[string]{Id: "B", Data: "B"})
	g.AddVertex(&graph.Vertex[string]{Id: "C", Data: "C"})
	g.AddVertex(&graph.Vertex[string]{Id: "D", Data: "D"})
	g.AddVertex(&graph.Vertex[string]{Id: "E", Data: "E"})

	vertexToBeFounded := "D"

	foundedVertex, error := graph.DFS(vertexToBeFounded, g.Vertices[vertexToBeFounded], &map[string]bool{}, queue.New[*graph.Vertex[string]]())

	if error != nil {
		t.Fatal(error)
	}

	if foundedVertex == nil {
		t.Fatalf("Graph.DFS() not founded any vertex.\nExpected: A\nReceived: nil")
	}

	if foundedVertex.Id != vertexToBeFounded {
		t.Fatalf("Graph.DFS() founded wrong vertex.\nExpected: A\nReceived: %s", foundedVertex.Id)
	}
}
