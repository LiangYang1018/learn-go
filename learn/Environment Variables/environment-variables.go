package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {

    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println("GO111MODULE:", os.Getenv("GO111MODULE"))
    x, err := os.Hostname()
    
    if err != nil {
		logrus.Fatalf("failed to get hostname: %v", err)
	}
    fmt.Println(x)
    os.Exit(0)
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}