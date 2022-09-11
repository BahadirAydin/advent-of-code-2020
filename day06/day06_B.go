package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines() [][]string {

	f, err := os.Open("input.txt")
	arr := make([][]string, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	group := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			arr = append(arr, group)
			group = nil
			continue
		}
		group = append(group, text)
	}
	arr = append(arr, group) // last group
	return arr
}
func inspectGroup(group []string) (count int) {
	m := make(map[int32]int)
	for _, v := range group {
		for _, s := range v {
			m[s] += 1
		}
	}
	for _, v := range m {
		if v == len(group) {
			count++
		}
	}
	return
}

func main() {
	data := ReadLines()
	sum := 0
	for _, v := range data {
		sum += inspectGroup(v)
	}
	fmt.Println(sum)
}
