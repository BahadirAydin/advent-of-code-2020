package main

import (
  "bufio"
  "fmt"
  "os"
)

// '.' means NOT SEAT, 'L' means empty, '#' means full
func ReadLines() [][]int32 {

  f, err := os.Open("input.txt")
  arr := make([][]int32, 0)
  if err != nil {
    fmt.Println("Error while reading.")
  }
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := make([]int32, 0)
    text := scanner.Text()
    for _, v := range text {
      line = append(line, v)
    }
    arr = append(arr, line)
  }
  return arr
}

func countNeighbors(data [][]int32, row int, column int) (neighbors int) {
  directions := [8]struct {
    x int
    y int
  }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
  rowLength := len(data)
  columnLength := len(data[0])
  for _, v := range directions {
    for i := 1; i < 50; i++ {
      x := v.x*i + row
      y := v.y*i + column
      if x < rowLength && y < columnLength && x >= 0 && y >= 0 && (data[x][y] == '#' || data[x][y] == 'e') {
        neighbors++
        break
      } else if x < rowLength && y < columnLength && x >= 0 && y >= 0 && (data[x][y] == 'L' || data[x][y] == 'f') {
        break
      }
    }
  }
  return
}
func changeState(data [][]int32) (flag bool) {
  for i := range data {
    for j := range data[i] {
      if data[i][j] == '.' {
        continue
      }
      n := countNeighbors(data, i, j)
      if n == 0 && data[i][j] == 'L' {
        flag = true
        data[i][j] = 'f' //SEATS THAT WILL GO EMPTY TO FULL AT THE END OF THE TURN
      } else if n >= 5 && data[i][j] == '#' {
        flag = true
        data[i][j] = 'e' // SEATS THAT WILL GO FULL TO EMPTY AT THE END OF THE TURN
      }
    }
  }
  for i := range data {
    for j := range data[i] {
      k := data[i][j]
      if k == 'f' {
        data[i][j] = '#'
      } else if k == 'e' {
        data[i][j] = 'L'
      }
    }
  }
  return
}
func countOccupied(data [][]int32) (sum int) {
  for i := range data {
    for j := range data[i] {
      k := data[i][j]
      if k == '#' {
        sum++
      }
    }
  }
  return
}

func main() {
  data := ReadLines()
  for changeState(data) {
  }
  result := countOccupied(data)
  fmt.Println(result)
}
