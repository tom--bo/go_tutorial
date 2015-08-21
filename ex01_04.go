package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "math/rand"
)

func main() {
    fmt.Println("Hello, 世界")
    fmt.Println("The time is", time.Now())
    fmt.Println("And if you try to open a file:")
    fmt.Println(os.Open("./ex02.go"))
    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "google.com"))
    fmt.Println("My favorite number is ", rand.Intn(10))
}
