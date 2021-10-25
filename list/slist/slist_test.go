package slist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	list := New()
	assert.Equal(t, uint(0), list.Len())
}

func TestList_PushBack(t *testing.T) {
	list := New()
	list.PushBack("test")
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, list.Head(), list.Tail())
	assert.Equal(t, list.Tail().Value, "test")
	assert.Equal(t, list.Front(), "test")
	assert.Equal(t, list.Back(), "test")

	list.PushBack(1)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Tail().Value, 1)
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Front(), "test")
	assert.Equal(t, list.Back(), 1)

	list.PushBack(5)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, list.Tail().Value, 5)
	assert.Equal(t, list.Head().Value, "test")

}

func TestList_PushFront(t *testing.T) {
	list := New()
	list.PushFront("test")
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, list.Head(), list.Tail())
	assert.Equal(t, list.Head().Value, "test")

	list.PushFront(1)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, 1)
	assert.Equal(t, list.Tail().Value, "test")

	list.PushFront(5)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, list.Head().Value, 5)
	assert.Equal(t, list.Tail().Value, "test")
}

func TestList_Insert(t *testing.T) {
	list := New()
	list.Insert("test", 0)
	assert.Equal(t, uint(1), list.Len())
	assert.Equal(t, list.Head(), list.Tail())
	assert.Equal(t, list.Head().Value, "test")

	list.Insert(999, 1)
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Tail().Value, 999)

	list.Insert(888, 10)
	assert.Equal(t, uint(3), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Next().Value, 999)
	assert.Equal(t, list.Tail().Value, 888)

	// "test" --> "test1" --> 999 --> 888
	list.Insert("test1", 1)
	assert.Equal(t, uint(4), list.Len())
	assert.Equal(t, list.Head().Value, "test")
	assert.Equal(t, list.Head().Next().Value, "test1")
	assert.Equal(t, list.Head().Next().Next().Value, 999)
	assert.Equal(t, list.Tail().Value, 888)
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
	assert.Equal(t, list.Tail().Value, "2")

	// 2 --> 3 --> 4, remove 3
	list.PushBack("3")
	list.PushBack("4")
	list.Remove(list.Find("3"))
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "2")
	assert.Equal(t, list.Head().Next().Value, "4")
	assert.Equal(t, list.Tail().Value, "4")

	// 2 --> 3 --> 4, remove 4
	list.Insert("3", 1)
	list.Remove(list.Find("4"))
	assert.Equal(t, uint(2), list.Len())
	assert.Equal(t, list.Head().Value, "2")
	assert.Equal(t, list.Head().Next().Value, "3")
	assert.Equal(t, list.Tail().Value, "3")

}

func TestList_Traverse(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	ret := make([]int, 0)
	list.Traverse(func(value interface{}) bool {
		ret = append(ret, value.(int))
		return true
	})
	assert.Equal(t, "[1 2 3 4]", fmt.Sprintf("%v", ret))
}
