package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Point struct {
  id        int64
  timestamp int64
}

func ReadLines() []Point {

  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make([]Point, 0)
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  scanner.Scan()
  count := 0
  text := strings.Split(scanner.Text(), ",")
  for _, v := range text {
    timestamp := count
    if v != "x" {
      conv, _ := strconv.Atoi(v)
      if timestamp > conv {
        timestamp = timestamp % conv
      }
      arr = append(arr, Point{int64(conv), int64(timestamp)})
    }
    count++
  }
  return arr
}
func findNum(p1 Point, p2 Point, start int64, increment int64) (i int64) {
  i = start
  for {
    if i%p1.id == p1.id-p1.timestamp || (p1.timestamp == 0 && i%p1.id == 0) {
      if i%p2.id == p2.id-p2.timestamp {
        return
      }
    }
    i += increment
  }
}
func main() {
  data := ReadLines()
  var start int64 = 0
  var increment int64 = 1
  for i := 0; i < len(data)-1; i++ {
    for j := i + 1; j < len(data); j++ {
      start = findNum(data[i], data[j], start, increment)
      if increment%data[i].id != 0 {
        increment = increment * data[i].id
      } else if increment%data[j].id != 0 {
        increment = increment * data[j].id
      }
    }
  }
  fmt.Println(start)
}
