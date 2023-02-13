package main

import (  
    "fmt"
)

func main() {
	ch := make(chan int)
	ch <- 5
	value := <-ch
	fmt.Println(value)

}
