// Iris - Decentralized Messaging Framework
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// Iris is dual licensed: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The framework is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the Iris framework may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)

// Package sortext contains extensions to the base Go sort package.
package sortext

import (
	"math/big"
	"sort"
)

// BigIntSlice attaches the methods of Interface to []*big.Int, sorting in increasing order.
type BigIntSlice []*big.Int

func (b BigIntSlice) Len() int           { return len(b) }
func (b BigIntSlice) Less(i, j int) bool { return b[i].Cmp(b[j]) < 0 }
func (b BigIntSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort is a convenience method.
func (b BigIntSlice) Sort() { sort.Sort(b) }

// BigRatSlice attaches the methods of Interface to []*big.Rat, sorting in increasing order.
type BigRatSlice []*big.Rat

func (b BigRatSlice) Len() int           { return len(b) }
func (b BigRatSlice) Less(i, j int) bool { return b[i].Cmp(b[j]) < 0 }
func (b BigRatSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort is a convenience method.
func (b BigRatSlice) Sort() { sort.Sort(b) }

// BigInts sorts a slice of *big.Ints in increasing order.
func BigInts(a []*big.Int) { sort.Sort(BigIntSlice(a)) }

// BigRats sorts a slice of *big.Rats in increasing order.
func BigRats(a []*big.Rat) { sort.Sort(BigRatSlice(a)) }

// BigIntsAreSorted tests whether a slice of *big.Ints is sorted in increasing order.
func BigIntsAreSorted(a []*big.Int) bool { return sort.IsSorted(BigIntSlice(a)) }

// BigRatsAreSorted tests whether a slice of *big.Rats is sorted in increasing order.
func BigRatsAreSorted(a []*big.Rat) bool { return sort.IsSorted(BigRatSlice(a)) }
