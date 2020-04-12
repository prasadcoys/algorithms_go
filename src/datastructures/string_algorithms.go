package datastructures

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