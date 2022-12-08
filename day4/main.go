package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "file"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error while opening: %s\n", fileName)
		return
	}

	scanner := bufio.NewScanner(file)
	enclosing := 0
	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		sections := strings.Split(line, ",")

		elfOneS := strings.Split(sections[0], "-")
		elfTwoS := strings.Split(sections[1], "-")

		elfOne := make([]int, 2)
		elfTwo := make([]int, 2)

		for i := 0; i < 2; i++ {
			elfOne[i], _ = strconv.Atoi(elfOneS[i])
			elfTwo[i], _ = strconv.Atoi(elfTwoS[i])
		}

		if elfOne[0] <= elfTwo[0] && elfOne[1] >= elfTwo[1] {
			/*
				A-------B
				  C----D
			*/
			enclosing++
		} else if elfOne[0] >= elfTwo[0] && elfOne[1] <= elfTwo[1] {
			/*
				A----B
				  C-------D
			*/
			enclosing++
		}

		if elfOne[0] <= elfTwo[1] && elfOne[1] >= elfTwo[0] {
			overlaps++
		}
	}
	fmt.Println("enclosing: ", enclosing)
	fmt.Println("Overlaps: ", overlaps)
}
