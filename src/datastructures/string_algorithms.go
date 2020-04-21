package datastructures

import (
	"strings"
)

func ZigZagText(text string, num int) string {
	zigzagtext := ""
	texts := buildZigZagArray(text, num)
	for _, chars := range texts {
		for _, char := range chars {
			zigzagtext = zigzagtext + char
		}
	}
	return zigzagtext
}

func buildZigZagArray(text string, num int) [][]string {
	// odd := false
	texts := make([][]string, num)
	textCounter := 0
	for {
		if textCounter > len(text)-1 {
			break
		}
		for i := 0; i < num; i++ {
			if textCounter > len(text)-1 {
				break
			}
			if texts[i] == nil {
				texts[i] = []string{}
			}
			texts[i] = append(texts[i], string(text[textCounter]))
			textCounter = textCounter + 1
		}
		for j := num - 2; j > 0; j-- {
			if textCounter > len(text)-1 {
				break
			}
			texts[j] = append(texts[j], string(text[textCounter]))
			textCounter = textCounter + 1
		}
	}

	return texts
}

func FindLongestCommonPrefix(texts []string) string {
	if len(texts) < 1 {
		return ""
	}
	positions := []rune{}
	endPosition := -1
	for _, character := range texts[0] {
		positions = append(positions, character)
		endPosition = endPosition + 1
	}
	for index, text := range texts {
		if index == 0 {
			continue
		}
		currPosition := -1
		if len(text)-1 < endPosition {
			endPosition = len(text) - 1
		}
		for _, character := range text {

			if currPosition+1 > endPosition {
				break
			}
			if character != positions[currPosition+1] {
				endPosition = currPosition
				break
			}
			currPosition = currPosition + 1
		}
	}

	return texts[0][:endPosition+1]
}

func IsValidParantheses(text string) bool {

	tokenMap := make(map[string]string, 3)
	tokenMap[")"] = "("
	tokenMap["]"] = "["
	tokenMap["}"] = "{"
	tokens := StringStack{}
	for _, token := range strings.Split(text, "") {
		matchingToken, ok := tokenMap[token]

		if ok && !tokens.isEmpty() && tokens.peek() == matchingToken {
			tokens.pop()
			continue
		}
		tokens.push(token)

	}

	return tokens.isEmpty()
}

func GetParanthesesCombination(num int) []string {

	parantheses := []string{}
	if num == 1 {
		return []string{"()"}
	}
	paranMap := make(map[string]bool, 0)
	for _, paranthesis := range GetParanthesesCombination(num - 1) {
		if paranMap[paranthesis+"()"] != true {
			parantheses = append(parantheses, paranthesis+"()")
			paranMap[paranthesis+"()"] = true
		}
		if paranMap["()"+paranthesis] != true {
			parantheses = append(parantheses, "()"+paranthesis)
			paranMap["()"+paranthesis] = true
		}
		if paranMap["("+paranthesis+")"] != true {
			parantheses = append(parantheses, "("+paranthesis+")")
			paranMap["("+paranthesis+")"] = true
		}
		individualBrackets := BracketEverySingleparanthesis(paranthesis)
		for _, bracket := range individualBrackets {
			if paranMap[bracket] != true {
				parantheses = append(parantheses, bracket)
				paranMap[bracket] = true
			}
		}
	}

	return parantheses
}

func BracketEverySingleparanthesis(text string) []string {
	output := []string{}
	i := 0
	for {
		if i > len(text)-1 {
			break
		}
		currentPositionOfBracket := i + strings.Index(text[i:], "()")
		output = append(output, text[0:currentPositionOfBracket]+"("+text[currentPositionOfBracket:currentPositionOfBracket+2]+")"+text[currentPositionOfBracket+2:])
		i = currentPositionOfBracket + i + 2
	}
	return output
}

func letterCasePermutation(S string) []string {
	permutations := []string{}

	letters := strings.Split(S, "")
	recursivelyAddLetterCasePermutations("", letters, &permutations)

	return permutations

}

func recursivelyAddLetterCasePermutations(currentString string, letters []string, permutations *[]string) {
	if len(letters) == 0 {
		(*permutations) = append(*permutations, currentString)
		return
	}
	character := []rune(letters[0])
	if character[0] >= 48 && character[0] <= 57 {
		recursivelyAddLetterCasePermutations(currentString+letters[0], letters[1:], permutations)
	} else {
		recursivelyAddLetterCasePermutations(currentString+strings.ToUpper(letters[0]), letters[1:], permutations)
		recursivelyAddLetterCasePermutations(currentString+strings.ToLower(letters[0]), letters[1:], permutations)
	}

}
