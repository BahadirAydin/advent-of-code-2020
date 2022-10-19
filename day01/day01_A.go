package main

import (
  "bufio"
  "fmt"
  "os"
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
func day01() int {
  arr := ReadIntegerLines()
  for i := 0; i < len(arr); i++ {
    for j := i + 1; j < len(arr); j++ {
      if arr[i]+arr[j] == 2020 {
        return arr[i] * arr[j]
      }
    }
  }
  return 0
}

func main() {
  fmt.Println(day01())
}
