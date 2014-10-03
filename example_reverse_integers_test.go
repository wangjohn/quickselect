package quickselect_test

import (
  "fmt"
  "github.com/wangjohn/quickselect"
)

func ExampleReverseIntegers() {
  integers := []int{5, 2, 6, 3, 1, 4}
  quickselect.QuickSelect(quickselect.Reverse(quickselect.IntSlice(integers)))
  fmt.Println(integers[:3])
}
