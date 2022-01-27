package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q Queue) New() Queue {
	q.items = []int{}
	return q
}

func (q *Queue) Enqueue(item int) {
	(*q).items = append((*q).items, item)
	return
}

func (q *Queue) Dequeue() int {
	firstItem := (*q).items[0]
	(*q).items = (*q).items[1:len(q.items)]
	return firstItem
}

func (q Queue) Read() int {
	//Reads the first item from the queue
	if len(q.items) >= 1 {
		return q.items[0]
	}
	return -1

}

func main() {
	cinemaQueue := Queue{}.New()

	cinemaQueue.Enqueue(1)
	cinemaQueue.Enqueue(2)
	cinemaQueue.Enqueue(3)
	cinemaQueue.Enqueue(4)
	cinemaQueue.Enqueue(5)
	cinemaQueue.Enqueue(6)

	fmt.Printf("First item in the Queue is %d\n", cinemaQueue.Read())
	fmt.Printf("Just Finished Serving Item: %d\n", cinemaQueue.Dequeue())
	fmt.Printf("Next Item in the Queue is %d\n", cinemaQueue.Read())

}
