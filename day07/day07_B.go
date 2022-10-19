package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

type Pair struct {
  key    string
  values []nameAndCount
}
type nameAndCount struct {
  name  string
  count int
}

// first two functions are I/O operations
func ReadLines() map[string][]nameAndCount {

  pairs := make(map[string][]nameAndCount)
  f, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    pair := inspectLine(scanner.Text())
    pairs[pair.key] = pair.values
  }
  return pairs
}
func inspectLine(str string) Pair {
  var result Pair
  r := regexp.MustCompile("^(.*?) bags")
  r2 := regexp.MustCompile("^(.*?)bag")
  key := r.FindString(str)
  result.key = key
  str = str[len(key)+9:]
  value := key
  for {
    num, _ := strconv.Atoi(string(str[0]))
    str = str[2:]
    value = r2.FindString(str)
    increment := 1
    if str[len(value)] == 's' {
      increment = 2
    }
    value += "s"

    result.values = append(result.values, nameAndCount{value, num})
    if len(str) < len(value)+increment {
      break
    }
    str = str[len(value)+increment:]
  }
  return result
}
func calculateShiny(pairs map[string][]nameAndCount, current string) (sum int) {
  sum = 1
  shiny := pairs[current]
  for _, v := range shiny {
    sum += calculateShiny(pairs, v.name) * v.count
  }
  fmt.Println(current, " over:", sum)
  return
}
func main() {
  data := ReadLines()
  fmt.Println(data)
  result := calculateShiny(data, "shiny gold bags") - 1 //bag doesn't contain itself
  fmt.Println(result)
}
