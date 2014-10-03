package quickselect_test

import (
  "fmt"
  "github.com/wangjohn/quickselect"
)

func ExampleQuickIntegers() {
  integers := []int{5, 2, 6, 3, 1, 4}
  quickselect.IntQuickSelect(integers, 3)
  fmt.Println(integers[:3])
}

func ExampleIntegers() {
  integers := []int{5, 2, 6, 3, 1, 4}
  quickselect.QuickSelect(quickselect.IntSlice(integers), 3)
  fmt.Println(integers[:3])
}

func ExampleReverseIntegers() {
  integers := []int{5, 2, 6, 3, 1, 4}
  quickselect.QuickSelect(quickselect.Reverse(quickselect.IntSlice(integers)))
  fmt.Println(integers[:3])
}
