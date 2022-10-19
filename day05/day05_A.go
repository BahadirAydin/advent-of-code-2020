package main

import (
  "bufio"
  "fmt"
  "os"
)

func ReadLines() []string {

  f, err := os.Open("input.txt")
  arr := make([]string, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    arr = append(arr, scanner.Text())
  }
  return arr
}
func findSeat(code string) (row int, column int) {
  lower := 0
  upper := 127
  for i := 0; i < 6; i++ {
    if code[i] == 'F' {
      upper = (lower + upper) / 2
    } else if code[i] == 'B' {
      lower = (lower+upper)/2 + 1
    }
  }
  if code[6] == 'F' {
    row = lower
  } else {
    row = upper
  }
  lower = 0
  upper = 7
  for i := 7; i < 9; i++ {
    if code[i] == 'L' {
      upper = (lower + upper) / 2
    } else if code[i] == 'R' {
      lower = (lower+upper)/2 + 1
    }
  }
  if code[9] == 'L' {
    column = lower
  } else {
    column = upper
  }
  return
}
func main() {
  data := ReadLines()
  highestId := 0
  for _, v := range data {
    r, c := findSeat(v)
    id := r*8 + c
    if id > highestId {
      highestId = id
    }
  }
  fmt.Println(highestId)
}
