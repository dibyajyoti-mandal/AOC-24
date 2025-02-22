package main

import (
	"fmt"

	aoc "github.com/shraddhaag/aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("../text.txt")

	fmt.Println("answer for part 1: ", countXMAS(input))
	fmt.Println("answer for part 2: ", countXMAS2(input))
}

func countXMAS(input []string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		for j, char := range input[i] {
			if char == 'X' {
				count += checkIndex(input, i, j)
			}
		}
	}
	return count
}

func checkIndex(input []string, i, j int) int {
	count := 0

	if j >= 3 {
		if input[i][j] == byte('X') && input[i][j-1] == byte('M') && input[i][j-2] == byte('A') && input[i][j-3] == byte('S') {
			count++
		}
	}

	if j <= (len(input[0]) - 1 - 3) {
		if input[i][j] == byte('X') && input[i][j+1] == byte('M') && input[i][j+2] == byte('A') && input[i][j+3] == byte('S') {
			count++
		}
	}

	if i >= 3 {
		if input[i][j] == byte('X') && input[i-1][j] == byte('M') && input[i-2][j] == byte('A') && input[i-3][j] == byte('S') {
			count++
		}
	}

	if i <= (len(input[0]) - 1 - 3) {
		if input[i][j] == byte('X') && input[i+1][j] == byte('M') && input[i+2][j] == byte('A') && input[i+3][j] == byte('S') {
			count++
		}
	}

	if i >= 3 && j <= (len(input[0])-1-3) {
		if input[i][j] == byte('X') && input[i-1][j+1] == byte('M') && input[i-2][j+2] == byte('A') && input[i-3][j+3] == byte('S') {
			count++
		}
	}

	if i >= 3 && j >= 3 {
		if input[i][j] == byte('X') && input[i-1][j-1] == byte('M') && input[i-2][j-2] == byte('A') && input[i-3][j-3] == byte('S') {
			count++
		}
	}

	if i <= (len(input[0])-1-3) && j <= (len(input[0])-1-3) {
		if input[i][j] == byte('X') && input[i+1][j+1] == byte('M') && input[i+2][j+2] == byte('A') && input[i+3][j+3] == byte('S') {
			count++
		}
	}

	if i <= (len(input[0])-1-3) && j >= 3 {
		if input[i][j] == byte('X') && input[i+1][j-1] == byte('M') && input[i+2][j-2] == byte('A') && input[i+3][j-3] == byte('S') {
			count++
		}
	}

	return count
}

func countXMAS2(input []string) int {
	count := 0
	for i := 1; i < len(input)-1; i++ {
		for j, char := range input[i] {
			if char == 'A' {
				count += checkIndex2(input, i, j)
			}
		}
	}
	return count
}

func checkIndex2(input []string, i, j int) int {
	count := 0

	if j < 1 || j > len(input[i])-2 {
		return 0
	}

	// M.M
	// .A.
	// S.S
	if input[i-1][j+1] == byte('M') && input[i-1][j-1] == byte('M') && input[i+1][j-1] == byte('S') && input[i+1][j+1] == byte('S') {
		count++
	}

	// S.S
	// .A.
	// M.M
	if input[i-1][j+1] == byte('S') && input[i-1][j-1] == byte('S') && input[i+1][j-1] == byte('M') && input[i+1][j+1] == byte('M') {
		count++
	}

	// M.S
	// .A.
	// M.S
	if input[i-1][j+1] == byte('S') && input[i-1][j-1] == byte('M') && input[i+1][j-1] == byte('M') && input[i+1][j+1] == byte('S') {
		count++
	}

	// S.M
	// .A.
	// S.M
	if input[i-1][j+1] == byte('M') && input[i-1][j-1] == byte('S') && input[i+1][j-1] == byte('S') && input[i+1][j+1] == byte('M') {
		count++
	}
	return count
}
