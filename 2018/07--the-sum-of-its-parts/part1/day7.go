package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	lines := strings.Split(string(b), "\n")
	stepDeps := map[string]map[string]bool{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := r.FindStringSubmatch(l)
		dep := matches[1]
		step := matches[2]
		if stepDeps[dep] == nil {
			stepDeps[dep] = map[string]bool{}
		}
		if stepDeps[step] == nil {
			stepDeps[step] = map[string]bool{}
		}
		stepDeps[step][dep] = true
	}

	for {
		didWork := false
		for step := 'A'; step <= 'Z'; step++ {
			deps, ok := stepDeps[string(step)]
			if !ok {
				continue
			}
			if len(deps) > 0 {
				didWork = true
			} else {
				fmt.Printf("%c", step)
				for _, deps := range stepDeps {
					delete(deps, string(step))
				}
				delete(stepDeps, string(step))
				didWork = true
				break
			}
		}
		if !didWork {
			break
		}
	}
	fmt.Println()
}
