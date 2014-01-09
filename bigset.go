// Package bigset implements a set type for storing non-negative integers
// using the Go standard library's math/big package.
//
// Most of the set operations return a *Set much like the operations in
// math/big return a *big.Int.
package bigset

import (
	"math/big"
)

type Set big.Int

// New constructs a new set with each of the elements n.
// Each n must be >= 0.
func New(n ...int) *Set {
	s := (*Set)(big.NewInt(0))
	s.setBits(1, n...)
	return s
}

func (s *Set) setBits(b uint, bits ...int) *Set {
	i := (*big.Int)(s)
	for _, n := range bits {
		i = i.SetBit(i, n, b)
	}
	*s = *(*Set)(i)
	return s
}

// Insert inserts each of the elements n in to the set.
// Each n must be >= 0.
func (s *Set) Insert(n ...int) *Set {
	return s.setBits(1, n...)
}

// Remove removes each of the elements n from the set, if it contains them.
// If the set does not contain an element n, Remove is a no-op for that element.
// Each n must be >= 0.
func (s *Set) Remove(n ...int) *Set {
	return s.setBits(0, n...)
}

// Contains returns true if the element n is a member of the set.
func (s *Set) Contains(n int) bool {
	return (*big.Int)(s).Bit(n) == uint(1)
}

// Len returns the number of elements in the set.
// Its complexity is O(n).
func (s *Set) Len() int {
	var l int
	zero := big.NewInt(0)
	v := new(big.Int).Set((*big.Int)(s))
	for l = 0; v.Cmp(zero) != 0; l++ {
		vMinusOne := new(big.Int).Sub(v, big.NewInt(1))
		v.And(v, vMinusOne)
	}
	return l
}

func union(z, s, t *Set) *Set {
	i := (*big.Int)(z)
	return (*Set)(i.Or((*big.Int)(s), (*big.Int)(t)))
}

// Union constructs a new set that is the result of s ∪ t.
func Union(s, t *Set) *Set {
	return union(new(Set), s, t)
}

// Union updates s to be the result of s ∪ t.
func (s *Set) Union(t *Set) *Set {
	return union(s, s, t)
}

func intersection(z, s, t *Set) *Set {
	i := (*big.Int)(z)
	return (*Set)(i.And((*big.Int)(s), (*big.Int)(t)))
}

// Intersection constructs a new set that is the result of s ∩ t.
func Intersection(s, t *Set) *Set {
	return intersection(new(Set), s, t)
}

// Intersection updates s to be the result of s ∩ t.
func (s *Set) Intersection(t *Set) *Set {
	return intersection(s, s, t)
}

func difference(z, s, t *Set) *Set {
	i := (*big.Int)(z)
	return (*Set)(i.And((*big.Int)(s), new(big.Int).Not((*big.Int)(t))))
}

// Difference constructs a new set that is the result of s ∖ t.
func Difference(s, t *Set) *Set {
	return difference(new(Set), s, t)
}

// Difference updates s to be the result of s ∖ t.
func (s *Set) Difference(t *Set) *Set {
	return difference(s, s, t)
}

func symmetricDifference(z, s, t *Set) *Set {
	i := (*big.Int)(z)
	return (*Set)(i.Xor((*big.Int)(s), (*big.Int)(t)))
}

// SymmetricDifference constructs a new set that is the result of s ∆ t.
func SymmetricDifference(s, t *Set) *Set {
	return symmetricDifference(new(Set), s, t)
}

// SymmetricDifference updates s to be the result of s ∆ t.
func (s *Set) SymmetricDifference(t *Set) *Set {
	return symmetricDifference(s, s, t)
}

// IsSubset returns true if s is a subset of t.
func (s *Set) IsSubset(t *Set) bool {
	return Intersection(s, t).Len() == s.Len()
}

// IsSuperset returns true if s is a superset of t.
func (s *Set) IsSuperset(t *Set) bool {
	return Intersection(s, t).Len() == t.Len()
}
