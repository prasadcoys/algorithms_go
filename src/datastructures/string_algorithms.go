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
