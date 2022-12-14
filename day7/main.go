package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("file")

	if err != nil {
		return
	}

	reader := bufio.NewReader(file)
	var bytes []byte
	path := make([]string, 0)
	sizeMap := make(map[string]int, 0)
	for {
		bytes, _, err = reader.ReadLine()
		if err != nil {
			fmt.Println("ERROR")
			break
		}
		line := string(bytes)

		fragments := strings.Split(line, " ")
		if fragments[0] == "$" {
			if fragments[1] == "cd" {
				if fragments[2] == ".." {
					path = path[:len(path)-1]
				} else if fragments[2] == "/" {
					path = append(path, "/")
				} else {
					path = append(path, fragments[2])
				}
			}
			if fragments[1] == "ls" {
				continue
			}
		} else if fragments[0] != "dir" {
			var size int
			size, err = strconv.Atoi(fragments[0])
			fmt.Println("========")
			fmt.Println(path)
			for i := 1; i < len(path)+1; i++ {
				str := strings.Join(path[:i], "/")
				if len(str) != 1 {
					str = str[1:]
				}
				fmt.Println(str)
				fmt.Println(path[:i])
				sizeMap[str] += size
			}
		}
	}
	count := 0
	for _, v := range sizeMap {
		if v <= 100000 {
			count += v
		}
	}

	fmt.Println(count)
}

func split(c rune) bool {
	return c == '/'
}
