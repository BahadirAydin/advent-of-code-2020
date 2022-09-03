package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadIntegerLines() []int {

	f, err := os.Open("input.txt")
	arr := make([]int, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		converted, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error while converting.")
		}
		arr = append(arr, converted)
	}
	return arr
}
func day02() int {
	arr := ReadIntegerLines()
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					return arr[i] * arr[j] * arr[k]
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(day02())
}
