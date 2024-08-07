package linkedlist_test

import (
	"testing"

	ll "github.com/pactus-project/pactus/util/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestDoublyLink_InsertAtHead(t *testing.T) {
	link := ll.New[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.InsertAtHead(4)

	assert.Equal(t, []int{4, 3, 2, 1}, link.Values())
	assert.Equal(t, 4, link.Length())
	assert.Equal(t, 4, link.Head.Data)
	assert.Equal(t, 1, link.Tail.Data)
}

func TestSinglyLink_InsertAtTail(t *testing.T) {
	link := ll.New[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	assert.Equal(t, []int{1, 2, 3, 4}, link.Values())
	assert.Equal(t, 4, link.Length())
	assert.Equal(t, 1, link.Head.Data)
	assert.Equal(t, 4, link.Tail.Data)
}

func TestDeleteAtHead(t *testing.T) {
	link := ll.New[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.DeleteAtHead()
	assert.Equal(t, []int{2, 3}, link.Values())
	assert.Equal(t, 2, link.Length())

	link.DeleteAtHead()
	assert.Equal(t, []int{3}, link.Values())
	assert.Equal(t, 1, link.Length())

	link.DeleteAtHead()
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())

	link.DeleteAtHead()
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())
}

func TestDeleteAtTail(t *testing.T) {
	link := ll.New[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.DeleteAtTail()
	assert.Equal(t, []int{1, 2}, link.Values())
	assert.Equal(t, 2, link.Length())

	link.DeleteAtTail()
	assert.Equal(t, []int{1}, link.Values())
	assert.Equal(t, 1, link.Length())

	link.DeleteAtTail()
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())

	link.DeleteAtTail()
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())
}

func TestDelete(t *testing.T) {
	link := ll.New[int]()
	n1 := link.InsertAtTail(1)
	n2 := link.InsertAtTail(2)
	n3 := link.InsertAtTail(3)
	n4 := link.InsertAtTail(4)

	link.Delete(n1)
	assert.Equal(t, []int{2, 3, 4}, link.Values())
	assert.Equal(t, 3, link.Length())

	link.Delete(n4)
	assert.Equal(t, []int{2, 3}, link.Values())
	assert.Equal(t, 2, link.Length())

	link.Delete(n2)
	assert.Equal(t, []int{3}, link.Values())
	assert.Equal(t, 1, link.Length())

	link.Delete(n3)
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())
}

func TestClear(t *testing.T) {
	link := ll.New[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)

	link.Clear()
	assert.Equal(t, []int{}, link.Values())
	assert.Equal(t, 0, link.Length())
}

func TestInsertAfter(t *testing.T) {
	link := ll.New[int]()
	e1 := link.InsertAtHead(1)
	e2 := link.InsertAfter(2, e1)
	link.InsertAfter(3, e2)
	link.InsertAfter(4, link.Head)
	link.InsertAfter(5, link.Tail)

	assert.Equal(t, []int{1, 4, 2, 3, 5}, link.Values())
}

func TestInsertBefore(t *testing.T) {
	link := ll.New[int]()
	e1 := link.InsertAtHead(1)
	e2 := link.InsertBefore(2, e1)
	link.InsertBefore(3, e2)
	link.InsertBefore(4, link.Head)
	link.InsertBefore(5, link.Tail)

	assert.Equal(t, []int{4, 3, 2, 5, 1}, link.Values())
}
