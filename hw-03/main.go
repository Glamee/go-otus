package main

import "fmt"

func main() {
    var size uint

    fmt.Printf("Enter a size of the desk: ")
    _, err := fmt.Scan(&size)
    if err != nil { 
        panic("Invalid value")
    }

    for i := range size {
        for j := range size {
            if (j+i)%2 == 0 {
                fmt.Printf("#")
            } else {
                fmt.Printf(" ")
            }
        }
        fmt.Println("")
    }
}