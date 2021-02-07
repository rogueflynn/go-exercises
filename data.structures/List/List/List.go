package List

import "fmt"

type Queue []string

func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() {
	// Dereference array and slice from index 1 to the end of the slice and assign to type queue
	if len(*q) > 0 {
		*q = (*q)[1:]
	} else {
		fmt.Println("empty")
	}
}

func (q *Queue) Get() {
	fmt.Println(*q)
}