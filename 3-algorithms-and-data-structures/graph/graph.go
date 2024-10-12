package graph

import "errors"

type Graph[T any] struct {
	Vertices []*Vertex[T]
}

type Vertex[T any] struct {
	Data T
	Adjacences []*Vertex[T]
}

func (graph *Graph[T]) AddVertex(vertex *Vertex[T]) (error) {
	if (graph == nil) {
		return errors.New("graph is nil")
	}

	graph.Vertices = append(graph.Vertices, vertex)

	return nil
}
