package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(fmt.Sprint("Error: Can't start viewer -", err));
	}
}

func StartViewer() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("Eternity II",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 800,
		sdl.WINDOW_SHOWN)

	check(err)
	defer window.Destroy()


	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	check(err)
	defer renderer.Destroy()

	renderer.Clear()

	running := true

	for running { }

	sdl.Quit()
}
