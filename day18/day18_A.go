package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadLines() []string {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}
func findIndex(str string) int {
	count := 0
	for k, v := range str {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
		if count == 0 {
			return k
		}
	}
	return -1
}
func evaluateExpression(op1 int, op2 int, operator byte) int {
	if operator == '*' {
		return op1 * op2
	} else if operator == '+' {
		return op1 + op2
	}
	return 0
}
func evaluateString(str string) int {
	var op1, op2 int
	var f1, f2 bool
	var operator byte
	numStr := ""
	str += " "
	for k := 0; k < len(str); k++ {
		v := str[k]
		if v == ' ' {
			if numStr != "" {
				conv, _ := strconv.Atoi(numStr)
				if !f1 {
					op1 = conv
					f1 = true
					numStr = ""
				} else {
					op2 = conv
					f2 = true
					numStr = ""
				}
			}
		} else if v == '*' || v == '+' {
			operator = v
		} else if v == '(' {

			index := findIndex(str[k:]) + k
			if !f1 {
				op1 = evaluateString(str[k+1 : index])
				f1 = true
				k += index - k
			} else {
				op2 = evaluateString(str[k+1 : index])
				f2 = true
				k += index - k
			}
		} else {
			numStr += string(v)
		}
		if f1 && f2 {
			op1 = evaluateExpression(op1, op2, operator)
			f2 = false
		}
	}
	return op1
}

func main() {
	data := ReadLines()
	sum := 0
	for _, v := range data {
		r := evaluateString(v)
		sum += r
	}
	fmt.Println(sum)

}
