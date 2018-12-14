package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	fmt.Println(getRegionTotal(test, 32))
	fmt.Println(getRegionTotal(input, 10000))
}

func getRegionTotal(points []point, limit int) int {
	region := 0
	_, max := getBoundingBox(points)
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			p := point{
				x: x,
				y: y,
			}
			pointTotal := getManhattanTotal(p, points)
			if pointTotal < limit {
				region++
			}
		}
	}

	return region
}

func getManhattanTotal(p point, points []point) int {
	total := 0
	for i := range points {
		total += getManhattanDistance(p, points[i])
	}
	return total
}

func getManhattanDistance(p1 point, p2 point) int {
	x := int(math.Abs(float64(p1.x) - float64(p2.x)))
	y := int(math.Abs(float64(p1.y) - float64(p2.y)))
	return x + y
}

func getBoundingBox(points []point) (point, point) {
	maxX := math.MinInt64
	maxY := math.MinInt64
	minX := math.MaxInt64
	minY := math.MaxInt64

	for i := 0; i < len(points); i++ {
		p := points[i]
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}

	max := point{
		maxX,
		maxY,
	}
	min := point{
		minX,
		minY,
	}

	return min, max
}

func getInput(filename string) []point {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oh noes", err)
	}

	lines := strings.Split(string(data), "\n")
	points := make([]point, 0)
	for _, line := range lines {
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := point{
			x: x,
			y: y,
		}
		points = append(points, p)
	}

	return points
}
