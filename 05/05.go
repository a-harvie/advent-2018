package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

var test = getInput("test.txt")
var input = getInput("input.txt")
var testResults = 10

func main() {
	in := test
	units := make([]rune, 4)
	for i := 0; i < 4; i++ {
		units[i] = rune('a' + i)
	}
	fmt.Println(string(units))

	shortest := len(in)
	for _, u := range units {
		candidate := removeUnits(u, in)
		fmt.Println(string(in), string(candidate))
		for {
			reaction, reacted := react(candidate)

			if reacted {
				candidate = reaction
			} else {
				fmt.Println(string(reaction))
				fmt.Println(len(reaction))
				break
			}
		}
		if len(candidate) < shortest {
			shortest = len(candidate)
		}
	}
	fmt.Println(shortest)

	units = make([]rune, 26)
	for i := 0; i < 26; i++ {
		units[i] = rune('a' + i)
	}
	in = input
	shortest = len(in)
	for _, u := range units {
		candidate := removeUnits(u, in)

		for {
			reaction, reacted := react(candidate)

			if reacted {
				candidate = reaction
			} else {
				break
			}
		}
		if len(candidate) < shortest {
			shortest = len(candidate)
		}
	}
	fmt.Println(shortest)

}

func removeUnits(u rune, r []rune) []rune {
	if unicode.IsUpper(u) {
		u = unicode.ToLower(u)
	}

	out := make([]rune, 0)
	for i := 0; i < len(r); i++ {
		if unicode.ToLower(r[i]) != u {
			out = append(out, r[i])
		}
	}

	return out
}

func react(r []rune) ([]rune, bool) {
	for i := 1; i < len(r); i++ {
		if reacts := unitsReact(r[i-1], r[i]); reacts {
			s := i + 1
			if s == len(r) {
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
