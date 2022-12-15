package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key [2]int

type Coord struct {
	x int
	y int
}

const sizeX = 416
const sizeY = 690

var idk = 0

func main() {
	file, err := os.Open("../commands.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	//init rope
	rope := make([]Coord, 10)
	for i := 0; i < 10; i++ {
		rope[i] = Coord{0, 0}
	}

	var bytes []byte

	visited := make(map[Coord]bool, 0)
	visited[rope[0]] = true
	vis := make([][]string, sizeY)
	for i := 0; i < sizeY; i++ {
		vis[i] = make([]string, sizeX)
		for j := 0; j < sizeX; j++ {
			vis[i][j] = "."
		}
	}
	for {
		bytes, _, err = reader.ReadLine()
		if err != nil {
			break
		}

		command := strings.Split(string(bytes), " ")
		dir := command[0]
		dist, _ := strconv.Atoi(command[1])
		for j := 0; j < dist; j++ {
			for i, v := range rope {
				if i == 0 {
					if dir == "U" {
						v.y += 1
					} else if dir == "R" {
						v.x += 1
					} else if dir == "D" {
						v.y -= 1
					} else {
						v.x -= 1
					}
				} else {
					distVec := Coord{rope[i-1].x - v.x, rope[i-1].y - v.y}
					if distVec.y == 2 && distVec.x == 0 {
						v.y += 1
					} else if distVec.y == -2 && distVec.x == 0 {
						v.y -= 1
					} else if distVec.x == 2 && distVec.y == 0 {
						v.x += 1
					} else if distVec.x == -2 && distVec.y == 0 {
						v.x -= 1
					} else if distVec.x > 1 && distVec.y >= 1 || distVec.x >= 1 && distVec.y > 1 { // rigt up
						v.x += 1
						v.y += 1
					} else if distVec.x < -1 && distVec.y >= 1 || distVec.x <= -1 && distVec.y > 1 { // left up
						v.x -= 1
						v.y += 1
					} else if distVec.x > 1 && distVec.y <= -1 || distVec.x >= 1 && distVec.y < -1 { // right down
						v.x += 1
						v.y -= 1
					} else if distVec.x < -1 && distVec.y <= -1 || distVec.x <= -1 && distVec.y < -1 { // left down
						v.x -= 1
						v.y -= 1
					}
				}
				rope[i] = v
				if i == 9 {
					visited[v] = true
				}
			}
		}
	}
	count := 0
	for k, v := range visited {
		if v {
			count++
			vis[sizeY/2-k.y][k.x+sizeX/2] = "#"
		}
	}
	vis[sizeY/2][sizeX/2] = "s"
	printV(vis)
	fmt.Println(count)
}

func printV(vis [][]string) {
	for y := 0; y < sizeY; y++ {
		fmt.Printf("%2d ", y)
		fmt.Println(vis[y])
	}
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			vis[i][j] = "."
		}
	}
}
