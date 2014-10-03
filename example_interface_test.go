package quickselect_test

import (
  "quickselect"
  "fmt"
)

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person
func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) { return a[i].Age < a[j].Age }

integerArray := []IntSlice{10, 5, 3, 15, 4, 9, 3, 4}

func Example() {
  people := []Person{
      {"Bob", 31},
      {"John", 42},
      {"Michael", 17},
      {"Jenny", 26},
  }

  quickselect.QuickSelect(ByAge(people), 2)
  fmt.Println(people[:2])
}
