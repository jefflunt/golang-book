package main
import "fmt"

func main() {
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
}
