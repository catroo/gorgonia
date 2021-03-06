// Generated by: main
// TypeWriter: smallset
// Directive: +gen on Type

package gorgonia

import (
	"bytes"
	"fmt"
)

/*
The MIT License (MIT)

Copyright (c) 2016 Xuanyi Chew (chewxy [AT] gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

//TypeSet is a set of Type
type typeSet Types

// NewTypeSet creates a new set of Type, given an input of any Type
func newTypeSet(a ...Type) typeSet {
	var set typeSet

	for _, v := range a {
		set = set.Add(v)
	}

	return set
}

// ContainsALl determines if all the wanted items are already in set
func (set typeSet) ContainsAll(ws ...Type) bool {
	for _, w := range ws {
		if !set.Contains(w) {
			return false
		}
	}
	return true
}

// Add adds an item into the set, and then returns a new set. If the item already exists, it returns the current set
func (set typeSet) Add(item Type) typeSet {
	if set.Contains(item) {
		return set
	}
	set = append(set, item)
	return set
}

// IsSubSetOf checks if the current set is a subset of the other set.
func (set typeSet) IsSubsetOf(other typeSet) bool {
	if len(set) > len(other) {
		return false
	}

	for _, v := range set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// IsSupersetOf checks if the current set is a superset of the other set
func (set typeSet) IsSupersetOf(other typeSet) bool {
	return other.IsSubsetOf(set)
}

// Intersect performs an intersection between two sets - only items that exist in both are returned
func (set typeSet) Intersect(other typeSet) typeSet {
	retVal := make(typeSet, 0)
	for _, o := range other {
		if set.Contains(o) {
			retVal = append(retVal, o)
		}
	}
	return retVal
}

//Union joins both sets together, keeping only unique items
func (set typeSet) Union(other typeSet) typeSet {
	retVal := make(typeSet, len(set))
	copy(retVal, set)
	for _, o := range other {
		if !retVal.Contains(o) {
			retVal = append(retVal, o)
		}
	}
	return retVal
}

// Difference returns a new set with items in the current set but not in the other set.
// Equivalent to  (set - other)
func (set typeSet) Difference(other typeSet) typeSet {
	retVal := make(typeSet, 0)
	for _, v := range set {
		if !other.Contains(v) {
			retVal = append(retVal, v)
		}
	}
	return retVal
}

// SymmetricDifference is the set of items that is not in each either set.
func (set typeSet) SymmetricDifference(other typeSet) typeSet {
	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

// Equals compares two sets and checks if it is the same
func (set typeSet) Equals(other typeSet) bool {
	if len(set) != len(other) {
		return false
	}

	for _, v := range set {
		if !other.Contains(v) {
			return false
		}
	}

	return true
}

// String for stuff
func (set typeSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("TypeSet[")
	for i, v := range set {
		if i == len(set)-1 {
			fmt.Fprintf(&buf, "%v", v)
		} else {
			fmt.Fprintf(&buf, "%v, ", v)
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// Contains determines if an item is in the set already
func (set typeSet) Contains(w Type) bool {
	for _, v := range set {
		if typeEq(v, w) {
			return true
		}
	}
	return false
}
