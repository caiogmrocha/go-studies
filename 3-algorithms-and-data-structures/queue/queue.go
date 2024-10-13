package queue

import (
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	list *list.List
};

func New[T any]() (*Queue[T]) {
	q := &Queue[T]{}

	if q.list == nil {
		q.list = list.New()
	}

	return q
}

func (q *Queue[T]) Enqueue(value T) {
	if q.list == nil {
		q.list = list.New()
	}

	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) ForEachPrint() {
	for element := q.list.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
