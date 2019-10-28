package data_structures

import (
	"testing"
)

func TestStackInit(t *testing.T) {
	s := NewYPStack()
	if s.Len() != 0 {
		t.Error("initialized stack should be 0 length")
	}

	v := "Hello"
	s.Push(&v)
	s.Push("Hello")
	s.Push(1)
	s.Push(true)

	if s.Len() != 4 {
		t.Error("Stack size should be 4")
	}

	if *(s.Get(0).(*string)) != "Hello" {
		t.Error("Wrong value")
	} else if s.Get(1).(string) != "Hello" {
		t.Error("Wrong value")
	} else if s.Get(2).(int) != 1 {
		t.Error("Wrong value")
	} else if s.Get(3).(bool) != true {
		t.Error("Wrong value")
	}

	if !s.Top().(bool) {
		t.Error("Wrong value")
	}
	if !s.Pop().(bool) {
		t.Error("Wrong value")
	}

	if s.Top().(int) != 1 || s.Pop().(int) != 1 {
		t.Error("Wrong value")
	}

	if s.Top().(string) != "Hello" || s.Pop().(string) != "Hello" {
		t.Error("Wrong value")
	}

	if *(s.Top().(*string)) != "Hello" || *(s.Pop().(*string)) != "Hello" {
		t.Error("Wrong value")
	}

	if s.Len() != 0 {
		t.Error("Stack size should be 0")
	}
}
