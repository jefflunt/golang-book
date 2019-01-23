package main
import "fmt"

func main() {
  fmt.Println("Double  :\n  Whatever, man. I do what I want.")
  fmt.Println(`Backtick:\nWhatever, man. I do what I want.`)
  fmt.Println(`
    Whatever, man.
    I do what I want.
  `)
  // This \/ doesn't work.
  // Println('Single  : Whatever, man. I do what I want.')

  fmt.Println("Char    : Whatever, man. I do what I want."[10]) // => "W", zero-based indexing
  fmt.Println("Concat  : Whatever, man." + " I do what I want.")
  fmt.Println("Series  : Whatever, man.", "I do what I want.")
}
