package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for i, arg := range os.Args[:] {
        s += sep + arg
        sep = " "
		fmt.Println(i,arg)
    }
    fmt.Println(s)

}