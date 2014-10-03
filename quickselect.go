/*
The quickselect package provides primitives for finding the smallest k elements
in slices and user-defined collections. The primitives used in the package are
modeled off of the standard sort library for Go. Quickselect uses Hoare's
Selection Algorithm which finds the smallest k elements in expected O(n) time,
and is thus an asymptotically optimal algorithm (and is faster than sorting or
heap implementations).
*/
package quickselect

import (
  "math/rand"
  "errors"
  "fmt"
)

/*
A type, typically a collection, which satisfies quickselect.Interface can be
used as data in the QuickSelect method. The interface is the same as the
interface required by Go's canonical sorting package (sort.Interface).

Note that the methods require that the elements of the collection be enumerated
by an integer index.
*/
type Interface interface {
  // Len is the number of elements in the collection
  Len() int
  // Less reports whether the element with
  // index i should sort before the element with index j
  Less(i, j int) bool
  // Swap swaps the order of elements i and j
  Swap(i, j int)
}

type reverse struct {
  // This embedded Interface permits Reverse to use the methods of
  // another Interface implementation.
  Interface
}

// Less returns the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
  return r.Interface.Less(j, i)
}

func Reverse(data Interface) (Interface) {
  return &reverse{data}
}

// The IntSlice type attaches the QuickSelect interface to an array of ints. It
// implements Interface so that you can call QuickSelect(k) on any IntSlice.
type IntSlice []int

func (t IntSlice) Len() (int) {
  return len(t)
}

func (t IntSlice) Less(i, j int) bool {
  return t[i] < t[j]
}

func (t IntSlice) Swap(i, j int) {
  t[i], t[j] = t[j], t[i]
}

// QuickSelect(k) mutates the IntSlice so that the first k elements in the
// IntSlice are the k smallest elements in the slice. This is a convenience
// method for QuickSelect
func (t IntSlice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

// The Float64Slice type attaches the QuickSelect interface to an array of
// float64s. It implements Interface so that you can call QuickSelect(k) on any
// Float64Slice.
type Float64Slice []float64

func (t Float64Slice) Len() (int) {
  return len(t)
}

func (t Float64Slice) Less(i, j int) bool {
  return t[i] < t[j] || isNaN(t[i]) && !isNaN(t[j])
}

func (t Float64Slice) Swap(i, j int) {
  t[i], t[j] = t[j], t[i]
}

// QuickSelect(k) mutates the Float64Slice so that the first k elements in the
// Float64Slice are the k smallest elements in the slice. This is a convenience
// method for QuickSelect
func (t Float64Slice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

// The StringSlice type attaches the QuickSelect interface to an array of
// float64s. It implements Interface so that you can call QuickSelect(k) on any
// StringSlice.
type StringSlice []string

func (t StringSlice) Len() (int) {
  return len(t)
}

func (t StringSlice) Less(i, j int) bool {
  return t[i] < t[j]
}

func (t StringSlice) Swap(i, j int) {
  t[i], t[j] = t[j], t[i]
}

// QuickSelect(k) mutates the StringSlice so that the first k elements in the
// StringSlice are the k smallest elements in the slice. This is a convenience
// method for QuickSelect
func (t StringSlice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
  return f != f
}

/*
Helper function that does all of the work for QuickSelect. This implements
Hoare's Selection Algorithm which finds the smallest k elements in an interface
in expected O(n) time.

The algorithm works by finding a random pivot element, and making sure all the
elements to the left are less than the pivot element and vice versa for
elements on the right. Recursing on this solves the selection algorithm.
*/
func randomizedMedianFinding(data Interface, low, high, k int) {
  var pivotIndex int

  for {
    if low >= high {
      return
    }
    pivotIndex = rand.Intn(high + 1 - low) + low
    pivotIndex = partition(data, low, high, pivotIndex)

    if k < pivotIndex {
      high = pivotIndex - 1
    } else if k > pivotIndex {
      low = pivotIndex + 1
    } else {
      return
    }
  }
}

/*
Helper function for the selection algorithm. Returns the partitionIndex.

It goes through all elements between low and high and makes sure that the
elements in the range [low, partitionIndex) are less than the element that was
originally in the pivotIndex and that the elements in the range
[paritionIndex + 1, high] are greater than the element originally in the
pivotIndex.
*/
func partition(data Interface, low, high, pivotIndex int) (int) {
  partitionIndex := low
  data.Swap(pivotIndex, high)
  for i := low; i < high; i++ {
    if data.Less(i, high) {
      data.Swap(i, partitionIndex)
      partitionIndex++
    }
  }
  data.Swap(partitionIndex, high)
  return partitionIndex
}

/*
QuickSelect swaps elements in the data provided so that the first k elements
(i.e. the elements occuping indices 0, 1, ..., k-1) are the smallest k elements
in the data.

QuickSelect implements Hoare's Selection Algorithm and runs in O(n) time, so it
is asymptotically faster than sorting or other heap-like implementations for
finding the smallest k elements in a data structure.

Note that k must be in the range [0, data.Len()), otherwise the QuickSelect
method will raise an error.
*/
func QuickSelect(data Interface, k int) (error) {
  length := data.Len()
  if (k < 0 || k >= length) {
    message := fmt.Sprintf("The specified index '%d' is outside of the data's range of indices [0,%d)", k, length)
    return errors.New(message)
  }
  randomizedMedianFinding(data, 0, length - 1, k)
  return nil
}
