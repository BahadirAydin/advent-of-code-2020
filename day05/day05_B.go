package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() []string {

	f, err := os.Open("input.txt")
	arr := make([]string, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}
func findSeat(code string) (row int, column int) {
	lower := 0
	upper := 127
	for i := 0; i < 6; i++ {
		if code[i] == 'F' {
			upper = (lower + upper) / 2
		} else if code[i] == 'B' {
			lower = (lower+upper)/2 + 1
		}
	}
	if code[6] == 'F' {
		row = lower
	} else {
		row = upper
	}
	lower = 0
	upper = 7
	for i := 7; i < 9; i++ {
		if code[i] == 'L' {
			upper = (lower + upper) / 2
		} else if code[i] == 'R' {
			lower = (lower+upper)/2 + 1
		}
	}
	if code[9] == 'L' {
		column = lower
	} else {
		column = upper
	}
	return
}
func main() {
	var arr [128][8]bool
	var ids [1032]bool
	data := ReadLines()
	for _, v := range data {
		r, c := findSeat(v)
		if r == 127 || r == 0 {
			continue
		}
		arr[r][c] = true
		ids[8*r+c] = true
	}
	for i := 1; i < 126; i++ {
		for j := 0; j < 8; j++ {
			if arr[i][j] == false {
				id := 8*i+j
				if ids[id+1] && ids[id-1]{
					fmt.Println(id)
				}
			}
		}
	}
}
