package main

import (
	"fmt"
	"os"
	"io"
	"time"
)

type sides struct {
	top int
	right int
	bottom int
	left int
}

type piece struct {
	id int

	// For every rotation
	rots [4]sides

	side *sides
	cell *cell
}

type cell struct {
	piece *piece
	next *cell
	top, right, bottom, left *cell
}

type board struct {
	pieces [256]piece
	cells [16][16]cell
}

func rotate(side sides, rot int) sides {
	ret := side

	for i := 0; i < rot; i++ {
		save := ret.top
		ret.top = ret.left
		ret.left = ret.bottom
		ret.bottom = ret.right
		ret.right = save
	}

	return ret
}

func load_pieces(r io.Reader, pieces []piece) {
	for i := 0; i < len(pieces); i++ {
		pieces[i].side = &pieces[i].rots[0]

		fmt.Fscanf(r,
			"%d,%d,%d,%d,%d", &pieces[i].id,
			&pieces[i].side.top, &pieces[i].side.right,
			&pieces[i].side.bottom, &pieces[i].side.left)

		for rot := 0; rot < 4; rot++ {
			pieces[i].rots[rot] = rotate(pieces[i].rots[0], rot)
		}

		pieces[i].cell = nil
	}
}

func init_board(r io.Reader, board *board) {
	load_pieces(r, board.pieces[0:256])

	// Link cell to neightbors
	for i := 0; i < len(board.cells); i++ {
		for j := 0; j < len(board.cells[i]); j++ {
			board.cells[i][j].piece = nil

			if i == 0 {
				board.cells[i][j].top = nil
			} else {
				board.cells[i][j].top = &board.cells[i - 1][j]
			}

			if i == len(board.cells) - 1 {
				board.cells[i][j].bottom = nil
			} else {
				board.cells[i][j].bottom = &board.cells[i + 1][j]
			}

			if j == 0 {
				board.cells[i][j].left = nil
			} else {
				board.cells[i][j].left = &board.cells[i][j - 1]
			}

			if j == len(board.cells[i]) - 1 {
				board.cells[i][j].right = nil
			} else {
				board.cells[i][j].right = &board.cells[i][j + 1]
			}
		}
	}

	// Link cells row by row
	for i := 0; i < len(board.cells); i++ {
		for j := 0; j < len(board.cells[i]); j++ {
			if j == len(board.cells[i]) - 1 && i == len(board.cells) - 1 {
				board.cells[i][j].next = nil
			} else if j == len(board.cells[i]) - 1 && i < len(board.cells) - 1 {
				board.cells[i][j].next = &board.cells[i + 1][0]
			} else {
				board.cells[i][j].next = &board.cells[i][j + 1]
			}
		}
	}
}

func does_fit(side sides, cell *cell) bool {
	if cell.top != nil && cell.top.piece != nil && cell.top.piece.side.bottom != side.top {
		return false
	} else if cell.bottom != nil && cell.bottom.piece != nil && cell.bottom.piece.side.top != side.bottom {
		return false
	} else if cell.left != nil && cell.left.piece != nil && cell.left.piece.side.right != side.left {
		return false
	} else if cell.right != nil && cell.right.piece != nil && cell.right.piece.side.left != side.right {
		return false
	}

	return true
}

var maxdepth = 0
var depth = 0

func backtrack(board *board, cell *cell) bool {
	// Temporary hack to display current backtracking depth
	if depth > maxdepth {
		fmt.Printf("\rDepth = %d", depth)
		maxdepth = depth
	}

	if cell == nil {
		return true
	} else if cell.piece != nil {
		return backtrack(board, cell.next)
	}

	for i := 0; i < len(board.pieces); i++ {
		if board.pieces[i].cell == nil {
			for rot := 0; rot < 4; rot++ {
				if does_fit(board.pieces[i].rots[rot], cell) {
					board.pieces[i].cell = cell
					board.pieces[i].side = &board.pieces[i].rots[rot]
					cell.piece = &board.pieces[i]

					depth++
					if backtrack(board, cell.next) {
						return true
					}
					depth--

					board.pieces[i].cell = nil
					cell.piece = nil
				}
			}
		}
	}

	return false
}

func print_board(board *board) {
	for i := 0; i < len(board.cells); i++ {
		for j := 0; j < len(board.cells[i]); j++ {
			if board.cells[i][j].piece != nil {
				fmt.Printf("%3d ", board.cells[i][j].piece.id)
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pieces file>\n", os.Args[0])
		os.Exit(1)
	}

	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err);
		os.Exit(2)
	}

	go StartViewer()

	var board board

	init_board(fd, &board)

	if backtrack(&board, &board.cells[0][0]) {
		print_board(&board)
	} else {
		fmt.Printf("No solutions found\n")
	}

	time.Sleep(10 * time.Second)
}
