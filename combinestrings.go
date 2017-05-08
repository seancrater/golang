package main
import "fmt"

func viaArray() {
  var glued [2]string
  glued[0] = "Hello"
  glued[1] = "World!"

  fmt.Println(glued)
}

func viaStr() {
  var glued = "Hello"
  glued += " World!"

  fmt.Println(glued)
}

func main() {
  viaArray()
  viaStr()
}
