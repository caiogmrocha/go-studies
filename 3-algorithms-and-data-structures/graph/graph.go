package graph

import "errors"

type Graph[T comparable] struct {
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


func (graph *Graph[T]) AddAdjacence(a *Vertex[T], b *Vertex[T]) (error) {
	bFounded := false
	aFounded := false

	for _, v := range graph.Vertices {
		if v.Data == a.Data {
			v.Adjacences = append(v.Adjacences, b)
			aFounded = true
		}

		if v.Data == b.Data {
			v.Adjacences = append(v.Adjacences, a)
			bFounded = true
		}

		if (aFounded && bFounded) {
			return nil
		}
	}

	if !aFounded {
		return errors.New("a not found")
	} else {
		return errors.New("b not found")
	}
}
