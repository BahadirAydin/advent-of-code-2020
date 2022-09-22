package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Compare struct {
	min int
	max int
}
type Criteria struct {
	c1 Compare
	c2 Compare
}

func ReadCriteria() ([]Criteria, string, []string) {

	f, _ := os.Open("input.txt")
	defer f.Close()
	arr := make([]Criteria, 0)
	nearbyTickets := make([]string, 0)
	var myTicket string
	scanner := bufio.NewScanner(f)
	r1 := regexp.MustCompile(": ([^;]*) or")
	r2 := regexp.MustCompile("or ([^;]*)")
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			scanner.Scan()
			scanner.Scan()
			myTicket = scanner.Text()
			scanner.Scan()
			scanner.Scan()
			for scanner.Scan() {
				nearbyTickets = append(nearbyTickets, scanner.Text())
			}
			break
		}
		match1 := r1.FindString(text)
		match2 := r2.FindString(text)
		match1 = match1[2 : len(match1)-3]
		match2 = match2[3:]
		c1 := strings.Split(match1, "-")
		c2 := strings.Split(match2, "-")
		conv1, _ := strconv.Atoi(c1[0])
		conv2, _ := strconv.Atoi(c1[1])
		comp1 := Compare{
			conv1,
			conv2,
		}
		conv1, _ = strconv.Atoi(c2[0])
		conv2, _ = strconv.Atoi(c2[1])
		comp2 := Compare{
			conv1,
			conv2,
		}
		crit := Criteria{
			comp1,
			comp2,
		}
		arr = append(arr, crit)
	}
	return arr, myTicket, nearbyTickets
}
func modifyTicket(ticket string) []int {
	subArr := make([]int, 0)

	splitTickets := strings.Split(ticket, ",")
	for _, t := range splitTickets {
		conv, _ := strconv.Atoi(t)
		subArr = append(subArr, conv)
	}
	return subArr
}
func modifyTickets(tickets []string) [][]int {
	arr := make([][]int, 0)
	for _, v := range tickets {
		subArr := modifyTicket(v)
		arr = append(arr, subArr)
	}
	return arr
}
func ifValid(criteria []Criteria, ticket []int) bool {
	for _, v := range ticket {
		flag := true
		for _, c := range criteria {
			c1 := c.c1
			c2 := c.c2
			if (v <= c1.max && v >= c1.min) || (v <= c2.max && v >= c2.min) {
				flag = false
				break
			}
		}
		if flag {
			return false
		}
	}
	return true
}
func findCrits(crit Criteria, tickets [][]int, used []int) []int {
	c1 := crit.c1
	c2 := crit.c2
	arr := make([]int, 0)
	columnSize := len(tickets[0])
	for column := 0; column < columnSize; column++ {
		flag := true
		for _, v := range tickets {
			curr := v[column]
			if (curr > c1.max || curr < c1.min) && (curr > c2.max || curr < c2.min) {
				flag = false
				break
			}
		}
		contains := false
		for _, c := range used {
			if c == column {
				contains = true
				break
			}
		}
		if flag && !contains {
			arr = append(arr, column)
		}
	}
	return arr
}
func findValidTickets(criteria []Criteria, tickets [][]int) [][]int {
	validTickets := make([][]int, 0)
	for _, v := range tickets {
		if ifValid(criteria, v) {
			validTickets = append(validTickets, v)
		}
	}
	return validTickets
}

func main() {
	criteria, myTicket, nearbyTickets := ReadCriteria()
	tickets := modifyTickets(nearbyTickets)
	myTicketInt := modifyTicket(myTicket)
	validTickets := findValidTickets(criteria, tickets)
	columnsMap := make(map[int]int)
	used := make([]int, 0)
	size := len(criteria)
	for len(used) < size {
		for i := 0; i < size; i++ {
			crits := findCrits(criteria[i], validTickets, used)
			if len(crits) == 1 {
				columnsMap[i] = crits[0]
				used = append(used, crits[0])
			}
		}
	}
	result := 1
	for i := 0; i < 6; i++ {
		result *= myTicketInt[columnsMap[i]]
	}
	fmt.Println(result)
}
