package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var test = getInput("../input/06_test.txt")
var input = getInput("../input/06.txt")
var testResults = 10

type point struct {
	x int
	y int
}

func main() {
	reacted := getFullyReacted(test)
	fmt.Println(len(reacted))
	reacted = getFullyReacted(input)
	fmt.Println(len(reacted))
}

func getInput(filename string) []rune {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oh noes", err)
	}

	lines := strings.Split(string(data), "\n")
	points := make([]point, 0)
	for i, line := range lines {
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := point{
			x: x,
			y: y,
		}
		points = append(points)
	}

	return bytes.Runes(data)
}
