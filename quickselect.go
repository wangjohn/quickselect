package quickselect

import (
  "math/rand"
)

type Interface interface {
  // Len is the number of elements in the collection
  Len() int
  // Less reports whether the element with
  // index i should sort before the element with index j
  Less(i, j int) bool
  // Swap swaps the order of elements i and j
  Swap(i, j)
}

func randomizedMedianFinding(data Interface, low, high, k int) (Interface) {
  if low == high {
    return low
  }

  var pivotIndex int

  for {
    pivotIndex = Rand.Intn(high + 1 - low) + low
    pivotIndex = partition(data, low, high, pivotIndex)

    if k < pivotIndex {
      high = pivotIndex - 1
    } else if k > pivotIndex {
      low = pivotIndex + 1
    } else {
      return data
    }
  }
}

func partition(data Interface, low, high, pivotIndex int) (int) {
  partitionIndex := low
  data.Swap(pivotIndex, high)
  for i := low; i++; i<high {
    if data.Less(i, pivotIndex) {
      data.Swap(i, partitionIndex)
      partitionIndex++
    }
  }
  data.Swap(partitionIndex, high)
  return partitionIndex
}

func QuickSelect(data Interface, k int) (Interface) {
  return randomizedMedianFinding(data, 0, data.Len() - 1, k)
}
