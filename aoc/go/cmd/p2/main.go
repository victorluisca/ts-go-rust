package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func parsePoint(line string) (*Point, error) {
	parts := strings.Split(line, ",")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Point{
		x: x,
		y: y,
	}, nil
}

func parseLine(line string) (*Line, error) {
	parts := strings.Split(line, " -> ")

	p1, err := parsePoint(parts[0])
	if err != nil {
		return nil, err
	}

	p2, err := parsePoint(parts[1])
	if err != nil {
		return nil, err
	}

	return &Line{
		p1: *p1,
		p2: *p2,
	}, nil
}

func isStraight(line *Line) *Line {
	if line.p1.x == line.p2.x || line.p1.y == line.p2.y {
		return line
	}
	return nil
}

func main() {
	lines := []Line{}

	for _, l := range strings.Split(getInput(), "\n") {
		line, err := parseLine(l)
		if err != nil {
			log.Fatal("not able to parse the line")
		}

		if isStraight(line) != nil {
			lines = append(lines, *line)
		}

	}

	fmt.Printf("%+v", lines)
}
