package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var quotes []string

	f, err := os.Open("s_quotes.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		quotes = append(quotes, en(sc.Text()))
	}

	rand.Seed(time.Now().UnixNano())
	ran := int32(len(quotes))

	fmt.Println(quotes[rand.Int31n(ran)])

	fmt.Println(quotes)
}

func en(S string) string {
	var n string
	for _, s := range S {
		n += string(s - rune(1))
	}
	return n
}

func de(S string) string {
	var n string
	for _, s := range S {
		n += string(s + rune(1))
	}
	return n
}