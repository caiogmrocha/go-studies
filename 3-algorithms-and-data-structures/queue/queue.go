package queue

import (
	"container/list"
	"errors"
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
	q.list.PushBack(value)
}

func (q *Queue[T]) Pop() (T, error) {
	if q.list.Len() == 0 {
		var zero T

		return zero, errors.New("queue is empty")
	}

	element := q.list.Back()
	value := element.Value.(T)
	q.list.Remove(element)

	return value, nil
}

func (q *Queue[T]) ForEachPrint() {
	for element := q.list.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
