package main
import "fmt"

func main() {
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
}
