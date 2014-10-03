package quickselect_test

import (
  "fmt"
  "github.com/wangjohn/quickselect"
)

func ExampleIntegers() {
  integers := []int{5, 2, 6, 3, 1, 4}
  quickselect.QuickSelect(quickselect.IntSlice(integers), 3)
  fmt.Println(integers[:3])
}
