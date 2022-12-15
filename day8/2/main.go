package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../file.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	treeMap := make([][]int8, 99)
	for i, _ := range treeMap {
		var lineBytes []byte
		lineBytes, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		line := string(lineBytes)
		treeMap[i] = make([]int8, len(line))
		for c, v := range line {
			treeMap[i][c] = int8(v - '0')
		}
	}
	bestScenic := 0
	for y := 0; y < len(treeMap); y++ {
		for x := 0; x < len(treeMap[0]); x++ {
			var distLeft int = 0
			var distRight int = 0
			var distTop int = 0
			var distBottom int = 0

			for sX := x - 1; sX >= 0; sX-- {
				distLeft++
				if treeMap[y][sX] >= treeMap[y][x] {
					break
				}
			}
			for sX := x + 1; sX < 99; sX++ {
				distRight++
				if treeMap[y][sX] >= treeMap[y][x] {
					break
				}
			}

			for sY := y - 1; sY >= 0; sY-- {
				distTop++
				if treeMap[sY][x] >= treeMap[y][x] {
					break
				}
			}

			for sY := y + 1; sY < 99; sY++ {
				distBottom++
				if treeMap[sY][x] >= treeMap[y][x] {
					break
				}
			}

			scenic := distBottom * distTop * distLeft * distRight
			if scenic > bestScenic {
				bestScenic = scenic
			}
		}
	}
	fmt.Println(bestScenic)
}
