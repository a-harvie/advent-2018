package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var tests = getInput("../input/04_test.txt")
var input = getInput("../input/04.txt")
var testResults = 240

// map minutes to time spent asleep during that minute
type shiftCount map[int]int

func main() {

	// tl := make([]shift, 0)

	// currentGuard := ""
	shiftCount := getShifts(tests)
	fmt.Println(shiftCount)
	out := countShifts(shiftCount)
	fmt.Println(out)

	shiftCount = getShifts(input)
	// fmt.Println(shiftCount)
	out = countShifts(shiftCount)
	fmt.Println(out)
}

// func printGrid(g [][]int) {
// 	for i := 0; i < len(g); i++ {
// 		fmt.Println(g[i])
// 	}
// }

// func gridCount(g [][]int) int {
// 	c := 0
// 	for i := 0; i < len(g); i++ {
// 		for j := 0; j < len(g[i]); j++ {
// 			if g[i][j] > 1 {
// 				c++
// 			}
// 		}
// 	}

// 	return c
// }

// func updateGrid(g [][]int, c claim) [][]int {
// 	x := c.offset.x
// 	y := c.offset.y
// 	w := c.size.w
// 	h := c.size.h

// 	for i := x; i < x+w; i++ {
// 		for j := y; j < y+h; j++ {
// 			if g[i][j] < 2 {
// 				g[i][j]++
// 			}
// 		}
// 	}

// 	return g
// }

func countShifts(shifts map[string]shiftCount) int {
	sleepyGuard := ""
	sleepyMinutes := 0
	minuteCount := 0

	for guard, shift := range shifts {
		minuteCount = 0
		fmt.Println(shift)
		for _, count := range shift {
			minuteCount += count
		}
		fmt.Printf("Guard %v slept for %v\n", guard, minuteCount)
		if minuteCount > sleepyMinutes {
			sleepyMinutes = minuteCount
			sleepyGuard = guard
		}
		fmt.Printf("Sleepy guard now %v\n", sleepyGuard)
	}
	fmt.Println("Sleepy guard", sleepyGuard)
	sleepyMinutes = 0
	sleepyMinute := 0
	for m, s := range shifts[sleepyGuard] {
		if s > sleepyMinutes {
			sleepyMinute = m
			sleepyMinutes = s
		}
	}
	fmt.Println("Sleepy minute", sleepyMinute, sleepyMinutes)
	sg, _ := strconv.Atoi(sleepyGuard)
	return sg * sleepyMinute
}

func getShifts(lines []string) map[string]shiftCount {
	shifts := make(map[string]shiftCount)
	sleeping := false
	currentGuard := ""
	for index, line := range lines {
		date, action, guard := parseInput(line)

		fmt.Println(date, action, guard)

		if guard != "" {
			currentGuard = guard
			_, found := shifts[guard]
			if !found {
				shifts[guard] = make(map[int]int)
			}
		} else if action == "falls asleep" {
			sleeping = true
		} else if action == "wakes up" {
			if sleeping {
				fellAsleep, _, _ := parseInput(lines[index-1])
				fmt.Println(fellAsleep.Minute(), date.Minute(), currentGuard)
				for i := fellAsleep.Minute(); i < date.Minute(); i++ {
					shifts[currentGuard][i]++
				}
			}
			sleeping = false
		}
	}

	return shifts
}

func parseInput(i string) (time.Time, string, string) {
	d := i[1:strings.LastIndex(i, "]")]
	date, _ := time.Parse("2006-01-02 15:04", d)

	action := i[strings.LastIndex(i, "]")+1:]
	action = strings.TrimSpace(action)

	guard := ""
	if strings.Contains(action, "#") {
		r, _ := regexp.Compile("([0-9]+)")
		guard = r.FindString(action)
	}

	return date, action, guard
}

func getInput(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oh noes", err)
	}

	lines := strings.Split(string(data), "\n")
	sort.Strings(lines)

	return lines
}
