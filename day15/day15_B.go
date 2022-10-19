package main

import "fmt"

func turn(lastSpoken int, turnNum int, track *map[int]int) (spoken int) {
  if v, ok := (*track)[lastSpoken]; ok {
    spoken = turnNum - v - 1
  } else {
    spoken = 0
  }
  (*track)[lastSpoken] = turnNum - 1
  return
}

func main() {
  var arr []int = []int{16, 12, 1, 0, 15, 7, 11}
  //var arr []int = []int{0, 3, 6}
  track := make(map[int]int)
  for k, v := range arr {
    track[v] = k + 1
  }
  spoken := arr[len(arr)-1]
  for i := len(arr) + 1; i <= 30000000; i++ {
    spoken = turn(spoken, i, &track)
  }
  fmt.Println(spoken)
}
