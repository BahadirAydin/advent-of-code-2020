package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction struct {
	code  string
	value int
}
type Ship struct {
	xPos int
	yPos int
	dir  string
}

func ReadIntegerLines() []Direction {

	f, _ := os.Open("input.txt")
	arr := make([]Direction, 0)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var direction Direction
		text := scanner.Text()
		converted, _ := strconv.Atoi(text[1:])
		direction.code = string(text[0])
		direction.value = converted
		arr = append(arr, direction)
	}
	return arr
}
func updateShip(ship *Ship, dir Direction) {
	code := dir.code
	value := dir.value
	if code == "N" {
		ship.yPos += value
	} else if code == "E" {
		ship.xPos += value
	} else if code == "S" {
		ship.yPos -= value
	} else if code == "W" {
		ship.xPos -= value
	} else if code == "F" {
		dir.code = ship.dir
		updateShip(ship, dir)
		return
	} else if (code == "L" && value == 90) || (code == "R" && value == 270) {
		if ship.dir == "W" {
			ship.dir = "S"
		} else if ship.dir == "S" {
			ship.dir = "E"
		} else if ship.dir == "E" {
			ship.dir = "N"
		} else if ship.dir == "N" {
			ship.dir = "W"
		}
	} else if (code == "R" && value == 90) || (code == "L" && value == 270) {
		if ship.dir == "W" {
			ship.dir = "N"
		} else if ship.dir == "N" {
			ship.dir = "E"
		} else if ship.dir == "E" {
			ship.dir = "S"
		} else if ship.dir == "S" {
			ship.dir = "W"
		}
	} else if value == 180 {
		if ship.dir == "W" {
			ship.dir = "E"
		} else if ship.dir == "N" {
			ship.dir = "S"
		} else if ship.dir == "E" {
			ship.dir = "W"
		} else if ship.dir == "S" {
			ship.dir = "N"
		}
	}
}
func main() {
	data := ReadIntegerLines()
	var ship Ship
	ship.dir = "E"
	for _, v := range data {
		updateShip(&ship, v)
	}
	x := ship.xPos
	y := ship.yPos
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	fmt.Println(x + y)
}
