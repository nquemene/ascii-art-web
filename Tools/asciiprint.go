package asciiartweb

import (
	"strings"
)

func AsciiPrint(str string, lines []string) string {
	var sb strings.Builder

	slicestr := strings.Split(str, "\n")

	for a := 0; a < len(slicestr); a++ {
		if len(slicestr[a]) == 0 {
			sb.WriteString("\n")
			continue
		}

		runes := []rune(slicestr[a])

		var char [][]string

		for j := 0; j < len(runes); j++ {
			r := runes[j]

			index := int(r) - 32
			start := index*9 + 2
			end := start + 7

			var slice []string
			for line := start; line <= end; line++ {
				if line-1 >= 0 && line-1 < len(lines) {
					slice = append(slice, lines[line-1])
				}
			}

			char = append(char, slice)
		}

		for k := 0; k < 8; k++ {
			for _, block := range char {
				if k < len(block) {
					sb.WriteString(block[k])
				}
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
