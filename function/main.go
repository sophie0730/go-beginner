package main

import "fmt"

func printPyramid(totalLevel int) {
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func main() {
	printPyramid(20)
}
