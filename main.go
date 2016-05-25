package main

import (
    "math/rand"
    "strconv"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) >= 2 {
        max, err := strconv.Atoi(os.Args[1])
        if err == nil {
            fmt.Printf("rand.Intn(%v) = %v\n", os.Args[1], rand.Intn(max))
            fmt.Printf("rand.Intn(%v) = %v\n", 200, rand.Intn(200))
            fmt.Printf("rand.Intn(%v) = %v\n", 300, rand.Intn(300))
            os.Exit(0)
        }
    }
    fmt.Printf("Usage: %v <randint-max>\n", os.Args[0])
    fmt.Printf("\n")
    fmt.Printf("       randint-max   = Generate random number between 0 and randint-max\n")
    fmt.Printf("\n")
    os.Exit(1)   
}
