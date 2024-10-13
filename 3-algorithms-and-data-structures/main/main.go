package main

import (
	"ads/graph"
	"ads/queue"
	"fmt"
)

func main() {
	testGraphPackage()
}

func testGraphPackage() {
	g := &graph.Graph[string]{}

	g.AddVertex(&graph.Vertex[string]{Id: "A", Data: "A"})
	g.AddVertex(&graph.Vertex[string]{Id: "B", Data: "B"})
	g.AddVertex(&graph.Vertex[string]{Id: "C", Data: "C"})
	g.AddVertex(&graph.Vertex[string]{Id: "D", Data: "D"})
	g.AddVertex(&graph.Vertex[string]{Id: "E", Data: "E"})

	g.AddAdjacence(g.Vertices["A"], g.Vertices["B"])
	g.AddAdjacence(g.Vertices["B"], g.Vertices["C"])
	g.AddAdjacence(g.Vertices["C"], g.Vertices["D"])
	g.AddAdjacence(g.Vertices["D"], g.Vertices["E"])
	g.AddAdjacence(g.Vertices["E"], g.Vertices["A"])

	vertexToBeFounded := "F"

	foundedVertex, error := graph.DFS(vertexToBeFounded, g.Vertices["A"], map[string]bool{}, queue.New[*graph.Vertex[string]]())

	if error != nil {
		fmt.Println("error in DFS:", error)

		return
	}

	if foundedVertex == nil {
		fmt.Printf("vertex not found: %s\n", vertexToBeFounded)

		return
	}

	fmt.Printf("vertex found: %s\n", foundedVertex.Id)
}

// func printNodeValue(n *bst.Node) (error) {
// 	if (n == nil) {
// 		return errors.New("n is nil")
// 	}

// 	fmt.Println(n.Value)

// 	return nil
// }

// func testBSTPackage() {
// 	tree := &bst.BST{}

// 	tree.Insert(10)
// 	tree.Insert(5)
// 	tree.Insert(15)
// 	tree.Insert(1)
// 	tree.Insert(7)

// 	tree.InOrderTraversal(printNodeValue)

// 	tree.Remove(5)

// 	tree.InOrderTraversal(printNodeValue)
// }
