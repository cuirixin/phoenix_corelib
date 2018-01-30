package utils

import "testing"

type Teacher struct {
	Name string
	Age int64
	Avator string
	Students []string
}

func TestReflectStructInfo(t *testing.T) {
	teacher := Teacher{"crx", 30, "", []string{"lll","ddd"}}
	ReflectStructInfo(teacher)
}
