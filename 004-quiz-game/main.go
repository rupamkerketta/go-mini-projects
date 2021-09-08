package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheynewallace/tabby"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Answer int

const (
	Right Answer = iota
	Wrong
	Invalid
)

type Stat struct {
	n1, n2, sum             int
	diff, rawInput, avgTime string
	answer                  Answer
}

var divs = []time.Duration{
	time.Duration(1), time.Duration(10), time.Duration(100), time.Duration(1000)}

func round(d time.Duration, digits int) time.Duration {
	switch {
	case d > time.Second:
		d = d.Round(time.Second / divs[digits])
	case d > time.Millisecond:
		d = d.Round(time.Millisecond / divs[digits])
	case d > time.Microsecond:
		d = d.Round(time.Microsecond / divs[digits])
	}
	return d
}

func main() {
	// Emojis
	rightIco, wrongIco, invalidIco := "🟩", "🟧", "🟥"

	fmt.Println(" --- Quiz Game --- ", time.Now().Local().String())
	fmt.Println("")

	// Flag for taking the upper limit for the random numbers
	flg := flag.Int("limit", 10, "Set the upper limit for the random numbers")
	flag.Parse()

	// List for storing the stats
	statList := make([]Stat, 0, 200)

	// stat var
	s := Stat{}

	rand.Seed(time.Now().UnixNano())

	// Standard input - keyboard
	reader := bufio.NewReader(os.Stdin)

	// True/False Channel
	flag := make(chan bool)

	// Counter var representing the no. of questions
	var c int

	// Correct var representing the no. of questions answered correctly
	var correct int

	// Var for storing the initial time
	var t1 time.Time

Loop:
	for c < 11 {
		n1, n2 := rand.Intn(*flg+1), rand.Intn(*flg+1)

		s.n1, s.n2 = n1, n2

		fmt.Printf(fmt.Sprintf(" %d\t+\t%d\t->\t", n1, n2))

		// Set the timer functio to a specific time limit
		timer := time.NewTimer(5 * time.Second)

		// Store the starting time
		t1 = time.Now()

		go func(startTime time.Time) {
			input, err1 := reader.ReadString('\n')
			if err1 != nil {
				fmt.Println("Error while taking the input", err1)
			}
			s.rawInput = strings.TrimSpace(input)

			// Send a bool value when the user enters a value
			flag <- true

			// Stop the timer on user input
			timer.Stop()

			// Find the difference or the time elapsed - basically how long the user
			// took to respond.
			s.diff = round(time.Now().Sub(startTime), 3).String()

			sum, err2 := strconv.Atoi(strings.TrimSpace(input))
			if err2 != nil {
				fmt.Println("Invalid input!!!")
				fmt.Println("")
				s.answer = Invalid
				s.diff = "\tnil"
			} else {
				// Increment the correct var if the question is answered correctly
				if sum == n1+n2 {
					correct++
					s.answer = Right
				} else {
					s.answer = Wrong
				}
			}

			// Compute the sum
			s.sum = n1 + n2

			// Appending stat to statList
			statList = append(statList, s)
		}(t1)

		// Wait for an input from a channel
		select {
		case <-flag:
			// Question counter increment
			c++
		case timeOut := <-timer.C:
			fmt.Println("\n --- TIME OUT !! --- ", timeOut.Local().String())
			break Loop
		}
	}
	precent := fmt.Sprintf("(%.2f %%)", float64(correct)/float64(c)*100)
	fmt.Println("Num of questions answered correctly:", correct, "out of", c, precent)
	fmt.Println("")

	// Super simple library for making go tables
	t := tabby.New()
	t.AddHeader("Num 1", "Num 2", "Expected Input", "Raw Input", "Status", "Time taken")

	avgTime, _ := time.ParseDuration("0s")
	var icon string
	for _, i := range statList {
		switch i.answer {
		case 0:
			icon = rightIco
		case 1:
			icon = wrongIco
		case 2:
			icon = invalidIco
		default:
			icon = ""
		}
		currTimeDiff, _ := time.ParseDuration(i.diff)
		avgTime += currTimeDiff
		t.AddLine(i.n1, i.n2, i.sum, i.rawInput, icon, i.diff)
	}
	t.Print()

	fmt.Println("\nAverage time taken: ", avgTime/time.Duration(c))
}
