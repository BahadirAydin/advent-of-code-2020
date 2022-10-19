package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
)

const STEP = 25

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
func sumBetween(data []int, startIndex int, item int) (int, bool) {
  sum := 0
  max := 0
  min := math.MaxInt32
  for i := startIndex; i < len(data); i++ {
    sum += data[i]
    if data[i] > max {
      max = data[i]
    }
    if data[i] < min {
      min = data[i]
    }
    if sum == item {
      return min + max, true
    } else if sum > item {
      return 0, false
    }
  }
  return 0, false
}

func main() {
  data := ReadIntegerLines()
  falseInt := 0
  for i := STEP; i < len(data); i++ {
    if possibleSums(data[i-STEP:i], data[i]) == false {
      falseInt = data[i]
      break
    }
  }
  for i := 0; i < len(data); i++ {

    if result, ok := sumBetween(data, i, falseInt); ok {
      fmt.Println(result)
      return
    }
  }
}
