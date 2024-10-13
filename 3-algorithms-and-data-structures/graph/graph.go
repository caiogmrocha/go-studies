package graph

import (
	"errors"

	"ads/queue"
)

type Graph[T comparable] struct {
	Vertices map[string](*Vertex[T])
}

type Vertex[T comparable] struct {
	Id string
	Data T
	Adjacences []*Vertex[T]
}

func New[T comparable]() *Graph[T] {
	g := &Graph[T]{}

	return g
}

func (graph *Graph[T]) AddVertex(vertex *Vertex[T]) (error) {
	if (graph == nil) {
		return errors.New("graph is nil")
	}

	if graph.Vertices == nil {
		graph.Vertices = map[string](*Vertex[T]){}
	}

	graph.Vertices[vertex.Id] = vertex

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

func DFS[T comparable](
	data T,
	origin *Vertex[T],
	visited *map[string]bool,
	pathQueue *queue.Queue[*Vertex[T]],
) (*Vertex[T], error) {
	if origin == nil {
		return nil, errors.New("null vertex")
	}

	if (origin.Data == data) {
		return origin, nil
	}

	(*visited)[origin.Id] = true

	for _, adjacence := range origin.Adjacences {
		if (!(*visited)[adjacence.Id]) {
			pathQueue.Enqueue(adjacence)
		}
	}

	var foundedVertex *Vertex[T]

	for pathQueue.Len > 0 && foundedVertex == nil {
		var zero *Vertex[T]

		nonVisitedVertex, error := pathQueue.Pop()

		if (error != nil) {
			return zero, error
		}

 		foundedVertex, error = DFS(data, nonVisitedVertex, visited, pathQueue)

		if (error != nil) {
			return zero, error
		}
	}

	return foundedVertex, nil
}
