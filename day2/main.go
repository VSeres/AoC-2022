package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName string = "file"

func main() {
	fmt.Println("hello")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Missing file: %s\n", fileName)
		return
	}

	scanner := bufio.NewScanner(file)
	pointSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		guide := strings.Split(line, " ")
		point := getMatchPoint(guide[0], guide[1])
		point += getValue(guide[1])
		fmt.Println(point)
		pointSum += point
	}
	fmt.Printf("sum: %d\n", pointSum)
}

func getMatchPoint(enemy, your string) int {
	// a = x, b = y, c = z
	win, draw := false, false
	switch enemy {
	case "A":
		if your == "X" {
			draw = true
		} else if your == "Y" {
			win = true
		}
	case "B":
		if your == "Y" {
			draw = true
		} else if your == "Z" {
			win = true
		}
	case "C":
		if your == "Z" {
			draw = true
		} else if your == "X" {
			win = true
		}
	}

	if draw {
		return 3
	}

	if win {
		return 6
	}
	return 0
}

func getValue(hand string) int {
	switch hand {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}
