package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	fh := NewQueue(100)

	if fh == nil {
		t.Error("TestNew Failed")
	}

	if fh.size != 100 {
		t.Errorf("Mismatch Size")
	}

}

func TestPush(t *testing.T) {
	fh := NewQueue(1)

	obj := 1

	fh.Push(obj)

	if fh.Len() != 1 {
		t.Errorf("Can't Push Queue")
	}
}

func TestPeek(t *testing.T) {
	fh := NewQueue(1)

	obj := 1

	fh.Push(obj)

	peekedObj := fh.Peek()

	if peekedObj != 1 {
		t.Errorf("Unexpected Result", peekedObj)
	}

	if fh.Len() != 1 {
		t.Errorf("Unexpected Queue Len %v", fh.Len())
	}
}

func TestPop(t *testing.T) {
	fh := NewQueue(1)

	obj := 1
	fh.Push(obj)

	obj2 := 2
	fh.Push(obj2)

	obj3 := 3
	fh.Push(obj3)

	obj4 := 4
	fh.Push(obj4)

	obj5 := 5
	fh.Push(obj5)

	t.Logf("queue")
	for i, v := range fh.nodes {
		t.Logf("%v %#v", i, v)
	}

	retObj := fh.Pop()

	if retObj != 1 {
		t.Errorf("wrong object %v", retObj)
	}
}

func BenchmarkPush_100000(b *testing.B) {
	b.ResetTimer()

	fh := NewQueue(100000)
	obj := 5

	for n := 0; n < b.N; n++ {
		fh.Push(obj)

	}
}

func BenchmarkPush_1000000(b *testing.B) {
	b.ResetTimer()

	fh := NewQueue(1000000)
	obj := 10

	for n := 0; n < b.N; n++ {
		fh.Push(obj)

	}
}

func BenchmarkPush_10000000(b *testing.B) {
	b.ResetTimer()

	fh := NewQueue(10000000)
	obj := 15

	for n := 0; n < b.N; n++ {
		fh.Push(obj)

	}
}
