package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key [2]int

func main() {
	file, err := os.Open("../commands.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)

	headX := 0
	headY := 0
	tailX := 0
	tailY := 0
	var bytes []byte
	visited := make(map[key]bool, 0)
	visited[key{tailX, tailY}] = true
	for {
		bytes, _, err = reader.ReadLine()
		if err != nil {
			break
		}

		command := strings.Split(string(bytes), " ")
		dir := command[0]
		dist, _ := strconv.Atoi(command[1])

		if dir == "R" {
			for i := 0; i < dist; i++ {
				headX++
				distX := headX - tailX
				if distX > 1 || distX < -1 {
					tailX = headX - 1
					tailY = headY
				}
				visited[key{tailX, tailY}] = true
			}
		} else if dir == "L" {
			for i := 0; i < dist; i++ {
				headX--
				distX := headX - tailX
				if distX > 1 || distX < -1 {
					tailX = headX + 1
					tailY = headY
				}
				visited[key{tailX, tailY}] = true
			}
		} else if dir == "D" {
			for i := 0; i < dist; i++ {
				headY--
				distY := headY - tailY
				if distY > 1 || distY < -1 {
					tailY = headY + 1
					tailX = headX
				}
				visited[key{tailX, tailY}] = true
			}
		} else if dir == "U" {
			for i := 0; i < dist; i++ {
				headY++
				distY := headY - tailY
				if distY > 1 || distY < -1 {
					tailY = headY - 1
					tailX = headX
				}
				visited[key{tailX, tailY}] = true
			}
		}
		fmt.Printf("-----\n%q\ntail: %d %d\nhead: %d %d\n", command, tailX, tailY, headX, headY)
	}
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}

	fmt.Println(count)

}
