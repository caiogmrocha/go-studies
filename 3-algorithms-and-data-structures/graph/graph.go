package graph

import "errors"

type Graph[T comparable] struct {
	Vertices map[string](*Vertex[T])
}

type Vertex[T any] struct {
	Data T
	Adjacences []*Vertex[T]
}

func (graph *Graph[T]) AddVertex(id string, vertex *Vertex[T]) (error) {
	if (graph == nil) {
		return errors.New("graph is nil")
	}

	if graph.Vertices == nil {
		graph.Vertices = map[string](*Vertex[T]){}
	}

	graph.Vertices[id] = vertex

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

func (graph *Graph[T]) DFS(data T) *Vertex[T] {


	return nil
}
