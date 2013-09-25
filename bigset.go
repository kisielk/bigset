package bigset

import (
	"math/big"
)

type Set big.Int

func NewSet(n ...int) *Set {
	s := (*Set)(big.NewInt(0))
	s.setBits(1, n...)
	return s
}

func (s *Set) setBits(b uint, bits ...int) {
	i := (*big.Int)(s)
	for _, n := range bits {
		i = i.SetBit(i, n, b)
	}
	*s = *(*Set)(i)
}

func (s *Set) Insert(n ...int) {
	s.setBits(1, n...)
}

func (s *Set) Remove(n ...int) {
	s.setBits(0, n...)
}

func (s *Set) Contains(n int) bool {
	return (*big.Int)(s).Bit(n) == uint(1)
}

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

func (s *Set) Union(t *Set) *Set {
	return (*Set)(new(big.Int).Or((*big.Int)(s), (*big.Int)(t)))
}

func (s *Set) Intersection(t *Set) *Set {
	return (*Set)(new(big.Int).And((*big.Int)(s), (*big.Int)(t)))
}

func (s *Set) Difference(t *Set) *Set {
	return (*Set)(new(big.Int).And((*big.Int)(s), new(big.Int).Not((*big.Int)(t))))
}

func (s *Set) SymmetricDifference(t *Set) *Set {
	return (*Set)(new(big.Int).Xor((*big.Int)(s), (*big.Int)(t)))
}
