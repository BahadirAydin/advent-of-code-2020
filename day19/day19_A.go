package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Rules [][]int

var maxLen int

func ReadLines() (map[int]Rules, map[int]string, []string) {
  f, _ := os.Open("input_A.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  connectedRules := make(map[int]Rules)
  singletonRules := make(map[int]string)
  strArray := make([]string, 0)
  mode := false

  for scanner.Scan() {
    str := scanner.Text()
    if mode {
      strArray = append(strArray, str)
      continue
    }
    if len(str) == 0 {
      mode = true
      continue
    }
    index := strings.Index(str, ":")
    ruleNumber, _ := strconv.Atoi(str[0:index])
    str = str[index+1:]
    numStr := ""
    rules := make([][]int, 0)
    currRule := 0
    str += " "
    singleton := strings.Index(str, "\"")
    if singleton != -1 {
      singletonRules[ruleNumber] = str[singleton+1 : singleton+2]
    } else {
      for i := 0; i < len(str); i++ {
        v := str[i]
        if v == ' ' {
          if numStr != "" {
            conv, _ := strconv.Atoi(numStr)
            numStr = ""
            if len(rules) == currRule {
              rules = append(rules, make([]int, 0))
            }
            rules[currRule] = append(rules[currRule], conv)
          }
        } else if v == '|' {
          currRule++
        } else {
          numStr += string(v)
        }
      }
      connectedRules[ruleNumber] = rules
    }
  }
  return connectedRules, singletonRules, strArray
}
func replace(n int, newRule []int, rules *[]int) {
  index := -1
  for k := range *rules {
    if (*rules)[k] == n {
      index = k
    }
  }
  if index == -1 {
    return
  }
  cpy := make([]int, len((*rules)[index+1:]))
  copy(cpy, (*rules)[index+1:])
  (*rules) = append(append((*rules)[:index], newRule...), cpy...)
}
func replaceAll(n int, newRules Rules, oldRules *Rules, change []int) [][]int {
  l := len(newRules) - 1
  for i := 0; i < l; i++ {
    cpy := make([]int, len(change))
    copy(cpy, change)
    *oldRules = append((*oldRules), cpy)
  }
  cpy := make([]int, len(change))
  copy(cpy, change)
  for _, v := range newRules {
    index := findRule(cpy, *oldRules)
    replace(n, v, &(*oldRules)[index])
  }
  return *oldRules
}
func findRule(rule []int, rules Rules) int {
  for k, v := range rules {
    if isEqual(rule, v) {
      return k
    }
  }
  return -1
}
func isEqual(s1 []int, s2 []int) bool {
  if len(s1) != len(s2) {
    return false
  }
  for i := 0; i < len(s1); i++ {
    if s1[i] != s2[i] {
      return false
    }
  }
  return true
}
func getNewRules(n int, connectedRules map[int]Rules) Rules {
  return connectedRules[n]
}
func getChangeIndex(rules Rules, singleton map[int]string, index int) (bool, int, []int, int) {
  for k := index; k < len(rules); k++ {
    rule := rules[k]
    if len(rule) >= maxLen {
      fmt.Println("yas")
      return true, 0, []int{}, 0
    }
    for _, v := range rule {
      if _, ok := singleton[v]; !ok {
        return false, v, rule, k
      }
    }
  }
  return true, 0, []int{}, 0
}
func turnToStr(rules Rules, singleton map[int]string) []string {
  arr := make([]string, 0)
  for _, rule := range rules {
    str := ""
    for _, ch := range rule {
      str += singleton[ch]
    }
    arr = append(arr, str)
  }
  return arr
}
func mutate(connectedRules map[int]Rules, singleton map[int]string, ruleNum int) [][]int {
  rules := connectedRules[ruleNum]
  over, n, change, index := getChangeIndex(rules, singleton, 0)
  for !over {
    connectedRules[ruleNum] = replaceAll(n, getNewRules(n, connectedRules), &rules, change)
    over, n, change, index = getChangeIndex(rules, singleton, index)
    rules = connectedRules[ruleNum]
  }
  return rules
}
func findAllStr(allRules map[int]Rules, singleton map[int]string) map[int][]string {
  arr := make(map[int][]string)
  for i := 1; i < len(allRules); i++ {
    arr[i] = turnToStr(allRules[i], singleton)
  }
  return arr
}
func divideHelper(arr []string, check string) []string {
  remaining := make([]string, 0)
  for _, v := range arr {
    for i := len(check); i >= 0; i-- {
      if check[:i] == v {
        remaining = append(remaining, check[i:])
      }
    }
  }
  return remaining
}
func divide(strMap map[int][]string, check string, rule []int, index int) bool {
  if check == "" {
    return true
  } else if len(check) > 0 && len(rule) == index {
    return false
  }
  curr := strMap[rule[index]]
  remaining := divideHelper(curr, check)
  for _, v := range remaining {
    ok := divide(strMap, v, rule, index+1)
    if ok {
      return true
    }
  }
  return false
}
func main() {
  connectedRules, singletonRules, check := ReadLines()
  for _, s := range check {
    if len(s) > maxLen {
      maxLen = len(s)
    }
  }
  for i := 1; i < len(connectedRules); i++ {
    fmt.Println(i)
    connectedRules[i] = mutate(connectedRules, singletonRules, i)
  }
  strMap := findAllStr(connectedRules, singletonRules)
  for k, v := range singletonRules {
    strMap[k] = []string{v}
  }
  rule := connectedRules[0][0]
  count := 0
  for _, v := range check {
    ok := divide(strMap, v, rule, 0)
    if ok {
      count++
    }
  }
  fmt.Println(count)
}
