package bigset

import (
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
