package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "../file.txt"

type stack []byte

func push(s *stack, value byte) {
	*s = append(*s, value)
}

func pop(s *stack) byte {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error while opening: %s\n", fileName)
		return
	}

	scanner := bufio.NewScanner(file)
	stacks := make([]stack, 9)
	for i := 0; i < 7; i++ {
		stacks[i] = make(stack, 0)
	}
	for scanner.Scan() {
		line := scanner.Text()

		if line[1] == '1' {
			scanner.Scan()
			break
		}

		for i := 0; i < 9; i++ {
			fmt.Printf(" %c ", line[i+1+i*3])
			if line[i+1+i*3] != ' ' {
				push(&stacks[i], line[i+1+i*3])
			}
		}
		fmt.Println()
	}
	for _, s := range stacks {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	for scanner.Scan() {
		line := scanner.Text()

		order := strings.Fields(line)

		ammount, _ := strconv.Atoi(order[1])
		from, _ := strconv.Atoi(order[3])
		to, _ := strconv.Atoi(order[5])

		for i := 0; i < ammount; i++ {
			value := pop(&stacks[from-1])
			push(&stacks[to-1], value)
		}
	}

	fmt.Println()
	for _, s := range stacks {
		fmt.Printf("%c", pop(&s))
	}
	println()

}

func printStack(s []stack) {
	maxLen := 0
	for _, s := range s {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	for i := 0; i < maxLen; i++ {
		for j := 0; j < 9; j++ {
			if len(s[j]) > i {
				fmt.Printf("%c ", s[j][i])
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}

}
