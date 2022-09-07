package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadLines() []string {

	f, err := os.Open("input.txt")
	arr := make([]string, 0)
	if err != nil {
		fmt.Println("Error while reading.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var str string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			arr = append(arr, str)
			str = ""
		}
		str += " " + scanner.Text() + " "

	}
	arr = append(arr, str) // to add the last one
	return arr
}
func matchByr(data string) bool {
	r, _ := regexp.Compile("byr:([0-9]+) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	found = found[len("byr:") : len(found)-1]
	num, err := strconv.Atoi(found)
	if err != nil {
		return false
	}
	if num <= 2002 && num >= 1920 {
		return true
	}
	return false
}
func matchIyr(data string) bool {
	r, _ := regexp.Compile("iyr:([0-9]+) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	found = found[len("iyr:") : len(found)-1]
	num, err := strconv.Atoi(found)
	if err != nil {
		return false
	}
	if num <= 2020 && num >= 2010 {
		return true
	}
	return false
}
func matchEyr(data string) bool {
	r, _ := regexp.Compile("eyr:([0-9]+) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	found = found[len("eyr:") : len(found)-1]
	num, err := strconv.Atoi(found)
	if err != nil {
		return false
	}
	if num <= 2030 && num >= 2020 {
		return true
	}
	return false
}
func matchHgt(data string) bool {
	cm := regexp.MustCompile("hgt:([0-9]+)cm")
	in := regexp.MustCompile("hgt:([0-9]+)in")
	found_cm := cm.FindString(data)
	found_in := in.FindString(data)
	if found_cm == "" && found_in == "" {
		return false
	} else if found_cm == "" {
		found_in = found_in[len("hgt:") : len(found_in)-2]
		num, _ := strconv.Atoi(found_in)
		if num >= 59 && num <= 76 {
			return true
		}
	} else {
		found_cm = found_cm[len("hgt:") : len(found_cm)-2]
		num, _ := strconv.Atoi(found_cm)
		if num >= 150 && num <= 193 {
			return true
		}
	}
	return false
}
func matchHcl(data string) bool {
	r := regexp.MustCompile("hcl:#([0-9a-f]{6}) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	return true
}
func matchEcl(data string) bool {
	r := regexp.MustCompile("ecl:([a-z]{3}) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	found = found[len("ecl:") : len(found)-1]
	// amb blu brn gry grn hzl oth
	if found == "amb" || found == "blu" || found == "brn" || found == "gry" || found == "grn" || found == "hzl" || found == "oth" {
		return true
	}
	return false
}
func matchPid(data string) bool {
	r := regexp.MustCompile("pid:([0-9]{9}) ")
	found := r.FindString(data)
	if found == "" {
		return false
	}
	return true
}
func matchAll(data string) bool {
	if matchByr(data) && matchIyr(data) && matchEyr(data) && matchHgt(data) && matchHcl(data) && matchEcl(data) && matchPid(data) {
		return true
	}
	return false
}
func main() {
	data := ReadLines()
	count := 0
	for _, v := range data {
		if matchAll(v) {
			count++
		}
	}
	fmt.Println(count)
}
