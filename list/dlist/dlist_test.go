package dlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := New()
	assert.Equal(t, uint(0), list.Len())
}

func TestPushBack(t *testing.T) {
	list := New()
	list.PushBack(1)
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 1, list.Head().Next().Value)
	assert.Equal(t, 1, list.Head().Prev().Value)
	assert.Equal(t, 1, list.Back())
	assert.Equal(t, 1, list.Front())


	list.PushBack(2)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 2, list.Head().Next().Value)
	assert.Equal(t, 2, list.Head().Prev().Value)
	assert.Equal(t, 1, list.Head().Next().Prev().Value)
	assert.Equal(t, 1, list.Head().Next().Next().Value)
	assert.Equal(t, 2, list.Back())
	assert.Equal(t, 1, list.Front())

	list.PushBack(3)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 2, list.Head().Next().Value)
	assert.Equal(t, 3, list.Head().Prev().Value)
	assert.Equal(t, 2, list.Head().Next().Next().Prev().Value)
	assert.Equal(t, 1, list.Head().Next().Next().Next().Value)
	assert.Equal(t, 3, list.Head().Next().Next().Value)

	assert.Equal(t, 3, list.Back())
	assert.Equal(t, 1, list.Front())

}

func TestPushFront(t *testing.T) {
	list := New()
	list.PushFront(1)
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 1, list.Head().Next().Value)
	assert.Equal(t, 1, list.Head().Prev().Value)


	list.PushFront(2)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, 2, list.Head().Value)
	assert.Equal(t, 1, list.Head().Next().Value)
	assert.Equal(t, 1, list.Head().Prev().Value)
	assert.Equal(t, 2, list.Head().Next().Prev().Value)
	assert.Equal(t, 2, list.Head().Next().Next().Value)

	list.PushFront(3)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, 3, list.Head().Value)
	assert.Equal(t, 2, list.Head().Next().Value)
	assert.Equal(t, 1, list.Head().Prev().Value)
	assert.Equal(t, 2, list.Head().Next().Next().Prev().Value)
	assert.Equal(t, 1, list.Head().Next().Next().Value)
	assert.Equal(t, 3, list.Head().Next().Next().Next().Value)

}

func TestInsert(t *testing.T) {
	list := New()
	list.Insert("test", 0)
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Prev().Value, "test")
	assert.Equal(t, list.Head().Next().Value, "test")

	list.Insert(999, 1)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Next().Value, 999)
	assert.Equal(t, list.Head().Next().Next().Value, "test")
	assert.Equal(t, list.Head().Next().Prev().Value, "test")

	list.Insert(888, 10)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Next().Next().Value, 888)
	assert.Equal(t, list.Head().Next().Value, 999)
	assert.Equal(t, list.Head().Next().Prev().Value, "test")

	// "test" --> "test1" --> 999 --> 888
	list.Insert("test1", 1)
	assert.Equal(t, uint(4), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Next().Value, "test1")
	assert.Equal(t, list.Head().Next().Next().Value, 999)
	assert.Equal(t, list.Head().Next().Next().Next().Value, 888)

	assert.Equal(t, list.Head().Prev().Value, 888)
	assert.Equal(t, list.Head().Prev().Prev().Value, 999)
	assert.Equal(t, list.Head().Prev().Prev().Prev().Value, "test1")
	assert.Equal(t, list.Head().Prev().Prev().Prev().Prev().Value, "test")
}

func TestList_Find(t *testing.T) {
	list := New()
	list.PushBack("1")
	list.PushBack("2")
	list.PushBack("3")
	list.PushBack("4")

	assert.Equal(t, 0, list.Find("1"))
	assert.Equal(t, 3, list.Find("4"))
	assert.Equal(t, -1, list.Find("5"))
}

func TestList_Remove(t *testing.T) {
	list := New()
	list.PushBack("1")
	list.Remove(list.Find("1"))
	assert.Equal(t, uint(0), list.Len())

	// 1 --> 2 , remove 1
	list.PushBack("1")
	list.PushBack("2")
	list.Remove(list.Find("1"))
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, list.Head().Value, "2")
	assert.Equal(t, list.Head().Next().Value, "2")
	assert.Equal(t, list.Head().Prev().Value, "2")

	// 2 --> 3 --> 4, remove 3
	list.PushBack("3")
	list.PushBack("4")
	list.Remove(list.Find("3"))
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "2")
	assert.Equal(t, list.Head().Prev().Value, "4")

	tail := list.Head().Next()
	assert.Equal(t, tail.Value, "4")
	assert.Equal(t, tail.Prev().Value, "2")
	assert.Equal(t, tail.Next().Value, "2")

	// 2 --> 3 --> 4, remove 4
	list.Insert("3", 1)
	list.Remove(list.Find("4"))
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "2")
	assert.Equal(t, list.Head().Next().Value, "3")

	tail2 := list.Head().Next()
	assert.Equal(t, tail2.Value, "3")
	assert.Equal(t, tail2.Prev().Value, "2")
	assert.Equal(t, tail2.Next().Value, "2")
}

func TestList_PopFrontAndBack(t *testing.T) {
	list := New()
	list.PushBack("1")
	assert.Equal(t, "1", list.PopBack())
	assert.Equal(t, uint(0), list.Len())

	list.PushBack("1")
	assert.Equal(t, "1", list.PopFront())
	assert.Equal(t, uint(0), list.Len())

	// 1 --> 2 , remove 1
	list.PushBack("1")
	list.PushBack("2")
	assert.Equal(t, "1", list.PopFront())

	list.PushBack("3")
	assert.Equal(t, "3", list.PopBack())
}
