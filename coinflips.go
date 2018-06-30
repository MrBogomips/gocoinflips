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

func coinFlipGenerator(throws chan<- string) {
	n := int(numOfThrows)
	for ; n > 0; n-- {
		if rand.Intn(2) == 1 {
			throws <- head
		} else {
			throws <- tail
		}
	}
	close(throws)
}

var numOfThrows uint
var oneliner bool
var printFormat string

type printer func(throws <-chan string, done *sync.WaitGroup)

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

func humanPrinter(throws <-chan string, done *sync.WaitGroup) {
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
	done.Done()
}

func csvPrinter(throws <-chan string, done *sync.WaitGroup) {
	bol := true
	for t := range throws {
		if !bol {
			fmt.Print(", ")
		}
		fmt.Print(t)
		bol = false
	}
	fmt.Println()
	done.Done()
}

func main() {
	flag.Parse()
	printer, found := printerFuncs[printFormat]
	if !found {
		fmt.Printf("Unrecognized output format %q\n", printFormat)
		flag.Usage()
		os.Exit(1)
	}
	throws := make(chan string)
	done := sync.WaitGroup{}
	done.Add(1)
	go coinFlipGenerator(throws)
	go printer(throws, &done)
	done.Wait()
}
