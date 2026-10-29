package asciiartweb

import (
	"bufio"
	"os"
)

func Loadbanner(banner string) ([]string, error) {
	file, err := os.Open("Banner/" + banner + ".txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	content := bufio.NewScanner(file)

	for content.Scan() {
		lines = append(lines, content.Text())
	}

	if err := content.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
