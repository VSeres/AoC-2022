package maind

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
	count := 4*99 - 4
	for y := 0; y < len(treeMap); y++ {
		for x := 0; x < len(treeMap[0]); x++ {
			var maxLeft int8 = 0
			var maxRight int8 = 0
			for sX := 0; sX < len(treeMap[0]); sX++ {
				if sX < x && maxLeft < treeMap[y][sX] {
					maxLeft = treeMap[y][sX]
				} else if sX > x && maxRight < treeMap[y][sX] {
					maxRight = treeMap[y][sX]
				}
			}

			var maxTop int8 = 0
			var maxDown int8 = 0
			for sY := 0; sY < len(treeMap); sY++ {
				if sY < y && maxTop < treeMap[sY][x] {
					maxTop = treeMap[sY][x]
				} else if sY > y && maxDown < treeMap[sY][x] {
					maxDown = treeMap[sY][x]
				}
			}
			if x == 0 || x == 98 || y == 0 || y == 98 {
				continue
			}
			if maxLeft < treeMap[y][x] || maxRight < treeMap[y][x] || maxDown < treeMap[y][x] || maxTop < treeMap[y][x] {
				count++
			}
		}
	}
	fmt.Println(count)
}
