package main
import "fmt"
import "reflect"

func main() {
  fmt.Println("Basic output:")
  fmt.Println("Hello Go")

  fmt.Println("\nArithmetic:")
  fmt.Println("5 + 5 :", 5 + 5)
  fmt.Println("5 - 3 :", 5 - 3)
  fmt.Println("5 / 2 :", 5 / 2)
  fmt.Println("5 / 2.0 :", 5 / 2.0)
  fmt.Println("5 * 3 :", 5 * 3)
  fmt.Println("5 * 3.0 :", 5 * 3.0)
  fmt.Println("5.0 * 3.0 :", 5.0 * 3.0)
  fmt.Println("5.1 * 3.2 :", 5.1 * 3.2)
  fmt.Println("5 % 2 :", 5 % 2)
  fmt.Println("(bitwise and) 5 & 2 :", 5 & 2)
  fmt.Println("(bitwise and) 6 & 2 :", 6 & 2)
  fmt.Println("(bitwise or)  5 | 2 :", 5 | 2)
  fmt.Println("(bitwise or)  6 & 2 :", 6 | 2)
  fmt.Println("(bit clear -- and not)  5 &^ 2 :", 5 &^ 2)
  fmt.Println("(bit clear -- and not)  2 &^ 5 :", 2 &^ 5)
  fmt.Println("(bit clear -- and not)  6 &^ 2 :", 6 &^ 2)
  fmt.Println("(bit clear -- and not)  2 &^ 6 :", 2 &^ 6)
  fmt.Println("(right shift) 8 >> 1 :", 8 >> 1)
  fmt.Println("(right shift) 8 >> 2 :", 8 >> 2)
  fmt.Println("(right shift) 8 >> 10 :", 8 >> 10)
  fmt.Println("(left shift) 8 << 1 :", 8 << 1)
  fmt.Println("(left shift) 8 << 2 :", 8 << 2)
  fmt.Println("(left shift) 8 << 10 :", 8 << 10)

  fmt.Println("\nEquality:")
  fmt.Println("(across incompatible types) 5 == true : compiler type error")
  fmt.Println("(across numeric types) 5 == 3.0 :", 5 == 3.0)
  fmt.Println("5 == 3 :", 5 == 3)
  fmt.Println("5 == 5 :", 5 == 5)
  fmt.Println("5 != 3 :", 5 != 3)
  fmt.Println("5 != 5 :", 5 != 5)
  fmt.Println("5 <= 5 :", 5 <= 5)
  fmt.Println("5 <= 3 :", 5 <= 3)
  fmt.Println("5 >= 7 :", 5 >= 7)
  fmt.Println("5 >= 5 :", 5 >= 5)
  fmt.Println("5 >= 3 :", 5 >= 3)

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

  fmt.Println("\nPointers")
  var p *int
  p = &i
  fmt.Println("(declare w/*)              var p *int :")
  fmt.Println("(assign p to address of i) p = &i")
  fmt.Println("(value of p, i.e. address) p:", p)
  fmt.Println("(dereference)              *p:", *p)
}
