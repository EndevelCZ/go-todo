package todo

import (
	"reflect"
	"testing"
)

func BenchmarkTodoToPbTodo(b *testing.B) {
	t := &Todo{
		ID:   1,
		Text: "Hello",
		Done: false,
	}
	for n := 0; n < b.N; n++ {
		TodoToPbTodo(t)
	}
}

func BenchmarkTodoToPbTodo2(b *testing.B) {
	t := &Todo{
		ID:   1,
		Text: "Hello",
		Done: false,
	}
	for n := 0; n < b.N; n++ {
		TodoToPbTodo2(t)
	}
}

func TestSerializer(t *testing.T) {
	t1 := &Todo{ID: 1, Text: "Hello", Done: false}
	t2 := &Todo{ID: 1, Text: "Hello", Done: false}
	m1 := TodoToPbTodo(t1)
	m2 := TodoToPbTodo2(t2)
	if !reflect.DeepEqual(m1, m2) {
		t.Errorf("serializers are not equal %v %v", t1, t2)
	}
}
