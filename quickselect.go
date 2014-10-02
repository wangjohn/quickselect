package quickselect

import (
  "math/rand"
  "errors"
  "fmt"
)

type Interface interface {
  // Len is the number of elements in the collection
  Len() int
  // Less reports whether the element with
  // index i should sort before the element with index j
  Less(i, j int) bool
  // Swap swaps the order of elements i and j
  Swap(i, j int)
}

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

func (t IntSlice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

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

func (t Float64Slice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

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

func (t StringSlice) QuickSelect(k int) (error) {
  return QuickSelect(t, k)
}

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
  return f != f
}

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

func QuickSelect(data Interface, k int) (error) {
  length := data.Len()
  if (k < 0 || k >= length) {
    message := fmt.Sprintf("The specified index '%d' is outside of the data's range of indices [0,%d)", k, length)
    return errors.New(message)
  }
  randomizedMedianFinding(data, 0, length - 1, k)
  return nil
}
