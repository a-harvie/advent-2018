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
	printEmptyGrid(test)
	grid := makeGrid(test)
	printGrid(grid, test)
	fmt.Println(testGrid(grid))

	grid2 := makeGrid(input)
	fmt.Println(testGrid(grid2))

}

func makeGrid(points []point) [][]int {

	_, max := getBoundingBox(points)

	grid := make([][]int, max.y+1)
	for y := range grid {
		grid[y] = make([]int, max.x+2) // example has an extra column
	}
	for y := range grid {
		for x := range grid[y] {
			closest := getClosestPoints(x, y, points)
			if len(closest) == 1 {
				grid[y][x] = closest[0]
			} else {
				grid[y][x] = -1
			}
		}
	}
	return grid
}

func testGrid(grid [][]int) int {

	seen := make(map[int]struct{})
	biggest := 0
	for y := range grid {
		for x := range grid[y] {
			val := grid[y][x]
			if val < 0 {
				continue
			}
			_, seenVal := seen[val]
			if seenVal {
				continue
			}
			seen[val] = struct{}{}

			i := isInfinite(x, y, grid)
			if !i {
				area := getArea(val, grid)
				if area > biggest {
					biggest = area
				}
			}

		}
	}
	return biggest
}

func getArea(val int, grid [][]int) int {
	area := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == val {
				area++
			}
		}
	}
	return area
}

func isInfinite(x int, y int, grid [][]int) bool {
	if x == 0 || y == 0 || y == (len(grid)-1) || x == (len(grid[y])-1) {
		return true
	}

	val := grid[y][x]
	for yy := range grid {
		for xx := range grid[yy] {
			val2 := grid[yy][xx]
			if val == val2 {
				if xx == 0 || yy == 0 || yy == len(grid)-1 || xx == len(grid[yy])-1 {
					return true
				}
			}
		}
	}
	return false
}

func getClosestPoints(x int, y int, points []point) []int {
	closestPoints := make([]int, 0)
	p1 := point{
		x,
		y,
	}

	for i := 0; i < len(points); i++ {

		p2 := points[i]
		dist := getManhattanDistance(p1, p2)
		if len(closestPoints) == 0 {
			closestPoints = append(closestPoints, i)
		} else {
			previousDist := getManhattanDistance(p1, points[closestPoints[0]])
			if dist < previousDist {
				closestPoints = make([]int, 1)
				closestPoints[0] = i
			} else if dist == previousDist {
				closestPoints = append(closestPoints, i)
			}
		}
	}
	return closestPoints
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

func printEmptyGrid(points []point) {
	_, max := getBoundingBox(points)

	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			o := ""
			for val, point := range points {
				if y == point.y && x == point.x {
					o = string(val + 65)
					break
				}
			}

			if o == "" {
				o = "."
			}
			fmt.Print(o)
		}
		fmt.Println()
	}
}

func printGrid(grid [][]int, points []point) {
	var o string
	for y := range grid {
		for x := range grid[y] {
			c := grid[y][x]

			if c < 0 {
				o = "."
			} else {
				offset := 97
				for _, point := range points {
					if y == point.y && x == point.x {
						offset = 65
						break
					}
				}
				o = string(c + offset)
			}
			fmt.Print(o)
		}
		fmt.Println()
	}
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
