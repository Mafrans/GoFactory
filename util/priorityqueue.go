package util

import "golang.org/x/exp/slices"

type PriorityQueueItem[T any] struct {
	Priority float64
	Value    T
}

type PriorityQueue[T any] []PriorityQueueItem[T]

func (queue *PriorityQueue[T]) Push(item T, priority float64) {
	*queue = append(*queue, PriorityQueueItem[T]{
		Priority: priority,
		Value:    item,
	})
}

func (queue *PriorityQueue[T]) Pop() PriorityQueueItem[T] {
	highest, index := (*queue).Highest()
	slices.Delete(*queue, index, index+1)

	return highest
}

func (queue *PriorityQueue[T]) Highest() (PriorityQueueItem[T], int) {
	var index int
	result := (*queue)[0]

	for i, item := range *queue {
		if result.Priority < item.Priority {
			result, index = item, i
		}
	}

	return result, index
}

func (queue *PriorityQueue[T]) Lowest() (PriorityQueueItem[T], int) {
	var index int
	result := (*queue)[0]

	for i, item := range *queue {
		if result.Priority > item.Priority {
			result, index = item, i
		}
	}

	return result, index
}
