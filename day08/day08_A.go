package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type Operation struct {
  code       string
  value      int
  isExecuted bool
}

func ReadLines() []Operation {

  f, err := os.Open("input.txt")
  arr := make([]Operation, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var p Operation
    text := scanner.Text()
    p.code = text[:3]
    p.value, _ = strconv.Atoi(text[4:])
    arr = append(arr, p)
  }
  return arr
}
func main() {
  accumulator := 0
  index := 0
  data := ReadLines()
  for {
    current := data[index]
    if current.isExecuted {
      break
    }
    data[index].isExecuted = true
    if current.code == "acc" {
      accumulator += current.value
      index++
    } else if current.code == "jmp" {
      index += current.value
    } else if current.code == "nop" {
      index++
    }
  }
  fmt.Println(accumulator)
}
