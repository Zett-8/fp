package quote

import (
	"bufio"
	"os"
)

func Import() ([]string, error) {
	var quotes []string

	f, err := os.Open("../quotes.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		quotes = append(quotes, sc.Text())
	}

	return quotes, nil
}
