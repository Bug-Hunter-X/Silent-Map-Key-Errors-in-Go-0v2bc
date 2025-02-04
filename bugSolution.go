```go
package main

import (


        "fmt"
)

func main() {
        m := make(map[string]int)
        value, ok := m["b"]
        if ok {
                fmt.Println("Value found:", value)
        } else {
                fmt.Println("Key not found")
        }
}
```