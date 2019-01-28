package main
import "fmt"

func main() {
  fmt.Println("\nPointers")
  i := 56
  var p *int
  p = &i
  fmt.Println("(declare w/*)              var p *int :")
  fmt.Println("(assign p to address of i) p = &i")
  fmt.Println("(value of p, i.e. address) p:", p)
  fmt.Println("(dereference)              *p:", *p)
}
