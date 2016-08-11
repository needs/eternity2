package main

import "fmt"

type piece struct {
	id int

	top int
	right int
	bottom int
	left int

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

func load_pieces(pieces []piece) {
	for i := 0; i < len(pieces); i++ {
		fmt.Scanf(
			"%d,%d,%d,%d,%d", &pieces[i].id,
			&pieces[i].top, &pieces[i].right,
			&pieces[i].bottom, &pieces[i].left)

		pieces[i].cell = nil
	}
}

func init_board(board *board) {
	load_pieces(board.pieces[0:256])

	// Link cell neightbors
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

func does_fit(piece *piece, cell *cell) bool {
	if cell.top != nil && cell.top.piece != nil && cell.top.piece.bottom != piece.top {
		return false
	} else if cell.bottom != nil && cell.bottom.piece != nil && cell.bottom.piece.top != piece.bottom {
		return false
	} else if cell.left != nil && cell.left.piece != nil && cell.left.piece.right != piece.left {
		return false
	} else if cell.right != nil && cell.right.piece != nil && cell.right.piece.left != piece.right {
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
			if does_fit(&board.pieces[i], cell) {
				board.pieces[i].cell = cell
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
	var board board

	init_board(&board)

	if backtrack(&board, &board.cells[0][0]) {
		print_board(&board)
	} else {
		fmt.Printf("No solutions found\n")
	}
}
