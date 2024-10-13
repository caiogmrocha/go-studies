package main

import "ads/queue"

func main() {
	q := &queue.Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	q.ForEachPrint()
}

// func testGraphPackage() {
// 	g := &graph.Graph[string]{}

// 	g.AddVertex("A", &graph.Vertex[string]{Data: "A"})
// 	g.AddVertex("B", &graph.Vertex[string]{Data: "B"})
// 	g.AddVertex("C", &graph.Vertex[string]{Data: "C"})
// 	g.AddVertex("D", &graph.Vertex[string]{Data: "D"})
// 	g.AddVertex("E", &graph.Vertex[string]{Data: "E"})

// 	g.AddAdjacence(g.Vertices["A"], g.Vertices["B"])
// 	g.AddAdjacence(g.Vertices["B"], g.Vertices["C"])
// 	g.AddAdjacence(g.Vertices["C"], g.Vertices["D"])
// 	g.AddAdjacence(g.Vertices["D"], g.Vertices["E"])
// 	g.AddAdjacence(g.Vertices["E"], g.Vertices["A"])

// 	fmt.Println("Vertices:")

// 	for k := range g.Vertices {
// 		fmt.Printf("%s\n", k)
// 	}
// }

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
