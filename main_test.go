package main

import (
	"math/rand"
	"reflect"
	"sort-in-go/bubble"
	"sort-in-go/insertion"
	"testing"
)

func TestSort(t *testing.T) {
	assertion := func(t *testing.T, got, want []int) {
		t.Helper()
		// reflect.DeepEqual() compares two elements of any type recursively. Perfect for slices!
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %q\nWant: %q", got, want)
		}
	}

	t.Run("Bubble Sort successfully sorts [6,9,3,5]", func(t *testing.T) {
		got := insertion.Sort([]int{6, 9, 3, 5, 3})
		want := []int{3, 3, 5, 6, 9}
		assertion(t, got, want)
	})

	t.Run("Insertion Sort successfully sorts [6,9,3,5]", func(t *testing.T) {
		got := insertion.Sort([]int{6, 9, 3, 5, 3})
		want := []int{3, 3, 5, 6, 9}
		assertion(t, got, want)
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	s := make([]int, 10000)
	for i := range s {
		s[i] = rand.Intn(10000)
	}
	b.ResetTimer()
	bubble.Sort(s)
}

func BenchmarkInsertionSort(b *testing.B) {
	s := make([]int, 10000)
	for i := range s {
		s[i] = rand.Intn(10000)
	}
	b.ResetTimer()
	insertion.Sort(s)
}