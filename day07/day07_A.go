package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Pair struct {
	key    string
	values []string
}

// first two functions are I/O operations
func ReadLines() map[string][]string {

	pairs := make(map[string][]string)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pair := inspectLine(scanner.Text())
		pairs[pair.key] = pair.values
	}
	return pairs
}
func inspectLine(str string) Pair {
	var result Pair
	r := regexp.MustCompile("^(.*?) bags")
	r2 := regexp.MustCompile("^(.*?)bag")
	key := r.FindString(str)
	result.key = key
	str = str[len(key)+11:]
	value := key
	for {
		value = r2.FindString(str)
		increment := 3
		if str[len(value)] == 's' {
			increment = 4
		}
		value += "s"
		result.values = append(result.values, value)
		if len(str) < len(value)+increment {
			break
		}
		str = str[len(value)+increment:]
	}
	return result
}
func reachShiny(pairs map[string][]string, current string) bool {
	if current == "shiny gold bags" {
		return true
	}
	if val, ok := pairs[current]; ok {
		for _, v := range val {
			if reachShiny(pairs, v) == true {
				return true
			}
		}
	}
	return false
}
func main() {
	data := ReadLines()
	count := -1 //minus 1 because of the key "shiny gold bags" 
	for k := range data {
		if reachShiny(data,k) {
			count++
		}
	}
	fmt.Println(count)
}
