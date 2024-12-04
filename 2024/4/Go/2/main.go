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
	fmt.Println(masX(characters))
}

const (
	first  int = 1
	second int = -1

	M rune = 77
	A rune = 65
	S rune = 83
)

func masX(characters [][]rune) (amount uint) {
	for vertical := 1; vertical < len(characters)-1; vertical++ {
		for horizontal := 1; horizontal < len(characters[vertical])-1; horizontal++ {
			if (characters)[vertical][horizontal] != A {
				continue
			}
			if diagonalMas(characters, vertical, horizontal, first) && diagonalMas(characters, vertical, horizontal, second) {
				amount++
			}
		}
	}
	return
}

func diagonalMas(characters [][]rune, vertical, horizontal, direction int) bool {
	if characters[vertical-direction][horizontal+1] == M {
		if characters[vertical+direction][horizontal-1] == S {
			return true
		}
	}
	if characters[vertical-direction][horizontal+1] == S {
		if characters[vertical+direction][horizontal-1] == M {
			return true
		}
	}
	return false
}
