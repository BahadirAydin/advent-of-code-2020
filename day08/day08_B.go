package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Operation struct {
	code       string
	value      int
	isExecuted bool
}

func ReadLines() []Operation {

	f, err := os.Open("input.txt")
	arr := make([]Operation, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var p Operation
		text := scanner.Text()
		p.code = text[:3]
		p.value, _ = strconv.Atoi(text[4:])
		arr = append(arr, p)
	}
	return arr
}
func switchCode(data []Operation, index int ) (cpy []Operation,found bool){
	cpy = make([]Operation, len(data))
	copy(cpy, data)
	count := 0
	for i:=0;i<len(cpy);i++{
		if(index == count){
			found = true
			if cpy[i].code == "jmp" {
				cpy[i].code = "nop"
				return 
			} else if cpy[i].code == "nop" {
				cpy[i].code = "jmp"
				break
			}
		} else {
			if cpy[i].code == "jmp" || cpy[i].code == "nop" {
				count++
			}
		}
	}
	return
}

func runAccumulator(data []Operation) (accumulator int ,index int){
	index = 0
	accumulator = 0
	for {
		if index >= len(data) {
			return
		}
		current := data[index]
		if current.isExecuted {
			return
		}
		data[index].isExecuted = true
		if current.code == "acc" {
			accumulator += current.value
			index++
		} else if current.code == "jmp" {
			index += current.value
		} else if current.code == "nop" {
			index++
		}
	}
}

func main() {
	data := ReadLines()
	changeIndex := 0
	for {	
		if newData, ok := switchCode(data,changeIndex); ok {
			result, index := runAccumulator(newData)
			if index == len(data) {
				fmt.Println(result)
				return
			}
			changeIndex++
		}
	}
}
