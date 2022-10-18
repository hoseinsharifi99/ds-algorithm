package main

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	root   *node
	Length int
}

func (l *LinkedList) insert(val interface{}, position int) bool {
	if position > l.Length || position < 0 {
		return false
	} else if position == 0 {
		newNode := &node{
			value: val,
			next:  l.root,
		}
		l.root = newNode
		l.Length++
	} else {
		currentNode := l.root
		currentIndex := 0
		for ; currentIndex < position-1; currentIndex++ {
			currentNode = currentNode.next
		}
		newNode := &node{
			value: val,
			next:  currentNode.next,
		}
		currentNode.next = newNode
		l.Length++
	}
	return true
}

func (l *LinkedList) Push(val interface{}) bool {
	return l.insert(val, l.Length)
}

func (l *LinkedList) unshift(val interface{}) bool {
	return l.insert(val, 0)
}

func (l *LinkedList) remove(position int) (val interface{}) {
	if l.Length == 0 || position >= l.Length {
		return nil
	}
	if position == 0 {
		val = l.root.value
		l.root = l.root.next
		l.Length--
	} else {
		currentNode := l.root
		currentIndex := 0

		for ; currentIndex < position-1; currentIndex++ {
			currentNode = currentNode.next
		}
		val = currentNode.next.value
		currentNode = currentNode.next.next
		l.Length--
	}
	return
}

func (l *LinkedList) pop() interface{} {
	return l.remove(l.Length - 1)
}

func (l *LinkedList) Shift() interface{} {
	return l.remove(0)
}

func (l *LinkedList) Traverse() (response []interface{}) {
	currentNode := l.root
	for currentNode != nil {
		response = append(response, currentNode.value)
		currentNode = currentNode.next
	}
	return
}

func (l *LinkedList) reverse() {
	curr := l.root
	var prev *node
	var next *node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.root = prev

}
