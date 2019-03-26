package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type step struct {
	ID       string
	children []*step
}

func main() {
	testSrc, testDst := getInput("../input/07_test.txt")

	root := buildTree(testSrc, testDst)
	fmt.Println(root)

	// inputSrc, inputDst := getInput("../input/07.txt")
}

func buildTree(src []string, dst []string) []step {
	rootIDs := getRootIDs(src, dst)
	fmt.Println("root", rootIDs)
	roots := make([]*step, 0)

	for _, ID := range rootIDs {
		s := step{
			ID: ID,
		}

		roots = append()

	}

	return roots
}

func buildFromStep(s *step, src []string, dst []string) *step {
	if len(src) == 0 {
		return s
	}
	remove := make([]int, 0)
	for i := range src {
		if src[i] == s.ID {
			remove = append(remove, i)
			child := step{
				ID: dst[i],
			}
			s.children = append(s.children, &child)
		}
	}
	for i := range remove {
		src = append(src[:i], src[i+1:]...)
		dst = append(dst[:i], dst[i+1:]...)
	}
	for i := range s.children {
		s.children[i] = buildFromStep(s.children[i], src, dst)
	}

	return s
}

func getRootIDs(src []string, dst []string) []string {
	rootIDs := make([]string, 0)
	for _, s := range src {
		isDest := false
		for _, d := range dst {
			if s == d {
				isDest = true
				break
			}
		}
		if !isDest && !sliceContains(rootIDs, s) {
			rootIDs = append(rootIDs, s)
		}
	}
	return rootIDs
}

func sliceContains(sl []string, s string) bool {
	for _, ss := range sl {
		if ss == s {
			return true
		}
	}

	return false
}

func getInput(filename string) ([]string, []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Oh noes", err)
	}

	lines := strings.Split(string(data), "\n")
	sources := make([]string, 0)
	dests := make([]string, 0)
	for _, line := range lines {
		//Step F must be finished before step E can begin.
		r, _ := regexp.Compile("Step ([A-Z]{1}) must be finished before step ([A-Z]{1}) can begin.")
		matches := r.FindStringSubmatch(line)
		sources = append(sources, matches[1])
		dests = append(dests, matches[2])
	}

	return sources, dests
}
