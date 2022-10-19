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

func possibleSums(data []int, item int) bool {
  for i := 0; i < len(data)-1; i++ {
    for j := i + 1; j < len(data); j++ {
      if item == data[i]+data[j] {
        return true
      }
    }
  }
  return false
}
func main() {
  data := ReadIntegerLines()
  for i := 25; i < len(data); i++ {
    if possibleSums(data[i-25:i], data[i]) == false {
      fmt.Println(data[i])
    }
  }
}
