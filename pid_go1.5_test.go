// +build gc,go1.5

package goid

import (
	"fmt"
	"testing"
	"unsafe"
)

var _ = unsafe.Sizeof(0)

//go:linkname procPin runtime.procPin
//go:nosplit
func procPin() int

//go:linkname procUnpin runtime.procUnpin
//go:nosplit
func procUnpin()

func getTID() int {
	tid := procPin()
	procUnpin()
	return tid
}

func TestParallelGetPid(t *testing.T) {
	ch := make(chan *string, 100)
	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			id := GetPid()
			expected := getTID()
			if id == expected {
				ch <- nil
				return
			}
			s := fmt.Sprintf("Expected %d, but got %d", expected, id)
			ch <- &s
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		val := <-ch
		if val != nil {
			t.Fatal(*val)
		}
	}
}

func TestGetPid(t *testing.T) {
	p1 := GetPid()
	p2 := getTID()
	if p1 != p2 {
		t.Fatalf("The result of GetPid %d procPin %d are not equal!", p1, p2)
	}
}

func BenchmarkGetPid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPid()
	}
}

func BenchmarkParallelGetPid(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetPid()
		}
	})
}
