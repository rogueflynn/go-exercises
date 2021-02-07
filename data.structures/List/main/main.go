package main

import (
	"example.com/List"
	)

func main() {
	var queue List.Queue
	queue.Enqueue("test")
	queue.Enqueue("test2")
	queue.Enqueue("test3")
	queue.Get()
	queue.Dequeue();
	queue.Get()
	queue.Dequeue();
	queue.Get()
}
