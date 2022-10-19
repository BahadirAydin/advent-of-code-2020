package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
)

type Item struct {
  foods     []string
  allergens []string
}

func ReadLines() []Item {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  arr := make([]Item, 0)
  for scanner.Scan() {
    arr = append(arr, Item{})
    text := scanner.Text()
    reg := regexp.MustCompile(`(?m)\bcontains\s+(.*)$`)
    allergens := reg.FindString(text)
    allergens = allergens[9:len(allergens)-1] + " "
    str := ""
    for i := 0; i < len(allergens); i++ {
      if allergens[i] == ',' || allergens[i] == ' ' {
        arr[len(arr)-1].allergens = append(arr[len(arr)-1].allergens, str)
        str = ""
        i++
        continue
      }
      str += string(allergens[i])
    }
    str = ""
    fmt.Println(allergens)
    for _, v := range text {
      if v == '(' {
        break
      }
      if v == ' ' {
        arr[len(arr)-1].foods = append(arr[len(arr)-1].foods, str)
        str = ""
        continue
      }
      str += string(v)
    }
  }
  return arr
}
func main() {
  data := ReadLines()
  fmt.Println(data)
}
