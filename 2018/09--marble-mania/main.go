package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flagPlayerCount := flag.Int("p", 9, "player count")
	flagLastMarble := flag.Int("m", 25, "last marble number")
	flag.Parse()

	playerCount := *flagPlayerCount
	lastMarble := *flagLastMarble

	winningScore := getWinningScore(playerCount, lastMarble)

	fmt.Println(strconv.Itoa(winningScore))
}

func getWinningScore(playerCount, lastMarble int) int {
	marbleCount := lastMarble + 1

	players := make([]int, playerCount)

	m := &marble{number: 0}
	m.clockwise = m
	m.counterClockwise = m

	currentPlayer := 1
	for marbleNumber := 1; marbleNumber < marbleCount; marbleNumber++ {
		if marbleNumber%23 != 0 {
			m = m.Clockwise(2)
			m = m.Insert(marbleNumber)
		} else {
			m = m.CounterClockwise(7)
			m.Remove()
			players[currentPlayer-1] += marbleNumber + m.number
			m = m.clockwise
		}
		currentPlayer++
		if currentPlayer > playerCount {
			currentPlayer = 1
		}
	}

	maxScore := 0
	for _, score := range players {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

type marble struct {
	number           int
	clockwise        *marble
	counterClockwise *marble
}

func (m *marble) Clockwise(steps int) *marble {
	currentMarble := m
	for i := 0; i < steps; i++ {
		currentMarble = currentMarble.clockwise
	}
	return currentMarble
}

func (m *marble) CounterClockwise(steps int) *marble {
	currentMarble := m
	for i := 0; i < steps; i++ {
		currentMarble = currentMarble.counterClockwise
	}
	return currentMarble
}

func (m *marble) Insert(marbleNumber int) *marble {
	newMarble := &marble{
		number:           marbleNumber,
		counterClockwise: m.counterClockwise,
		clockwise:        m,
	}
	m.counterClockwise.clockwise = newMarble
	m.counterClockwise = newMarble
	return newMarble
}

func (m *marble) Remove() {
	m.counterClockwise.clockwise = m.clockwise
	m.clockwise.counterClockwise = m.counterClockwise
}

func (m *marble) String() string {
	s := []string{}
	cm := m
	for {
		s = append(s, strconv.Itoa(cm.number))
		cm = cm.clockwise
		if cm == m {
			break
		}
	}
	return strings.Join(s, " ")
}
