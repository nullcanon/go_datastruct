package utils

type Visitor func(v interface{}) bool
type Compare func(left interface{}, right interface{}) int

