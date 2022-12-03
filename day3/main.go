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

	// highestPriority := 0
	prioritySum := 0
	for scanner.Scan() {
		line := scanner.Text()

		fullSize := len(line)

		firsCompartment := line[:fullSize/2]
		secondCompartment := line[fullSize/2:]

		firstMap := make(map[rune]bool, 5)
		for _, val := range firsCompartment {
			if firstMap[val] {
				continue
			}
			firstMap[val] = true
		}

		secondMap := make(map[rune]bool, 5)
		for _, val := range secondCompartment {
			if secondMap[val] {
				continue
			}
			secondMap[val] = true
		}

		fmt.Println("---")
		for k := range firstMap {
			fmt.Println(k)
			if secondMap[k] {
				fmt.Printf("Match: %c value: %d\n", k, getPrioroty(k))
				prioritySum += getPrioroty(k)
			}
		}

		// if highestPriority < prioritySum {
		// 	highestPriority = prioritySum
		// }
	}
	fmt.Printf("Priority sum: %d\n", prioritySum)

}

func getPrioroty(char rune) int {
	ascii := int(char)
	if ascii < 91 {
		return ascii - 38
	}

	return ascii - 96
}
