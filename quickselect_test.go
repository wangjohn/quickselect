package quickselect

import (
  "testing"
)

type TestData struct {
  Array []int
}

func (t TestData) Len() (int) {
  return len(t.Array)
}

func (t TestData) Less(i, j int) bool {
  return t.Array[i] < t.Array[j]
}

func (t TestData) Swap(i, j int) {
  t.Array[i], t.Array[j] = t.Array[j], t.Array[i]
}

func TestQuickSelectWithSimpleArray(t *testing.T) {
  fixture := TestData{[]int{50, 20, 30, 25, 45, 2, 6, 10, 3, 4, 5}}
  err := QuickSelect(fixture, 5)
  if err != nil {
    t.Errorf("Shouldn't have raised error: '%s'", err.Error())
  }

  smallestK := fixture.Array[:5]
  expectedK := []int{2,3,4,5,6}
  if !hasSameElements(smallestK, expectedK) {
    t.Errorf("Expected smallest K elements to be '%s', but got '%s'", expectedK, smallestK)
  }
}

func TestQuickSelectWithRepeatedElements(t *testing.T) {
  fixture := TestData{[]int{2, 10, 5, 3, 2, 6, 2, 6, 10, 3, 4, 5}}
  err := QuickSelect(fixture, 5)
  if err != nil {
    t.Errorf("Shouldn't have raised error: '%s'", err.Error())
  }

  smallestK := fixture.Array[:5]
  expectedK := []int{2,2,2,3,3}
  if !hasSameElements(smallestK, expectedK) {
    t.Errorf("Expected smallest K elements to be '%s', but got '%s'", expectedK, smallestK)
  }
}

func TestQuickSelectEmptyDataStructure(t *testing.T) {
  fixture := TestData{[]int{}}
  err := QuickSelect(fixture, 0)
  if err == nil {
    t.Errorf("Should have raised error on index outside of array length.")
  }

  err = QuickSelect(fixture, 5)
  if err == nil {
    t.Errorf("Should have raised error on index outside of array length.")
  }

  err = QuickSelect(fixture, -1)
  if err == nil {
    t.Errorf("Should have raised error on index outside of array length.")
  }
}

func hasSameElements(array1, array2 []int) (bool) {
  elements := make(map[int]int)

  for _, elem1 := range array1 {
    elements[elem1]++
  }

  for _, elem2 := range array2 {
    elements[elem2]--
  }

  for _, count := range elements {
    if count != 0 {
      return false
    }
  }
  return true
}
