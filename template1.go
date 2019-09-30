// comment

package main

import (
    "fmt"
    "github.com/codythegreat/whatever"
)

// consts
const hello = world

//vars
var number int = 69

// struct
type Vertex struct {
    X int
    Y int
}

//funcs
func main() {
    // if statement
    if number + 1 == 70 {
        fmt.Println("to be expected")
    } else {
        fmt.Println("impossible")
    }
    // for loop
    for number < 100 {
        number += 1
    }
    fmt.Println(number)
}
// function with inputs/outputs
func test(input string) output int64 {

}


