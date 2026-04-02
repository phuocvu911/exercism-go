package alphametics

import (
	"errors"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	// Split puzzle into left-hand side words and result
	parts := strings.Split(puzzle, " == ")
	if len(parts) != 2 {
		return nil, errors.New("invalid puzzle format")
	}

	addends := strings.Split(parts[0], " + ")
	result := strings.TrimSpace(parts[1])

	// Collect unique letters and identify leading letters (can't be zero)
	letterSet := make(map[rune]bool)
	leadingSet := make(map[rune]bool)

	allWords := append(addends, result)
	for _, word := range allWords {
		word = strings.TrimSpace(word)
		for i, ch := range word {
			letterSet[ch] = true
			if i == 0 {
				leadingSet[ch] = true
			}
		}
	}

	// Build ordered slice of unique letters
	letters := make([]rune, 0, len(letterSet))
	for ch := range letterSet {
		letters = append(letters, ch)
	}

	if len(letters) > 10 {
		return nil, errors.New("too many unique letters (max 10)")
	}

	// Precompute column coefficients for each letter.
	// For SEND + MORE == MONEY, each letter gets a coefficient:
	// positive for addends, negative for result.
	// A valid assignment satisfies sum(coeff[letter] * digit) == 0.
	coeffs := make(map[rune]int)

	for _, word := range addends {
		word = strings.TrimSpace(word)
		n := len(word)
		place := 1
		for i := n - 1; i >= 0; i-- {
			coeffs[rune(word[i])] += place
			place *= 10
		}
	}

	result = strings.TrimSpace(result)
	n := len(result)
	place := 1
	for i := n - 1; i >= 0; i-- {
		coeffs[rune(result[i])] -= place
		place *= 10
	}

	// Convert to parallel slices for the solver
	numLetters := len(letters)
	letterCoeffs := make([]int, numLetters)
	isLeading := make([]bool, numLetters)
	for i, ch := range letters {
		letterCoeffs[i] = coeffs[ch]
		isLeading[i] = leadingSet[ch]
	}

	// Backtracking solver
	assignment := make([]int, numLetters)
	for i := range assignment {
		assignment[i] = -1
	}
	used := [10]bool{}

	var solve func(idx int, runningSum int) bool
	solve = func(idx int, runningSum int) bool {
		if idx == numLetters {
			return runningSum == 0
		}

		for digit := 0; digit <= 9; digit++ {
			if used[digit] {
				continue
			}
			if digit == 0 && isLeading[idx] {
				continue
			}
			used[digit] = true
			assignment[idx] = digit
			if solve(idx+1, runningSum+letterCoeffs[idx]*digit) {
				return true
			}
			used[digit] = false
			assignment[idx] = -1
		}
		return false
	}

	if !solve(0, 0) {
		return nil, errors.New("no solution found")
	}

	result2 := make(map[string]int, numLetters)
	for i, ch := range letters {
		result2[string(ch)] = assignment[i]
	}
	return result2, nil
}

