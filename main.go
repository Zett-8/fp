package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/zett-8/fp/terminal"
)

func main() {
	args := os.Args[1:]
	args = append([]string{"push"}, args...)

	out, err := exec.Command("git", args...).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		os.Exit(0)
	}

	err = printFortune()
	if err != nil {
		fmt.Println("Sorry! No fortune paper because of error!")
	}
	os.Exit(0)
}

func printFortune() error {
	width, err := terminal.Width()
	if err != nil {
		return err
	}

	qb, err := Asset("quotes.txt")
	if err != nil {
		return err
	}
	quotes := strings.Split(string(qb), "\n")

	rand.Seed(time.Now().UnixNano())
	ran := int32(len(quotes))
	q := crp(quotes[rand.Int31n(ran)])

	whiteSpace := math.Max((float64(width)-float64(len(q)))/2, 0)

	fmt.Println("")
	fmt.Println("  %%" + strings.Repeat("~", int(width)-8) + "%%  ")
	fmt.Println("")
	fmt.Println(strings.Repeat(" ", int(whiteSpace)) + q)
	fmt.Println("")
	fmt.Println("  %%" + strings.Repeat("~", int(width)-8) + "%%  ")
	fmt.Println("")

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
