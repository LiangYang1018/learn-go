package main

import "os"

func main() {

    //panic("a problem")

    _, err := os.Create("file1.txt")
    if err != nil {
        panic(err)
    }
}