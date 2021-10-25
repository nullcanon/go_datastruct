package utils

type Container interface {
	PushBack(value interface{})
	PopBack() interface{}
	PushFront(value interface{})
	PopFront() interface{}
	Len() uint
	Empty() bool
	Back() interface{}
	Front() interface{}
}