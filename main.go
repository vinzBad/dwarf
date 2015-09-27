package main

import (
	"log"

	"github.com/vinzBad/grid"
)

func check(method string, err error) {
	if err != nil {
		log.Fatalln(method, "failed:", err)
	}
}

type DwarfScene struct {
	Alien grid.Sprite
}

func (s *DwarfScene) Setup() {
	var err error
	alienTx, err := grid.CreateTexture("alienGreen.png")
	check("grid.CreateTexture", err)
	s.Alien, err = grid.CreateSprite(alienTx)
	check("grid.CreateSprite", err)
}

func (s *DwarfScene) Update(dt float32) {

}

func (s *DwarfScene) Draw(dt float32) {
	s.Alien.Draw()
}

func main() {
	game := &grid.Game{
		Width:      640,
		Height:     480,
		Fullscreen: false,
		Resizable:  false,
		Title:      "Dwarf",
		Scene:      &DwarfScene{},
	}

	check("grid.Run", grid.Run(game))
}
