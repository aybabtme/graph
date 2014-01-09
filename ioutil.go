package graph

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type graphScanner struct {
	*bufio.Scanner
}

func newGraphScanner(rd io.Reader) *graphScanner {
	gs := graphScanner{bufio.NewScanner(rd)}
	gs.Scanner.Split(bufio.ScanWords)
	return &gs
}

func (w *graphScanner) NextInt() (int, error) {
	if w.Scan() {
		if err := w.Err(); err != nil {
			return 0, err
		}

		return strconv.Atoi(w.Text())
	}
	return 0, fmt.Errorf("couldn't scan")
}

func (w *graphScanner) NextEdge() (from int, to int, err error) {
	from, err = w.NextInt()
	if err != nil {
		return
	}
	to, err = w.NextInt()
	if err != nil {
		return
	}
	return
}

type weightGraphScanner struct {
	*bufio.Scanner
}

func newWeighGraphScanner(rd io.Reader) *weightGraphScanner {
	ws := weightGraphScanner{bufio.NewScanner(rd)}
	ws.Scanner.Split(bufio.ScanWords)
	return &ws
}

func (w *weightGraphScanner) NextInt() (int, error) {
	if w.Scan() {
		if err := w.Err(); err != nil {
			return 0, err
		}

		return strconv.Atoi(w.Text())
	}
	return 0, fmt.Errorf("couldn't scan")
}

func (w *weightGraphScanner) NextFloat() (float64, error) {
	if w.Scan() {
		if err := w.Err(); err != nil {
			return 0, err
		}

		return strconv.ParseFloat(w.Text(), 64)
	}
	return 0, fmt.Errorf("couldn't scan")
}

func (w *weightGraphScanner) NextEdge() (from int, to int, weight float64, err error) {
	from, err = w.NextInt()
	if err != nil {
		return
	}
	to, err = w.NextInt()
	if err != nil {
		return
	}
	weight, err = w.NextFloat()
	if err != nil {
		return
	}

	return
}
