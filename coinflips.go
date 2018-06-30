package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
)

const (
	head = "head"
	tail = "tail"
)

func coinFlip() string {
	if rand.Intn(2) == 1 {
		return head
	} else {
		return tail
	}
}

var numOfThrows uint
var oneliner bool
var printFormat string

type printer func(throws <-chan string)

var printerFuncs map[string]printer

func init() {
	const (
		numOfThrowsDefault = 10
		numOfTrhowsUsage   = "Number of throws"
	)
	flag.UintVar(&numOfThrows, "n", numOfThrowsDefault, numOfTrhowsUsage+" (shorthand)")
	flag.UintVar(&numOfThrows, "number", numOfThrowsDefault, numOfTrhowsUsage)
	flag.BoolVar(&oneliner, "oneline", false, "Prints throws on one line")
	flag.StringVar(&printFormat, "format", "human", "Output format: 'human', 'csv'")

	printerFuncs = make(map[string]printer)
	printerFuncs["human"] = humanPrinter
	printerFuncs["csv"] = csvPrinter
}

func getFormatString() string {
	digits := int(math.Log10(float64(numOfThrows)) + 1)
	eol := "\n"
	if oneliner {
		eol = "\r"
	}
	return fmt.Sprintf("(%%%dd): %%s     Heads: %%%dd, Tails: %%%dd%s", digits, digits, digits, eol)
}

func humanPrinter(throws <-chan string) {
	format := getFormatString()
	heads, tails, index := 0, 0, 0
	for t := range throws {
		index++
		switch t {
		case head:
			heads++
		case tail:
			tails++
		}
		fmt.Printf(format, index, t, heads, tails)
	}
	if oneliner {
		fmt.Println()
	}
	wg.Done()
}

func csvPrinter(throws <-chan string) {
	bol := true
	for t := range throws {
		if !bol {
			fmt.Print(", ")
		}
		fmt.Print(t)
		bol = false
	}
	fmt.Println()
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	flag.Parse()
	n := int(numOfThrows)
	throws := make(chan string)
	printer, found := printerFuncs[printFormat]
	if !found {
		fmt.Printf("Unrecognized output format %q\n", printFormat)
		flag.Usage()
		os.Exit(1)
	}
	wg.Add(1)
	go printer(throws)
	for ; n > 0; n-- {
		throws <- coinFlip()
	}
	close(throws)
	wg.Wait()
}
