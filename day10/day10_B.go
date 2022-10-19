package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
)

func ReadIntegerLines() []int {

  f, err := os.Open("input.txt")
  arr := make([]int, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    converted, err := strconv.Atoi(scanner.Text())
    if err != nil {
      fmt.Println("Error while converting.")
    }
    arr = append(arr, converted)
  }
  return arr
}
func countArrangements(data []int, index int) (sum uint64) {
  if index+3 < len(data) && data[index+3] == data[index]+3 {
    sum += countArrangements(data, index+1)
    sum += countArrangements(data, index+2)
    sum += countArrangements(data, index+3)
  } else if index+2 < len(data) && data[index+2] == data[index]+2 {
    sum += countArrangements(data, index+1)
    sum += countArrangements(data, index+2)
  } else if index+2 < len(data) && data[index+2] == data[index]+3 {
    sum += countArrangements(data, index+1)
    sum += countArrangements(data, index+2)
  } else if index < len(data) {
    sum += countArrangements(data, index+1)
  } else {
    return 1
  }
  return sum
}
func subGraphs(data []int) (sum uint64) {
  sum = 1
  start := 0
  for k := 0; k < len(data); k++ {
    if k < len(data)-1 && data[k]+3 == data[k+1] {
      sum *= countArrangements(data[start:k+1], 0)
      start = k + 1
      k += 1
    } else if k == len(data)-1 {
      sum *= countArrangements(data[start:k+1], 0)
    }
  }
  return sum
}
func main() {
  data := ReadIntegerLines()
  data = append(data, 0)
  sort.Ints(data)
  sum := subGraphs(data)
  fmt.Println(sum)
}
