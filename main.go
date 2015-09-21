package main

import (
	"github.com/vinzBad/grid"
)

type DwarfScene struct {
}

func (s *DwarfScene) Update(dt float32) {

}

func (s *DwarfScene) Draw(dt float32) {

}

func main() {
	scene := DwarfScene{}

	game := &grid.Game{
		Width:      640,
		Height:     480,
		Fullscreen: false,
		Resizable:  false,
		Title:      "Dwarf",
		Scene:      &scene,
	}

	grid.Run(game)
}
