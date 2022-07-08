package main

import (
	"mafrans/gorogue/game"
	"mafrans/gorogue/gameobject"
	"mafrans/gorogue/input"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	input.OnKeyDown(tcell.KeyCtrlC, func() {
		game.Stop()
		os.Exit(0)
	})

	game.Instantiate(gameobject.NewPlayer())

	// Starting the game blocks the main thread
	game.Start()
}
