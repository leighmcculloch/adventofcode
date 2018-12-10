package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x, y   int
	vx, vy int
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("position=< *([-0-9]+), *([-0-9]+)> velocity=< *([-0-9]+), *([-0-9]+)>")
	lines := strings.Split(string(b), "\n")
	points := []point{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := r.FindStringSubmatch(l)
		p := point{
			x:  atoi(matches[1]),
			y:  atoi(matches[2]),
			vx: atoi(matches[3]),
			vy: atoi(matches[4]),
		}
		points = append(points, p)
	}

	_, minH := getWidthHeight(points)
	lastPoints := make([]point, len(points))
	for t := 0; ; t++ {
		copy(lastPoints, points)
		for i := range points {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
		_, h := getWidthHeight(points)
		if h < minH {
			minH = h
		} else {
			fmt.Printf("time is at %d seconds\n", t)
			display(lastPoints)
			break
		}
	}
}

func getWidthHeight(points []point) (int, int) {
	minX, maxX, minY, maxY := points[0].x, points[0].x, points[0].y, points[0].y
	for _, p := range points {
		minX = min(minX, p.x)
		maxX = max(maxX, p.x)
		minY = min(minY, p.y)
		maxY = max(maxY, p.y)
	}
	return maxX - minX + 1, maxY - minY + 1
}

func display(points []point) {
	minX, maxX, minY, maxY := points[0].x, points[0].x, points[0].y, points[0].y
	type staticPoint struct {
		x, y int
	}
	board := map[staticPoint]bool{}
	for _, p := range points {
		minX = min(minX, p.x)
		maxX = max(maxX, p.x)
		minY = min(minY, p.y)
		maxY = max(maxY, p.y)
		board[staticPoint{p.x, p.y}] = true
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if board[staticPoint{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
