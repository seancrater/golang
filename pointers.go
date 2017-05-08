package main
import "fmt"

func main() {
  i := 6
  p := &i

  fmt.Println(*p) // Reading value through pointer
}
