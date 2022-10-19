package main

import (
  "bufio"
  "fmt"
  "os"
)

type Point struct {
  x          int
  y          int
  z          int
  active     bool
  activate   bool
  deactivate bool
}

func ReadLines() []Point {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  arr := make([]Point, 0)
  y_pos := 0
  for scanner.Scan() {
    text := scanner.Text()
    x_pos := 0
    for _, v := range text {
      if v == '#' {
        arr = append(arr, Point{x_pos, y_pos, 0, true, false, false})
      } else {
        arr = append(arr, Point{x_pos, y_pos, 0, false, false, false})
      }
      x_pos++
    }
    y_pos++
  }
  return arr
}

func produceNeighbors(p Point) [26]Point {
  var neighbors [26]Point
  // arr is copy-pasted from https://stackoverflow.com/questions/71134868/get-26-nearest-neighbors-of-a-point-in-3d-space-vectorized
  arr := [26][3]int{
    {-1, -1, 0},
    {-1, -1, 1},
    {-1, 0, -1},
    {-1, 0, 0},
    {-1, 0, 1},
    {-1, 1, -1},
    {-1, 1, 0},
    {-1, 1, 1},
    {0, -1, -1},
    {0, -1, 0},
    {0, -1, 1},
    {0, 0, -1},
    {-1, -1, -1},
    {0, 0, 1},
    {0, 1, -1},
    {0, 1, 0},
    {0, 1, 1},
    {1, -1, -1},
    {1, -1, 0},
    {1, -1, 1},
    {1, 0, -1},
    {1, 0, 0},
    {1, 0, 1},
    {1, 1, -1},
    {1, 1, 0},
    {1, 1, 1},
  }
  for k, v := range arr {
    neighbors[k].x = v[0] + p.x
    neighbors[k].y = v[1] + p.y
    neighbors[k].z = v[2] + p.z
  }

  return neighbors
}
func contains(data *[]Point, check Point) (bool, int) {
  for k, p := range *data {
    if p.x == check.x && p.y == check.y && p.z == check.z {
      return true, k
    }
  }
  return false, 0
}
func increaseBoundary(data *[]Point) {
  for _, v := range *data {
    neighbors := produceNeighbors(v)
    for _, n := range neighbors {
      if ok, _ := contains(data, n); !ok {
        *data = append(*data, n)
      }
    }
  }
}
func turn(data *[]Point) {
  increaseBoundary(data)
  for index, v := range *data {
    neighbors := produceNeighbors(v)
    for k, curr := range neighbors {
      if ok, index := contains(data, curr); ok {
        neighbors[k].active = (*data)[index].active
      }
    }
    n := activeNeighbors(neighbors)
    if v.active && !(n == 2 || n == 3) {
      (*data)[index].deactivate = true
    } else if !v.active && n == 3 {
      (*data)[index].activate = true
    }
  }
  for k, v := range *data {
    if v.activate {
      (*data)[k].active = true
      (*data)[k].activate = false
    } else if v.deactivate {
      (*data)[k].active = false
      (*data)[k].deactivate = false
    }
  }
}
func activeNeighbors(n [26]Point) (c int) {
  for _, v := range n {
    if v.active {
      c++
    }
  }
  return
}
func activeCount(data []Point) (c int) {
  for _, v := range data {
    if v.active {
      c++
    }
  }
  return
}
func main() {
  data := ReadLines()
  for i := 0; i < 6; i++ {
    turn(&data)
  }
  fmt.Println(activeCount(data))
}
