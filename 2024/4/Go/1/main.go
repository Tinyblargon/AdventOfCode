package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("test-data")
	if err != nil {
		log.Fatal(err)
	}

	var characters [][]rune
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		characters = append(characters, runes)
	}
	fmt.Println(scan(characters))
}

const (
	up    int = -1
	down  int = 1
	left  int = -1
	right int = 1

	X rune = 88
	M rune = 77
	A rune = 65
	S rune = 83
)

func scan(characters [][]rune) (amount uint) {
	for vertical := range characters {
		for horizontal := range characters[vertical] {
			if (characters)[vertical][horizontal] != X {
				continue
			}
			if !horizontalOutBounds(characters, vertical, horizontal, right) {
				amount += horizontalScan_Unsafe(characters, vertical, horizontal, right)
			}
			if !horizontalOutBounds(characters, vertical, horizontal, left) {
				amount += horizontalScan_Unsafe(characters, vertical, horizontal, left)
			}
			if !verticalOutBounds(characters, vertical, up) {
				amount += verticalScan_Unsafe(characters, vertical, horizontal, up)
			}
			if !verticalOutBounds(characters, vertical, down) {
				amount += verticalScan_Unsafe(characters, vertical, horizontal, down)
			}
			if !verticalOutBounds(characters, vertical, up) && !horizontalOutBounds(characters, vertical, horizontal, right) {
				amount += scanDiagonal_Unsafe(characters, vertical, horizontal, up, right)
			}
			if !verticalOutBounds(characters, vertical, up) && !horizontalOutBounds(characters, vertical, horizontal, left) {
				amount += scanDiagonal_Unsafe(characters, vertical, horizontal, up, left)
			}
			if !verticalOutBounds(characters, vertical, down) && !horizontalOutBounds(characters, vertical, horizontal, right) {
				amount += scanDiagonal_Unsafe(characters, vertical, horizontal, down, right)
			}
			if !verticalOutBounds(characters, vertical, down) && !horizontalOutBounds(characters, vertical, horizontal, left) {
				amount += scanDiagonal_Unsafe(characters, vertical, horizontal, down, left)
			}
		}
	}
	return
}

func horizontalScan_Unsafe(characters [][]rune, vertical, horizontal, direction int) uint {
	if (characters)[vertical][horizontal+direction] == M {
		if (characters)[vertical][horizontal+direction*2] == A {
			if (characters)[vertical][horizontal+direction*3] == S {
				return 1
			}
		}
	}
	return 0
}

func horizontalOutBounds(characters [][]rune, vertical, horizontal, direction int) bool {
	return horizontal+3*direction < 0 || horizontal+3*direction >= len((characters)[vertical])
}

func verticalScan_Unsafe(characters [][]rune, vertical, horizontal, direction int) uint {
	if (characters)[vertical+direction][horizontal] == M {
		if (characters)[vertical+direction*2][horizontal] == A {
			if (characters)[vertical+direction*3][horizontal] == S {
				return 1
			}
		}
	}
	return 0
}

func verticalOutBounds(characters [][]rune, vertical, direction int) bool {
	return vertical+3*direction < 0 || vertical+3*direction >= len(characters)
}

func scanDiagonal_Unsafe(characters [][]rune, vertical, horizontal, verticalDirection, horizontalDirection int) uint {
	if (characters)[vertical+verticalDirection][horizontal+horizontalDirection] == M {
		if (characters)[vertical+verticalDirection*2][horizontal+horizontalDirection*2] == A {
			if (characters)[vertical+verticalDirection*3][horizontal+horizontalDirection*3] == S {
				return 1
			}
		}
	}
	return 0
}
