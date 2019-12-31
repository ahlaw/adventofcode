package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

type deck struct {
	size   int64
	offset int64
	step   int64
}

func newDeck(size int64) deck {
	return deck{size: size, offset: 0, step: 1}
}

func (d *deck) unreverse() {
	d.step *= -1
	d.offset += 1
	d.offset *= -1
	d.offset %= d.size
}

func (d *deck) uncut(n int64) {
	d.offset += n + d.size
	d.offset %= d.size
}

func (d *deck) unincrement(n int64) {
	inc := big.NewInt(n)
	deckSize := big.NewInt(d.size)
	stepSize := big.NewInt(d.step)
	offsetSize := big.NewInt(d.offset)
	incFactor := big.NewInt(0)
	incFactor.ModInverse(inc, deckSize)
	stepSize.Mul(incFactor, stepSize)
	stepSize.Mod(stepSize, deckSize)
	offsetSize.Mul(incFactor, offsetSize)
	offsetSize.Mod(offsetSize, deckSize)
	d.step = stepSize.Int64()
	d.offset = offsetSize.Int64()
}

func (d deck) getCard(pos, numShuffles int64) int64 {
	a := big.NewInt(d.step)
	b := big.NewInt(d.offset)
	mod := big.NewInt(d.size)
	m := big.NewInt(-1)
	n := big.NewInt(numShuffles)
	p := big.NewInt(pos)
	A := big.NewInt(0)
	B := big.NewInt(1)
	A.Exp(a, n, mod)
	B.ModInverse(a.Sub(a, B), mod)
	m.Add(m, A)
	m.Mul(m, B)
	m.Mul(m, b)
	n.Mul(A, p)
	A.Add(m, n)
	A.Mod(A, mod)
	return A.Int64()
}

func (d *deck) unshuffle(r io.Reader) {
	scanner := bufio.NewScanner(r)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if line == "deal into new stack" {
			d.unreverse()
		} else if strings.HasPrefix(line, "cut") {
			var n int64
			fmt.Sscanf(line, "cut %d", &n)
			d.uncut(n)
		} else if strings.HasPrefix(line, "deal with increment ") {
			var n int64
			fmt.Sscanf(line, "deal with increment %d", &n)
			d.unincrement(n)
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	deckSize := int64(119315717514047)
	numShuffles := int64(101741582076661)
	d := newDeck(deckSize)
	d.unshuffle(f)
	cardPos := int64(2020)
	fmt.Print("Part2: ")
	fmt.Println(d.getCard(cardPos, numShuffles))
}
