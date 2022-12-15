package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var screen = ""

func main() {
	file, err := os.Open("ints.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	var bytes []byte
	clock := 0
	sigStrength := 0
	x := 1
	for {
		bytes, _, err = reader.ReadLine()
		if err != nil {
			break
		}

		line := string(bytes)

		inst := strings.Split(line, " ")

		if inst[0] == "noop" {
			caluculateSignal(&clock, &sigStrength, x)
		} else if inst[0] == "addx" {
			caluculateSignal(&clock, &sigStrength, x)
			caluculateSignal(&clock, &sigStrength, x)
			var v int
			v, err = strconv.Atoi(inst[1])
			if err != nil {
				fmt.Println(err)
			}
			x += v
		}
	}
	fmt.Println(sigStrength)
	for i, v := range screen {
		if i+1%41 == 0 {
			fmt.Println()
		}
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

func caluculateSignal(clock *int, sigStrength *int, x int) {
	*clock += 1
	drawPos := *clock % 40
	if drawPos >= x-1 && drawPos <= x+1 {
		screen += "#"
	} else {
		screen += "."
	}
	if *clock == 20 || *clock == 60 || *clock == 100 || *clock == 140 || *clock == 180 || *clock == 220 {
		*sigStrength += *clock * x
	}
}
