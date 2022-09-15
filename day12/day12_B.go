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
type Waypoint struct {
	xPos int
	yPos int
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
		direction.value = int(converted)
		arr = append(arr, direction)
	}
	return arr
}
func updateShip(ship *Ship, p *Waypoint, dir Direction) {
	code := dir.code
	value := dir.value
	if code == "N" {
		p.yPos += value
	} else if code == "E" {
		p.xPos += value
	} else if code == "S" {
		p.yPos -= value
	} else if code == "W" {
		p.xPos -= value
	} else if code == "F" {
		ship.yPos += value * p.yPos
		ship.xPos += value * p.xPos
	} else if (code == "L" && value == 90) || (code == "R" && value == 270) {
		rotatePoint(90, p)
	} else if (code == "R" && value == 90) || (code == "L" && value == 270) {
		rotatePoint(270, p)
	} else if value == 180 {
		rotatePoint(180, p)
	}
}

func rotatePoint(angle int, p *Waypoint) {
	if angle == 90 {
		tmp := p.xPos
		p.xPos = -1 * p.yPos
		p.yPos = tmp
	} else if angle == 270 {
		tmp := p.xPos
		p.xPos = p.yPos
		p.yPos = -1 * tmp
	} else if angle == 180 {
		p.xPos *= -1
		p.yPos *= -1
	}
}

func main() {
	data := ReadIntegerLines()
	var ship Ship
	var p Waypoint
	p.xPos = 10
	p.yPos = 1
	ship.dir = "E"
	for _, v := range data {
		updateShip(&ship, &p, v)
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
