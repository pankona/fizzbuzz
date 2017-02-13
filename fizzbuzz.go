package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func fizzBuzz(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i <= n; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
