package main

import (
	"fmt"
	"os"
	"io"
	"time"
)

type piece struct {
	id int

	top int
	right int
	bottom int
	left int
}

func load_pieces(r io.Reader) [256]piece {
	var pieces [256]piece

	for i := 0; i < 256; i++ {
		fmt.Fscanf(r,
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

	print_pieces(load_pieces(fd))

	time.Sleep(10 * time.Second)
}
