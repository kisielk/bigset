package bigset

import (
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	s := New()
	if s.Contains(0) {
		t.Fatal("set should not contain 0")
	}
	if s.Len() != 0 {
		t.Fatal("Len() != 0")
	}

	s.Insert(0)
	if !s.Contains(0) {
		t.Fatal("set should contain 0")
	}
	if s.Len() != 1 {
		t.Fatal("Len() != 1")
	}

	s.Insert(0)
	if !s.Contains(0) {
		t.Fatal("set should contain 0")
	}
	if s.Len() != 1 {
		t.Fatal("Len() != 1")
	}

	s.Remove(0)
	if s.Contains(0) {
		t.Fatal("set should not contain 0")
	}
	if s.Len() != 0 {
		t.Fatal("Len() != 0")
	}
}

func TestUnion(t *testing.T) {
	a := New(1, 2, 3)
	b := New(3, 4, 5)
	c := a.Union(b)
	if l := c.Len(); l != 5 {
		t.Fatalf("Len(): got %d, want %d", l, 5)
	}

	for _, n := range []int{1, 2, 3, 4, 5} {
		if !c.Contains(n) {
			t.Errorf("union does not contain %d", n)
		}
	}
}

func TestIntersection(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(3, 4, 5)
	c := a.Intersection(b)
	if l := c.Len(); l != 2 {
		t.Fatalf("Len(): got %d, want %d", l, 2)
	}

	for _, n := range []int{3, 4} {
		if !c.Contains(n) {
			t.Errorf("intersection does not contain %d", n)
		}
	}
}

func TestDifference(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(1, 2)
	c := a.Difference(b)
	if l := c.Len(); l != 2 {
		t.Fatalf("Len(): got %d, want %d", l, 2)
	}

	for _, n := range []int{3, 4} {
		if !c.Contains(n) {
			t.Errorf("difference does not contain %d", n)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(1, 4, 5)
	c := a.SymmetricDifference(b)
	if l := c.Len(); l != 3 {
		t.Fatalf("Len(): got %d, want %d", l, 3)
	}

	for _, n := range []int{2, 3, 5} {
		if !c.Contains(n) {
			t.Errorf("symmetric difference does not contain %d", n)
		}
	}
}

func TestSuperSubSet(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(1, 4)
	c := New(1, 5)
	d := New(1, 4)

	super := []struct {
		a, b    *Set
		isSuper bool
	}{
		{a, b, true},
		{a, c, false},
		{b, a, false},
		{b, c, false},
		{c, a, false},
		{c, b, false},
		{b, d, true},
		{d, b, true},
	}
	for n, test := range super {
		if isSuper := test.a.IsSuperset(test.b); isSuper != test.isSuper {
			t.Errorf("%d: IsSuperset got %v, want %v", n, isSuper, test.isSuper)
		}
		if isSub := test.b.IsSubset(test.a); isSub != test.isSuper {
			t.Errorf("%d: IsSubset got %v, want %v", n, isSub, test.isSuper)
		}
	}
}

func BenchmarkIntersection10000_1000(b *testing.B)    { benchmarkIntersection(b, 10000, 1000) }
func BenchmarkIntersection100000_1000(b *testing.B)   { benchmarkIntersection(b, 100000, 1000) }
func BenchmarkIntersection1000000_1000(b *testing.B)  { benchmarkIntersection(b, 1000000, 1000) }
func BenchmarkIntersection10000000_1000(b *testing.B) { benchmarkIntersection(b, 10000000, 1000) }

func BenchmarkIntersection10000_10000(b *testing.B)    { benchmarkIntersection(b, 10000, 10000) }
func BenchmarkIntersection100000_10000(b *testing.B)   { benchmarkIntersection(b, 100000, 10000) }
func BenchmarkIntersection1000000_10000(b *testing.B)  { benchmarkIntersection(b, 1000000, 10000) }
func BenchmarkIntersection10000000_10000(b *testing.B) { benchmarkIntersection(b, 10000000, 10000) }

func benchmarkIntersection(b *testing.B, max, nNumbers int) {
	setA := New()
	setB := New()
	for i := 0; i < nNumbers; i++ {
		setA.Insert(rand.Intn(max))
		setB.Insert(rand.Intn(max))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Intersection(setB)
	}
}

func BenchmarkUnion10000_1000(b *testing.B)    { benchmarkUnion(b, 10000, 1000) }
func BenchmarkUnion100000_1000(b *testing.B)   { benchmarkUnion(b, 100000, 1000) }
func BenchmarkUnion1000000_1000(b *testing.B)  { benchmarkUnion(b, 1000000, 1000) }
func BenchmarkUnion10000000_1000(b *testing.B) { benchmarkUnion(b, 10000000, 1000) }

func BenchmarkUnion10000_10000(b *testing.B)    { benchmarkUnion(b, 10000, 10000) }
func BenchmarkUnion100000_10000(b *testing.B)   { benchmarkUnion(b, 100000, 10000) }
func BenchmarkUnion1000000_10000(b *testing.B)  { benchmarkUnion(b, 1000000, 10000) }
func BenchmarkUnion10000000_10000(b *testing.B) { benchmarkUnion(b, 10000000, 10000) }

func benchmarkUnion(b *testing.B, max, nNumbers int) {
	setA := New()
	setB := New()
	for i := 0; i < nNumbers; i++ {
		setA.Insert(rand.Intn(max))
		setB.Insert(rand.Intn(max))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Union(setB)
	}
}

func BenchmarkDifference10000_1000(b *testing.B)    { benchmarkDifference(b, 10000, 1000) }
func BenchmarkDifference100000_1000(b *testing.B)   { benchmarkDifference(b, 100000, 1000) }
func BenchmarkDifference1000000_1000(b *testing.B)  { benchmarkDifference(b, 1000000, 1000) }
func BenchmarkDifference10000000_1000(b *testing.B) { benchmarkDifference(b, 10000000, 1000) }

func BenchmarkDifference10000_10000(b *testing.B)    { benchmarkDifference(b, 10000, 10000) }
func BenchmarkDifference100000_10000(b *testing.B)   { benchmarkDifference(b, 100000, 10000) }
func BenchmarkDifference1000000_10000(b *testing.B)  { benchmarkDifference(b, 1000000, 10000) }
func BenchmarkDifference10000000_10000(b *testing.B) { benchmarkDifference(b, 10000000, 10000) }

func benchmarkDifference(b *testing.B, max, nNumbers int) {
	setA := New()
	setB := New()
	for i := 0; i < nNumbers; i++ {
		setA.Insert(rand.Intn(max))
		setB.Insert(rand.Intn(max))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.Difference(setB)
	}
}

func BenchmarkSymmetricDifference10000_1000(b *testing.B) {
	benchmarkSymmetricDifference(b, 10000, 1000)
}
func BenchmarkSymmetricDifference100000_1000(b *testing.B) {
	benchmarkSymmetricDifference(b, 100000, 1000)
}
func BenchmarkSymmetricDifference1000000_1000(b *testing.B) {
	benchmarkSymmetricDifference(b, 1000000, 1000)
}
func BenchmarkSymmetricDifference10000000_1000(b *testing.B) {
	benchmarkSymmetricDifference(b, 10000000, 1000)
}

func BenchmarkSymmetricDifference10000_10000(b *testing.B) {
	benchmarkSymmetricDifference(b, 10000, 10000)
}
func BenchmarkSymmetricDifference100000_10000(b *testing.B) {
	benchmarkSymmetricDifference(b, 100000, 10000)
}
func BenchmarkSymmetricDifference1000000_10000(b *testing.B) {
	benchmarkSymmetricDifference(b, 1000000, 10000)
}
func BenchmarkSymmetricDifference10000000_10000(b *testing.B) {
	benchmarkSymmetricDifference(b, 10000000, 10000)
}

func benchmarkSymmetricDifference(b *testing.B, max, nNumbers int) {
	setA := New()
	setB := New()
	for i := 0; i < nNumbers; i++ {
		setA.Insert(rand.Intn(max))
		setB.Insert(rand.Intn(max))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setA.SymmetricDifference(setB)
	}
}
