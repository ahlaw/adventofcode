package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type deck struct {
	cards []int
}

func newDeck(size int) deck {
	cards := make([]int, size)
	for i := range cards {
		cards[i] = i
	}
	return deck{cards}
}

func (d *deck) reverse() {
	for i := 0; i < len(d.cards)/2; i++ {
		j := len(d.cards) - i - 1
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *deck) cut(n int) {
	if n < 0 {
		n += len(d.cards)
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (d *deck) increment(n int) {
	newCards := make([]int, len(d.cards))
	for i, card := range d.cards {
		newCards[(i*n)%len(d.cards)] = card
	}
	d.cards = newCards
}

func (d deck) getPos(selected int) int {
	for pos, card := range d.cards {
		if card == selected {
			return pos
		}
	}
	return -1
}

func (d *deck) shuffle(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "deal into new stack" {
			d.reverse()
		} else if strings.HasPrefix(line, "cut") {
			var n int
			fmt.Sscanf(line, "cut %d", &n)
			d.cut(n)
		} else if strings.HasPrefix(line, "deal with increment ") {
			var n int
			fmt.Sscanf(line, "deal with increment %d", &n)
			d.increment(n)
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	deckSize := 10007
	d := newDeck(deckSize)
	d.shuffle(f)
	fmt.Print("Part1: ")
	fmt.Println(d.getPos(2019))
}
