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

func ReadLines() ([]Criteria, string, []string) {

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
func modifyTickets(tickets []string) [][]int {
  arr := make([][]int, 0)
  for _, v := range tickets {
    subArr := make([]int, 0)
    splitTickets := strings.Split(v, ",")
    for _, t := range splitTickets {
      conv, _ := strconv.Atoi(t)
      subArr = append(subArr, conv)
    }
    arr = append(arr, subArr)
  }
  return arr
}
func findInvalid(criteria []Criteria, ticket []int) int {
  sum := 0
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
      sum += v
    }
  }
  return sum
}

func main() {
  criteria, _, nearbyTickets := ReadLines()
  tickets := modifyTickets(nearbyTickets)
  sum := 0
  for _, v := range tickets {
    sum += findInvalid(criteria, v)
  }
  fmt.Println(sum)

}
