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

	f, err := os.Open("quotes.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		quotes = append(quotes, de(sc.Text()))
	}

	rand.Seed(time.Now().UnixNano())
	ran := int32(len(quotes))

	fmt.Println(quotes[rand.Int31n(ran)])
}

//func en(S string) string {
//	var n string
//	for _, s := range S {
//		if string(s) == " " {
//			n += " "
//		} else {
//			n += string(s - rune(1))
//		}
//	}
//	return n
//}

func de(S string) string {
	var n string
	for _, s := range S {
		if string(s) == " " {
			n += " "
		} else {
			n += string(s + rune(1))
		}
	}
	return n
}