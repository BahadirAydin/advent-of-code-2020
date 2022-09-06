package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() [][]bool {

	f, err := os.Open("input.txt")
	arr := make([][]bool, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := make([]bool, 0)

		str := scanner.Text()
		for _, v := range str {
			// 46 is the byte value for "." and 35 is the byte value for "#"
			if v == 46 {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}
		arr = append(arr, row)
	}
	return arr
}
func traverse(step_right int, step_down int, arr [][]bool) int {
	count := 0
	row_len := len(arr[0])
	column_len := len(arr)
	curr_row := 0
	curr_column := 0
	for curr_column < column_len {
		if arr[curr_column][curr_row] {
			count++
		}
		curr_row += step_right
		curr_column += step_down
		if curr_row >= row_len {
			curr_row -= row_len
		}
	}
	return count
}
func main() {
	data := ReadLines()
	a := traverse(1, 1, data)
	b := traverse(3, 1, data)
	c := traverse(5, 1, data)
	d := traverse(7, 1, data)
	e := traverse(1, 2, data)

	fmt.Println(a * b * c * d * e)
}
