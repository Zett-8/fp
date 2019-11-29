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
	args := os.Args[1:]
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		os.Exit(0)
	}

	if args[0] == "push" {
		fmt.Println(string(out))
		err := printFortune()
		if err != nil {
			fmt.Println("Sorry! No fortune paper because of error!")
		}
		os.Exit(0)
	}

	fmt.Println(string(out))
}


func printFortune() error {
	var quotes []string

	f, err := os.Open("quotes.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		quotes = append(quotes, sc.Text())
	}

	rand.Seed(time.Now().UnixNano())
	ran := int32(len(quotes))

	fmt.Println(crp(quotes[rand.Int31n(ran)]))

	return nil
}


func crp(S string) string {
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