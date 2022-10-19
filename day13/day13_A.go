package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func ReadLines() (int, []int) {

  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make([]int, 0)
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  time, _ := strconv.Atoi(scanner.Text())
  scanner.Scan()
  text := strings.Split(scanner.Text(), ",")
  for _, v := range text {
    if v != "x" {
      conv, _ := strconv.Atoi(v)
      arr = append(arr, conv)
    }
  }
  return time, arr
}
func smallestBusModulo(time int, data []int) (int, int) {
  smallest := 10000000
  smallestId := -1
  for _, v := range data {
    if v-time%v < smallest {
      smallest = v - time%v
      smallestId = v
    }
  }
  return smallest, smallestId
}
func main() {
  time, data := ReadLines()
  waitTime, id := smallestBusModulo(time, data)
  fmt.Println(waitTime * id)
}
