package test

import (
	"testing"
	"strconv"
	"fmt"
)

const SLICE_COUNT =  1000

func initSlice() []string{
	s:=make([]string, SLICE_COUNT)
	for i:=0;i< SLICE_COUNT;i++{
		s[i]=strconv.Itoa(i)
	}
	return s;
}


func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}


func BenchmarkForSlice(b *testing.B) {
	s:=initSlice()

	b.ResetTimer()
	for i:=0; i<b.N;i++  {
		ForSlice(s)
	}
}


func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}



func BenchmarkRangeForSlice(b *testing.B) {
	s:=initSlice()

	b.ResetTimer()
	for i:=0; i<b.N;i++  {
		RangeForSlice(s)
	}
}

func RangeForSlice2(s []string) {
	for i, _ := range s {
		_,_ = i,s[i]
	}
}

func BenchmarkRangeForSlice2(b *testing.B) {
	s:=initSlice()

	b.ResetTimer()
	for i:=0; i<b.N;i++  {
		RangeForSlice2(s)
	}
}


func RangeForMap1(m map[int]string) {
	for k, v := range m {
		_, _ = k, v
	}
}

const N = 1000

func initMap() map[int]string {
	m := make(map[int]string, N)
	for i := 0; i < N; i++ {
		m[i] = fmt.Sprint("foo",i)
	}
	return m
}

func BenchmarkRangeForMap1(b *testing.B) {
	m:=initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap1(m)
	}
}


func RangeForMap2(m map[int]string) {
	for k, _ := range m {
		_, _ = k, m[k]
	}
}

func BenchmarkRangeForMap2(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap2(m)
	}
}


func RangeForMap3(m map[int]string) {
	for _,_=range m {

	}
}

func BenchmarkRangeForMap3(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap3(m)
	}
}