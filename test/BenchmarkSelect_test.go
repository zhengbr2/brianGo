package test

import "testing"

func BenchmarkSelectShared(b *testing.B) {
	idleShared := make(chan struct{})
	b.RunParallel(func(pb *testing.PB) {
		ch := make(chan int, 1)
		for pb.Next() {
			select {
			case ch <- 1:
			case <-ch:
			case <-idleShared:
			}
		}
	})
}


func BenchmarkSelectShared2(b *testing.B) {

	b.ResetTimer()
	idleShared := make(chan struct{})
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1)

			select {
			case ch <- 1:
			case <-ch:
			case <-idleShared:
			}

	}
}

func BenchmarkSelectPrivate(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		ch := make(chan int, 1)
		idlePrivate := make(chan struct{})
		for pb.Next() {
			select {
			case ch <- 1:
			case <-ch:
			case <-idlePrivate:
			}
		}
	})
}
