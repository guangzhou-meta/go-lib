package collections

import "github.com/guangzhou-meta/go-lib/tools"

// 非线程安全
type priorityQueue struct {
	h   []interface{}
	cmp tools.Comparator
}

// New 默认小根堆
// 通过传入的数组建堆，时间复杂度 O(n)，如果不断 push 是nlogn
func New(q []interface{}) *priorityQueue {
	n := len(q)
	pq := new(priorityQueue)
	if n == 0 {

	} else {
		pq.cmp = tools.GetCmp(q[0])
		pq.h = q
		for i := n / 2; i >= 0; i-- {
			pq.down(i)
		}
	}
	return pq
}

func (pq *priorityQueue) down(u int) {
	t := u
	if u*2+1 < len(pq.h) && pq.cmp(pq.h[t], pq.h[u*2+1]) > 0 {
		t = u*2 + 1
	}
	if u*2+2 < len(pq.h) && pq.cmp(pq.h[t], pq.h[u*2+2]) > 0 {
		t = u*2 + 2
	}
	if t != u {
		pq.h[t], pq.h[u] = pq.h[u], pq.h[t]
		pq.down(t)
	}
}

func (pq *priorityQueue) up(u int) {
	for u/2 >= 0 && pq.cmp(pq.h[u/2], pq.h[u]) > 0 {
		pq.h[u/2], pq.h[u] = pq.h[u], pq.h[u/2]
		u /= 2
	}
}

func (pq *priorityQueue) Push(x interface{}) {
	if pq.cmp == nil {
		pq.cmp = tools.GetCmp(x)
	}
	if pq.cmp == nil {
		panic("unknown comparator")
	}
	pq.h = append(pq.h, x)
	pq.up(len(pq.h) - 1)
}

func (pq *priorityQueue) Pop() (res interface{}) {
	res = pq.h[0]
	pq.h[0] = pq.h[len(pq.h)-1]
	pq.h = pq.h[:len(pq.h)-1]
	pq.down(0)
	return
}

func (pq *priorityQueue) Peek() (res interface{}) {
	return pq.h[0]
}

func (pq *priorityQueue) IsEmpty() bool {
	return len(pq.h) == 0
}
