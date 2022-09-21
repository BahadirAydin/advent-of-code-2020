package main

//IMPORTANT: You should add "mask = end" to the end of your input as a newline to run this code

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	mask    string
	indexes []int64
	values  []int
}

func ReadLines() []Operation {

	f, _ := os.Open("input.txt")
	defer f.Close()
	arr := make([]Operation, 0)
	scanner := bufio.NewScanner(f)
	var tmp Operation
	indexes := make([]int64, 0)
	values := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text[1] == 'a' {
			cpyIndex := make([]int64, len(indexes))
			copy(cpyIndex, indexes)
			cpyValues := make([]int, len(values))
			copy(cpyValues, values)
			tmp.indexes = cpyIndex
			tmp.values = cpyValues
			arr = append(arr, tmp)
			tmp.mask = text[7:]
			indexes = nil
			values = nil
			continue
		}
		rightInd := strings.Index(text, "]")
		a, _ := strconv.Atoi(text[4:rightInd])
		value, _ := strconv.Atoi(text[rightInd+4:])
		index := int64(a)
		indexes = append(indexes, index)
		values = append(values, value)
	}
	return arr[1:]
}
func reverseStr(str string) (reversed string) {
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return
}
func findResultStr(mask string, v int64) string {

	result := ""
	s := strconv.FormatInt(v, 2)
	for i := 0; i < len(mask); i++ {
		if mask[35-i] == '0' {
			if i < len(s) {
				result += string(s[len(s)-1-i])
			} else {
				result += "0"
			}
		} else if mask[35-i] == '1' {
			result += "1"
		} else {
			result += "X"
		}
	}
	result = reverseStr(result)
	return result
}
func findPossibilites(result string, index int, exp float64, num int, pos *[]int) {
	for i := index; i >= 0; i-- {
		if result[i] == '1' {
			num += int(math.Pow(2, exp))
		} else if result[i] == 'X' {
			newNum := num + int(math.Pow(2, exp))
			findPossibilites(result, i-1, exp+1, newNum, pos)
			findPossibilites(result, i-1, exp+1, num, pos)
			return
		}
		exp++
	}
	*pos = append(*pos, num)
}
func main() {
	data := ReadLines()
	space := make(map[int]int)
	for _, v := range data {
		mask := v.mask
		for i := 0; i < len(v.indexes); i++ {
			r := findResultStr(mask, v.indexes[i])
			pos := make([]int, 0)
			findPossibilites(r, len(r)-1, 0, 0, &pos)
			for _, p := range pos {
				space[p] = v.values[i]
			}
		}
	}
	sum := 0
	for _, v := range space {
		sum += v
	}
	fmt.Println(sum)
}
