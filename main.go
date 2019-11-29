package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var quotes []string

	args := os.Args[1:]
	out, err := exec.Command("git push", args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		os.Exit(0)
	}

	if args[0] == "push" {
		f, err := os.Open("quotes.txt")
		if err != nil {
			fmt.Println(string(out))
			os.Exit(0)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			quotes = append(quotes, sc.Text())
		}

		rand.Seed(time.Now().UnixNano())
		ran := int32(len(quotes))

		fmt.Println(de(quotes[rand.Int31n(ran)]))
		os.Exit(0)
	}

	fmt.Println(string(out))
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