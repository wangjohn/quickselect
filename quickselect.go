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
  Swap(i, j int)
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

func QuickSelect(data Interface, k int) {
  randomizedMedianFinding(data, 0, data.Len() - 1, k)
}
