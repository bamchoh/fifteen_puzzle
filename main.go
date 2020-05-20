package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	termbox "github.com/nsf/termbox-go"
)

type board struct {
	board [][]int
	size  int
	x     int
	y     int
}

func newBoard(size int) *board {
	tmp := make([][]int, 4)
	x := 0
	y := 0
	for i := 0; i < size; i++ {
		tmp[i] = make([]int, 4)
		for j := 0; j < size; j++ {
			if (i+1)*(j+1) == size*size {
				tmp[i][j] = -1
				x = j
				y = i
			} else {
				tmp[i][j] = i*size + (j + 1)
			}
		}
	}
	return &board{board: tmp, size: size, x: x, y: y}
}

func (b *board) swap(i, j, k, l int) {
	v := b.board[i][j]
	b.board[i][j] = b.board[k][l]
	b.board[k][l] = v
	b.x, b.y = l, k
}

func (b *board) draw() {
	coldef := termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	x := 0
	y := 0
	hr := func() {
		termbox.SetCell(0, y, '+', coldef, coldef)
		for x := 1; x < b.size*3; x = x + 3 {
			termbox.SetCell(x+0, y, '-', coldef, coldef)
			termbox.SetCell(x+1, y, '-', coldef, coldef)
			termbox.SetCell(x+2, y, '+', coldef, coldef)
		}
		y++
	}
	hr()
	for _, line := range b.board {
		x = 0
		termbox.SetCell(x, y, '|', coldef, coldef)
		for _, cell := range line {
			if cell == -1 {
				termbox.SetCell(x+1, y, ' ', coldef, coldef)
				termbox.SetCell(x+2, y, ' ', coldef, coldef)
				termbox.SetCell(x+3, y, '|', coldef, coldef)
			} else {
				s := fmt.Sprintf("%2d|", cell)
				termbox.SetCell(x+1, y, rune(s[0]), coldef, coldef)
				termbox.SetCell(x+2, y, rune(s[1]), coldef, coldef)
				termbox.SetCell(x+3, y, rune(s[2]), coldef, coldef)
			}
			x += 3
		}
		y++
		hr()
	}
	termbox.Flush()
}

func scan() (string, error) {
	if !sc.Scan() {
		return "", errors.New("Unexpected error")
	}
	v := sc.Text()
	switch v {
	case "a", "s", "d", "w":
		return v, nil
	default:
		return "", fmt.Errorf("Invalid rune:%s", v)
	}
}

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func main() {
	size := 4
	b := newBoard(size)
	var err error
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		b.draw()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft:
				// left
				i, j := b.y, b.x
				if j+1 >= b.size {
					b.draw()
					continue
				}
				b.swap(i, j, i, j+1)
			case termbox.KeyArrowRight:
				// right
				i, j := b.y, b.x
				if j-1 < 0 {
					b.draw()
					continue
				}
				b.swap(i, j, i, j-1)
			case termbox.KeyArrowUp:
				// up
				i, j := b.y, b.x
				if i+1 >= b.size {
					b.draw()
					continue
				}
				b.swap(i, j, i+1, j)
			case termbox.KeyArrowDown:
				// down
				i, j := b.y, b.x
				if i-1 < 0 {
					b.draw()
					continue
				}
				b.swap(i, j, i-1, j)
			}
		}
	}
}
