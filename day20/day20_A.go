package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tile struct {
	id  int
	arr [10][10]bool
}

func ReadLines() []Tile {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	tiles := make([]Tile, 0)
	var arr [10][10]bool
	x, y := 0, 0
	var curr Tile
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			curr.arr = arr
			tiles = append(tiles, curr)
			x, y = 0, 0
		} else if text[0] == 'T' {
			curr.id, _ = strconv.Atoi(text[5:9])
		} else {
			for _, v := range text {
				if v == '#' {
					arr[y][x] = true
				} else {
					arr[y][x] = false
				}
				x++
			}
			x = 0
			y++
		}
	}
	curr.arr = arr
	tiles = append(tiles, curr)
	return tiles
}
func findBorder(tile Tile, pos string) (b [10]bool) {
	arr := tile.arr
	if pos == "left" {
		for i := 0; i < 10; i++ {
			b[i] = arr[i][0]
		}
	} else if pos == "right" {
		for i := 0; i < 10; i++ {
			b[i] = arr[i][9]
		}
	} else if pos == "bottom" {
		for i := 0; i < 10; i++ {
			b[i] = arr[9][i]
		}
	} else if pos == "top" {
		for i := 0; i < 10; i++ {
			b[i] = arr[0][i]
		}
	}
	return
}
func transpose(arr [10][10]bool) (transposed [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			transposed[j][i] = arr[i][j]
		}
	}
	return
}
func reverseRows(arr [10][10]bool) (reversed [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			reversed[i][9-j] = arr[i][j]
		}
	}
	return
}
func reverseColumns(arr [10][10]bool) (reversed [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			reversed[j][i] = arr[9-j][i]
		}
	}
	return
}
func rotate(arr [10][10]bool, degree int) (rotated [10][10]bool) {
	if degree == 90 {
		rotated = reverseRows(transpose(arr))
	} else if degree == 270 {
		rotated = reverseColumns(transpose(arr))
	} else if degree == 180 {
		rotated = reverseColumns(reverseRows(arr))
	}
	return
}
func allRotations(arr [10][10]bool) (all [4][10][10]bool) {
	all[0] = arr
	all[1] = rotate(arr, 90)
	all[2] = rotate(arr, 180)
	all[3] = rotate(arr, 270)
	return
}
func flipVertical(arr [10][10]bool) (flipped [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			flipped[9-i][j] = arr[i][j]
		}
	}
	return
}
func flipHorizontal(arr [10][10]bool) (flipped [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			flipped[i][9-j] = arr[i][j]
		}
	}
	return
}
func allOrientations(t Tile) (all [12]Tile) {
	arr := t.arr
	id := t.id
	vertical := flipVertical(arr)
	horizontal := flipHorizontal(arr)
	rot1 := allRotations(arr)
	rot2 := allRotations(vertical)
	rot3 := allRotations(horizontal)
	index := 0
	for ; index < 4; index++ {
		all[index].id = id
		all[index].arr = rot1[index]
	}
	for ; index < 8; index++ {
		all[index].id = id
		all[index].arr = rot2[index-4]
	}
	for ; index < 12; index++ {
		all[index].id = id
		all[index].arr = rot3[index-8]
	}
	return
}
func findLeft(index int, data []Tile) int {
	leftBorder := findBorder(data[index], "left")
	for k, v := range data {
		if k == index {
			continue
		}
		all := allOrientations(v)
		for _, v2 := range all {
			rightBorder := findBorder(v2, "right")
			if leftBorder == rightBorder {
				data[k] = v2
				return 1
			}
		}
	}
	return 0
}
func findRight(index int, data []Tile) int {
	rightBorder := findBorder(data[index], "right")
	for k, v := range data {
		if k == index {
			continue
		}
		all := allOrientations(v)
		for _, v2 := range all {
			leftBorder := findBorder(v2, "left")
			if rightBorder == leftBorder {
				data[k] = v2
				return 1
			}

		}
	}
	return 0
}
func findTop(index int, data []Tile) int {
	topBorder := findBorder(data[index], "top")
	for k, v := range data {
		if k == index {
			continue
		}
		all := allOrientations(v)
		for _, v2 := range all {
			bottomBorder := findBorder(v2, "bottom")
			if topBorder == bottomBorder {
				data[k] = v2
				return 1
			}
		}
	}
	return 0
}
func findBottom(index int, data []Tile) int {
	bottomBorder := findBorder(data[index], "bottom")
	for k, v := range data {
		if k == index {
			continue
		}
		all := allOrientations(v)
		for _, v2 := range all {
			topBorder := findBorder(v2, "top")
			if bottomBorder == topBorder {
				data[k] = v2
				return 1
			}
		}
	}
	return 0
}
func execute(index int, data []Tile) (sum int) {
	sum += findLeft(index, data)
	sum += findRight(index, data)
	sum += findTop(index, data)
	sum += findBottom(index, data)
	//above four functions used can be combined into a single function since the left-right-up-bottom does not have a real meaning in the problem but it works and I don't want to do it :)
	return
}
func main() {
	data := ReadLines()
	res := 1
	for i := 0; i < len(data); i++ {
		n := execute(i, data)
		if n == 2 {
			fmt.Println(data[i].id)
			res *= data[i].id
		}
	}
	fmt.Println(res)
}
