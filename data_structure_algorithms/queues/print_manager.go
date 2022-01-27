package main

import (
	"fmt"
)

func main() {

	documentQueue := PrintQueue{}.New()

	documentQueue.Enqueue("First Document")
	documentQueue.Enqueue("Second Document")
	documentQueue.Enqueue("Third Document")
	documentQueue.Enqueue("Fourth Document")
	documentQueue.Enqueue("Fifth Document")

	documentQueue.Run()
}

type PrintQueue struct {
	documents []string
}

func (q PrintQueue) New() PrintQueue {
	q.documents = []string{}
	return q
}

func (q *PrintQueue) Enqueue(item string) {
	(*q).documents = append((*q).documents, item)
	return
}

func (q *PrintQueue) Dequeue() string {
	if len((*q).documents) >= 1 {
		firstItem := (*q).documents[0]
		(*q).documents = (*q).documents[1:len(q.documents)]
		return firstItem
	}
	return "No More Documents to Print"
}

func (q *PrintQueue) Run() {
	for len((*q).documents) >= 1 {
		fmt.Printf("Currently Printing: %s\n", (*q).Dequeue())
	}

	return
}

func (q PrintQueue) Read() string {
	//Reads the first item from the queue
	if len(q.documents) >= 1 {
		return q.documents[0]
	}
	return "No more documents to Print!!!"
}
