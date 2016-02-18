package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var avg = []int{}
	file := os.Args[1]
	scores := bufio.NewScanner(file)
	for scores.Scan() {
		fmt.Println(scores.Text())
	}
	fmt.Println(len(avg))
}
