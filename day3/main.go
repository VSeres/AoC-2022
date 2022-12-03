package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "file"

func main() {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error while opening: %s", fileName)
	}

	scanner := bufio.NewScanner(file)

	prioritySum := 0
	badgePrioritySum := 0

	rucksacks := make([]map[rune]bool, 3)
	rucksackIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		rucksacks[rucksackIndex] = make(map[rune]bool, 3)

		prioritySum += getPriorotyFromLine(line)

		for _, v := range line {
			if !rucksacks[rucksackIndex][v] {
				rucksacks[rucksackIndex][v] = true
			}
		}

		rucksackIndex++
		if rucksackIndex == 3 {
			for k := range rucksacks[0] {
				if !rucksacks[1][k] {
					continue
				}

				if rucksacks[2][k] {
					badgePrioritySum += getPrioroty(byte(k))
				}
			}
			rucksackIndex = 0
		}

		// if highestPriority < prioritySum {
		// 	highestPriority = prioritySum
		// }
	}
	fmt.Printf("Priority sum: %d\n", prioritySum)
	fmt.Printf("Bage priority sum: %d\n", badgePrioritySum)

}

func getPrioroty(char byte) int {
	// fmt.Println(char)
	ascii := int(char)
	if ascii < 91 {
		return ascii - 38
	}

	return ascii - 96
}

func getPriorotyFromLine(line string) int {
	fullSize := len(line)
	half := fullSize / 2
	for i := 0; i < half; i++ {
		for j := half; j < fullSize; j++ {
			if line[i] == line[j] {
				return getPrioroty(line[i])
			}
		}
	}

	return -1
}
