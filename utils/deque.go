package utils

type Stack[T any] interface {
	Empty() bool
	Peek() T
	Pop() (T, bool)
	Push(item T)
}

type Queue[T any] interface {
	Empty() bool
	Peek() T
	Dequeue() (T, bool)
	Enqueue(item T)
}

type node[T any] struct {
	data T
	prev *node[T]
	next *node[T]
}

type Deque[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewDeque[T any]() *Deque[T] {
	var deque Deque[T]
	return &deque
}

func NewDequeFrom[T any](slice []T) *Deque[T] {
	var deque Deque[T]
	for _, value := range slice {
		deque.Enqueue(value)
	}
	return &deque
}

func (d *Deque[T]) Push(item T) {
	newNode := &node[T]{
		data: item,
		prev: nil,
		next: nil,
	}

	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.length++
}

func (d *Deque[T]) Pop() (T, bool) {
	if d.length == 0 {
		var nope T
		return nope, false
	} else {
		temp := d.head
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		} else {
			d.tail = nil
		}
		d.length--
		return temp.data, true
	}
}

func (d *Deque[T]) Peek() T {
	return d.head.data
}

func (d *Deque[T]) Enqueue(item T) {
	newNode := &node[T]{
		data: item,
		prev: nil,
		next: nil,
	}

	if d.length == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.length++
}

func (d *Deque[T]) Dequeue() (T, bool) {
	return d.Pop()
}

func (d *Deque[T]) Empty() bool {
	return d.length == 0
}

func (d *Deque[T]) Clear() {
	var deque Deque[T]
	*d = deque
}

func (d *Deque[T]) Len() int {
	return d.length
}
