package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
)

func ReadLines() []string {

  f, err := os.Open("input.txt")
  arr := make([]string, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  var str string
  for scanner.Scan() {
    if len(scanner.Text()) == 0 {
      arr = append(arr, str)
      str = ""
    }
    str += scanner.Text()

  }
  arr = append(arr, str) // to add the last one
  return arr
}
func matchStr(check string, data *string) bool {
  if match, _ := regexp.MatchString(check, *data); match {
    return true
  }
  return false
}
func matchAll(check *[7]string, data *string) bool {
  for _, v := range *check {
    if matchStr(v, data) == false {
      return false
    }
  }
  return true
}
func main() {
  data := ReadLines()
  expected := [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
  count := 0
  for _, v := range data {
    if matchAll(&expected, &v) {
      count++
    }
  }
  fmt.Println(count)
}
