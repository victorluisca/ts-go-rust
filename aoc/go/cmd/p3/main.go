package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	treeCount := 0

	for row, line := range strings.Split(getInput(), "\n") {
		if string(line[(row*3)%len(line)]) == "#" {
			treeCount += 1
		}
	}

	fmt.Printf("treeCount %v\n", treeCount)
}
