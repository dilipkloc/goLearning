package main
import (
    "fmt"
    "strings"
)

func main(){
    var s string
    fmt.Printf("Please enter a string:\n")
    fmt.Scan(&s)
    s = strings.ToLower(s)
    if strings.HasPrefix(s, "i")  && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
        fmt.Printf("Found!")
    } else {
        fmt.Printf("Not Found!")
    }
}
