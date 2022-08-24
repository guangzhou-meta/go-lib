package collections

import (
	"testing"
)

func TestPriorityQueue_Empty(t *testing.T) {
	cases := []struct {
		name   string
		q      []interface{}
		expect bool
	}{
		{"空数组", []interface{}{}, true},
		{"空数组", []interface{}{}, false},
		{"非空数组", []interface{}{10, 32, 2, 19, 28, 69}, false},
		{"非空数组", []interface{}{10, 32, 2, 19, 28, 69}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if pq := New(c.q); pq.IsEmpty() != c.expect {
				t.Fatalf("%s New(%T) expected %t, but got %t",
					c.name, c.q, c.expect, pq.IsEmpty())
			}
		})
	}
}

func TestPriorityQueue_Push_Peek(t *testing.T) {
	cases := []struct {
		name   string
		q      []interface{}
		expect []interface{}
	}{
		{"非空数组", []interface{}{10, 32, 2, 19, 28, 69}, []interface{}{10, 10, 2, 2, 2, 2}},
		{"非空数组", []interface{}{69, 32, 22, 19, 20, 11}, []interface{}{69, 32, 22, 19, 19, 11}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pq := New([]interface{}{})
			for i := 0; i < len(c.q); i++ {
				x := c.q[i]
				pq.Push(x)
				if pq.Peek() != c.expect[i] {
					t.Fatalf("%s peek expected %d, but got %d",
						c.name, c.expect, pq.Peek())
				}
			}
		})
	}
}

func TestPriorityQueue_Push_Pop(t *testing.T) {
	cases := []struct {
		name   string
		q      []interface{}
		expect []interface{}
	}{
		{"非空数组", []interface{}{10, 32, 2, 19, 28, 69}, []interface{}{10, 32, 2, 19, 28, 69}},
		{"非空数组", []interface{}{69, 32, 22, 19, 20, 11}, []interface{}{69, 32, 22, 19, 20, 11}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pq := New([]interface{}{})
			for i := 0; i < len(c.q); i++ {
				x := c.q[i]
				pq.Push(x)
				y := pq.Pop()
				if y != c.expect[i] {
					t.Fatalf("%s pop expected %d, but got %d",
						c.name, c.expect, y)
				}
			}
		})
	}
}

func TestPriorityQueue_New_Push_Pop(t *testing.T) {
	cases := []struct {
		name   string
		q      []interface{}
		expect []interface{}
	}{
		{"非空数组", []interface{}{10, 32, 2, 19, 28, 69}, []interface{}{2, 10, 19, 28, 32, 69}},
		{"非空数组", []interface{}{69, 32, 22, 19, 20, 11}, []interface{}{11, 19, 20, 22, 32, 69}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pq := New(c.q)
			idx := 0
			for !pq.IsEmpty() {
				y := pq.Pop()
				if y != c.expect[idx] {
					t.Fatalf("%s pop expected %d, but got %d",
						c.name, c.expect, y)
				}
				idx++
			}
		})
	}
}
