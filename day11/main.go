package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []uint64
	operation []string
	test      int
	m1        int
	m2        int
	inspeted  int
}

func main() {
	file, err := os.Open("monkeys.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	monkeys := make([]Monkey, 8)
	monkeyInxex := -1
	simplifier := 1
	for {
		var readBytes []byte
		readBytes, _, err = reader.ReadLine()
		if err != nil {
			break

		}
		line := string(readBytes)
		fields := strings.Fields(line)

		if fields[0] == "Monkey" {
			monkeyInxex += 1
		} else if fields[0] == "Starting" {
			for i := 2; i < len(fields); i++ {
				importance, _ := strconv.Atoi(strings.Trim(fields[i], ","))
				monkeys[monkeyInxex].items = append(monkeys[monkeyInxex].items, uint64(importance))
			}
		} else if fields[0] == "Operation:" {
			monkeys[monkeyInxex].operation = fields[3:]
		} else if fields[0] == "Test:" {
			monkeys[monkeyInxex].test, _ = strconv.Atoi(fields[3])
			simplifier *= monkeys[monkeyInxex].test
		} else if fields[0] == "If" && fields[1] == "true:" {
			monkeys[monkeyInxex].m1, _ = strconv.Atoi(fields[5])
		} else if fields[0] == "If" && fields[1] == "false:" {
			monkeys[monkeyInxex].m2, _ = strconv.Atoi(fields[5])
		}
	}

	for t := 0; t < 10000; t++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for j := 0; j < len(monkey.items); j++ {
				item := &monkey.items[j]
				var v1 uint64
				var tmp int
				if monkey.operation[0] == "old" {
					v1 = *item
				} else {
					tmp, _ = strconv.Atoi(monkey.operation[0])
					v1 = uint64(tmp)
				}
				var v2 uint64
				if monkey.operation[2] == "old" {
					v2 = *item
				} else {
					tmp, _ = strconv.Atoi(monkey.operation[2])
					v2 = uint64(tmp)
				}
				var worryLevel uint64
				if monkey.operation[1] == "+" {
					worryLevel = v1 + v2
				} else if monkey.operation[1] == "*" {
					worryLevel = v1 * v2
				}
				// worryLevel /= 3
				worryLevel %= uint64(simplifier)
				monkey.inspeted += 1
				*item = worryLevel
				if worryLevel%uint64(monkey.test) == 0 {
					monkeys[monkey.m1].items = append(monkeys[monkey.m1].items, worryLevel)
				} else {
					monkeys[monkey.m2].items = append(monkeys[monkey.m2].items, worryLevel)
				}
			}
			monkey.items = make([]uint64, 0)
		}
	}
	max1 := 0
	max2 := 0
	for _, v := range monkeys {
		if v.inspeted > max1 {
			max2 = max1
			max1 = v.inspeted
		} else if v.inspeted > max2 {
			max2 = v.inspeted
		}
	}
	fmt.Println(monkeys)
	fmt.Println(max1, max2, max1*max2)
}
