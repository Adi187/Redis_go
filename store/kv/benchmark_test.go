package kv

import (
	"strconv"
	"sync"
	"testing"
)

// One benchmark operation is a batch of exactly 5,000 reader goroutines.
// A barrier ensures all readers are created before any start calling Get.
// Timing includes goroutine creation, scheduling, Get, and waiting for completion.
func BenchmarkGet5000Readers(b *testing.B) {
	const readers = 5000
	s := NewStore(0)
	s.data["key"] = "value"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(readers)
		for j := 0; j < readers; j++ {
			go func() {
				defer wg.Done()
				<-start
				value, ok, err := s.Get("key")
				if err != nil || !ok || value != "value" {
					b.Error("unexpected Get result", value, ok, err)
				}
			}()
		}
		close(start)
		wg.Wait()
	}
	b.StopTimer()
	b.ReportMetric(float64(b.N)*readers/b.Elapsed().Seconds(), "reads/s")
}

// These benchmarks include the artificial delays in Get and Incr.
// Setup uses the private map to keep Set's console logging out of the timings.
func BenchmarkGet(b *testing.B) {
	for _, parallel := range []bool{false, true} {
		name := "Serial"
		if parallel {
			name = "Parallel"
		}
		b.Run(name, func(b *testing.B) {
			s := NewStore(0)
			s.data["key"] = "value"
			b.ReportAllocs()
			b.ResetTimer()
			read := func() {
				value, ok, err := s.Get("key")
				if err != nil || !ok || value != "value" {
					b.Error("unexpected Get result", value, ok, err)
				}
			}
			if parallel {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						read()
					}
				})
			} else {
				for i := 0; i < b.N; i++ {
					read()
				}
			}
		})
	}
}

func BenchmarkIncr(b *testing.B) {
	for _, parallel := range []bool{false, true} {
		name := "Serial"
		if parallel {
			name = "ParallelSharedKey"
		}
		b.Run(name, func(b *testing.B) {
			s := NewStore(0)
			s.data["counter"] = "0"
			b.ReportAllocs()
			b.ResetTimer()
			incr := func() {
				if _, err := s.Incr("counter"); err != nil {
					b.Error(err)
				}
			}
			if parallel {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						incr()
					}
				})
			} else {
				for i := 0; i < b.N; i++ {
					incr()
				}
			}
			b.StopTimer()
			if got := s.data["counter"]; got != strconv.Itoa(b.N) {
				b.Fatalf("counter = %s, want %d", got, b.N)
			}
		})
	}
}
