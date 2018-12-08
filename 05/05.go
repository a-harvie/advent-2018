package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

var test = getInput("../input/05_test.txt")
var input = getInput("../input/05.txt")
var testResults = 10

func main() {
	reacted := getFullyReacted(test)
	fmt.Println(len(reacted))
	reacted = getFullyReacted(input)
	fmt.Println(len(reacted))
}

func getFullyReacted(r []rune) []rune {
	for {
		reaction, reacted := react(r)

		if reacted {
			r = reaction
		} else {
			// fmt.Println(string(reaction))
			// fmt.Println(len(reaction))
			break
		}
	}

	return r
}

func react(r []rune) ([]rune, bool) {
	for i := 1; i < len(r); i++ {
		if reacts := unitsReact(r[i-1], r[i]); reacts {
			s := i + 1
			if s > len(r) {
				s = len(r)
			}
			return append(r[0:i-1], r[s:len(r)]...), true
		}
	}

	return r, false
}

func unitsReact(u1 rune, u2 rune) bool {
	if unicode.IsLower(u1) && unicode.IsLower(u2) {
		return false
	}

	if unicode.IsUpper(u1) && unicode.IsUpper(u2) {
		return false
	}

	if unicode.ToLower(u1) != unicode.ToLower(u2) {
		return false
	}

	return true
}

func getInput(filename string) []rune {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oh noes", err)
	}

	return bytes.Runes(data)
}
