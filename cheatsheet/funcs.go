package main
import "fmt"

func foo() string { return "foo" }
func abStr(a string, b string) string { return a + b }
func join(strList []string) string {
  return ""
}

func main() {
  fmt.Println("foo:", foo())
  fmt.Println("abStr", abStr("foo", "bar"))

  var strList := string[]{ "foo", "bar", "baz" }
  fmt.Println("join", join(strList})
}
