package main
import "fmt"
import "reflect"

func main() {
  fmt.Println("\nLogic:")
  fmt.Println("true && false :", true && false)
  fmt.Println("true && true :", true && true)
  fmt.Println("false && false :", false && false)

  fmt.Println("!true && false :", !true && false)
  fmt.Println("!true && true :", !true && true)
  fmt.Println("!false && !false :", !false && !false)

  fmt.Println("true || false :", true || false)
  fmt.Println("true || true :", true || true)
  fmt.Println("false || false :", false || false)
  fmt.Println("(across non-bool types) false || nil : compiler type error")

  fmt.Println("\nTypes, declariation, assignment:")
  fmt.Println("5 == 3.0 :", reflect.TypeOf(5 == 3.0))

  i := 56
  fmt.Println("(declare and assign, inferred type) i := 56 :", reflect.TypeOf(i))

  j := 123.0
  fmt.Println("(declare and assign, inferred type) j := 123.0 :", reflect.TypeOf(j))

  k := 5 * 3.0
  fmt.Println("(declare and assign, inferred type) k := 5 * 3.0 :", reflect.TypeOf(k))
}
