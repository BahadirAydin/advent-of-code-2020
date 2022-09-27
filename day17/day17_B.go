package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x      int
	y      int
	z      int
	w      int
	active bool
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
				arr = append(arr, Point{x_pos, y_pos, 0, 0, true})
			} else {
				arr = append(arr, Point{x_pos, y_pos, 0, 0, false})
			}
			x_pos++
		}
		y_pos++
	}
	return arr
}

func produceNeighbors(p Point) [80]Point {
	var neighbors [80]Point
	count := 0
	for dz := p.z - 1; dz <= p.z+1; dz++ {
		for dy := p.y - 1; dy <= p.y+1; dy++ {
			for dx := p.x - 1; dx <= p.x+1; dx++ {
				for dw := p.w - 1; dw <= p.w+1; dw++ {
					if (dx != p.x) || (dy != p.y) || (dz != p.z) || (dw != p.w) {
						neighbors[count] = Point{dx, dy, dz, dw, false}
						count++
					}
				}
			}
		}
	}
	return neighbors
}
func contains(data *[]Point, check Point) (bool, int) {
	for k, p := range *data {
		if p.x == check.x && p.y == check.y && p.z == check.z && p.w == check.w {
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
	deactivate := make([]int, 0)
	activate := make([]int, 0)
	for index, v := range *data {
		neighbors := produceNeighbors(v)
		for k, curr := range neighbors {
			if ok, index := contains(data, curr); ok {
				neighbors[k].active = (*data)[index].active
			}
		}
		n := activeNeighbors(neighbors)
		if v.active && !(n == 2 || n == 3) {
			deactivate = append(deactivate, index)
		} else if !v.active && n == 3 {
			activate = append(activate, index)
		}
	}
	for _, v := range deactivate {
		(*data)[v].active = false
	}
	for _, v := range activate {
		(*data)[v].active = true
	}
	for k := 0; k < len(*data); k++ {
		v := (*data)[k]
		if !v.active {
			(*data)[k] = (*data)[len(*data)-1]
			(*data) = (*data)[:len(*data)-1]
			k--
		}
	}
}
func activeNeighbors(n [80]Point) (c int) {
	for _, v := range n {
		if v.active {
			c++
		}
	}
	return
}
func main() {
	data := ReadLines()
	for i := 0; i < 6; i++ {
		increaseBoundary(&data)
		turn(&data)
	}
	fmt.Println(len(data))
}
