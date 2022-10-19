package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Data struct {
  min  int
  max  int
  char byte
  str  string
}

func ReadLines() []Data {

  f, err := os.Open("input.txt")
  arr := make([]Data, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    str := scanner.Text()
    res := strings.Split(str, " ")
    nums := strings.Split(res[0], "-")
    min, _ := strconv.Atoi(nums[0])
    max, _ := strconv.Atoi(nums[1])
    char := res[1][0]
    arr = append(arr, Data{min, max, char, res[2]})
  }
  return arr
}
func checkString(min int, max int, check byte, str string) bool {
  count := 0
  for i := 0; i < len(str); i++ {
    if str[i] == check {
      count++
    }
  }
  if count >= min && count <= max {
    return true
  }
  return false
}
func main() {
  data := ReadLines()
  valid := 0
  for _, v := range data {
    status := checkString(v.min, v.max, v.char, v.str)
    if status {
      valid++
    }
  }
  fmt.Println(valid)
}
