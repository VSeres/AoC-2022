package main

import (
	"fmt"
	"os"
)

const fileName = "file"

/*
first task 4
second task 14
*/
const chunk = 14

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error while opening: %s\n", fileName)
		return
	}

	var offset int64 = 0
	buff := make([]byte, chunk)
	for {
		_, err = file.ReadAt(buff, offset)
		if err != nil {
			fmt.Println(err)
			break
		}
		if !duplicate(buff) {
			break
		}
		offset += 1
	}
	fmt.Println(offset + chunk)

}

func duplicate(buff []byte) bool {
	for i := 0; i < chunk; i++ {
		for j := i + 1; j < chunk; j++ {
			if buff[i] == buff[j] {
				return true
			}
		}
	}
	return false
}
