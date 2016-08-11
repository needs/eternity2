package main

import "fmt"

type piece struct {
	id int

	top int
	right int
	bottom int
	left int
}

func load_pieces() [256]piece {
	var pieces [256]piece

	for i := 0; i < 256; i++ {
		fmt.Scanf(
			"%d,%d,%d,%d,%d", &pieces[i].id,
			&pieces[i].top, &pieces[i].right,
			&pieces[i].bottom, &pieces[i].left)
	}

	return pieces
}

func print_pieces(pieces [256]piece) {
	for i := 0; i < 256; i++ {
		fmt.Printf(
			"%3d: %2d %2d %2d %2d\n", pieces[i].id,
			pieces[i].top, pieces[i].right,
			pieces[i].bottom, pieces[i].left)
	}
}

func main() {
	print_pieces(load_pieces())
}
