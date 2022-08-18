package main

import (
	"container/heap"
	"fmt"
)

// Estructura de cada Nodo-Elemento de la Cola-Queue
type Task struct {
	priority int
	value    string
}

type TaskPQ []Task

// Longitu Cola
func (self TaskPQ) Len() int {
	return len(self)
}

// Elegir segun Prioridad
func (self TaskPQ) Less(i, j int) bool {
	return self[i].priority < self[j].priority
}

// Intercambiar. Se usa cuando se saca un elemento
func (self TaskPQ) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// Encolar: Enqueue
func (self *TaskPQ) Push(x interface{}) {
	*self = append(*self, x.(Task))
}

// Descolar: Dequeue
func (self *TaskPQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}

func main() {
	// Valores Iniciales
	pq := &TaskPQ{
		{3, "W"},
		{2, "H"},
		{1, "Y"}}

	// heapify: Inicar Cola, Amontonar
	heap.Init(pq)

	// enqueue: Meter a la Cola
	heap.Push(pq, Task{2, "A"})
	heap.Push(pq, Task{7, "B"})
	heap.Push(pq, Task{4, "SH"})

	// enqueue: Sacar de la Cola
	fmt.Println(heap.Pop(pq))

	// dequeue all: Sacar a todos de la Cola
	for pq.Len() != 0 {
		fmt.Println(heap.Pop(pq))
	}
}
