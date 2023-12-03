package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mm = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func find_coordinate(s string) int {
	ns := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(s, "")
	coo, _ := strconv.Atoi(ns[0:1] + ns[len(ns)-1:])
	return coo
}

func transform(s string) string {
	fmt.Println(s)
	i := 1
	for {
		subs := s[0:i]
		for number, digit := range mm {
			r, _ := regexp.Compile(number)
			if r.MatchString(subs) {
				ns := strings.Replace(s, number, strconv.Itoa(digit), 1)
				s = ns
				i = 1
				break
			}
		}
		i++
		if i > len(s) {
			break
		}
	}
	fmt.Println(s)
	return s
}

func he(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// lines of text
	// each calibration value = combine first digit w/ last digit = single two digit number
	// read file
	// for each line, transform text number to digit number
	// for each line, combine the first and the last digits
	// add iteration item to sum
	// OBS: a single digit is both the first and the last

	var input string
	var sum = 0
	input = os.Args[1]

	f, err := os.Open(input)
	fmt.Printf("Input is: %s\n", input)
	he(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tline := transform(line)
		coo := find_coordinate(tline)
		fmt.Println(coo)
		sum += coo
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
