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
func main() {
  data := ReadIntegerLines()
  sort.Ints(data)
  oneDiff, threeDiff := 0, 1 //threeDiff starts with 1 to account for our charger at the end that is not in the list. (always a 3 diff)
  prev := 0
  for _, v := range data {
    diff := v - prev
    if diff == 1 {
      oneDiff++
    } else if diff == 3 {
      threeDiff++
    }
    prev = v
  }
  fmt.Println(oneDiff * threeDiff)
}
