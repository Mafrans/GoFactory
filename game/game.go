package game

import (
	"mafrans/gorogue/input"
	"time"

	"github.com/gdamore/tcell/v2"
)

var Screen tcell.Screen
var stopUpdateLoop bool
var stopDrawLoop bool
var stopEventLoop bool

func Start() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	Screen = s

	Screen.Init()
	Screen.Sync()
	Screen.Show()
	Screen.EnableMouse()

	stopDrawLoop = false
	stopUpdateLoop = false
	stopEventLoop = false
	go startUpdateLoop()
	go startEventLoop()
	startDrawLoop()
}

func Stop() {
	stopUpdateLoop = true
	stopDrawLoop = true
	stopEventLoop = true
	Screen.DisableMouse()
	Screen.Fini()
}

func startUpdateLoop() {
	startTime := time.Now()
	lastTick := time.Now()
	for !stopUpdateLoop {
		update(time.Since(startTime), time.Since(lastTick))
		lastTick = time.Now()
	}
}

func startDrawLoop() {
	for !stopDrawLoop {
		draw()
	}
}

func startEventLoop() {
	for !stopEventLoop {
		handleEvent(Screen.PollEvent())
	}
}

func handleEvent(e tcell.Event) {
	switch e := e.(type) {
	case *tcell.EventKey:
		input.HandleKeyEvent(e)
	case *tcell.EventMouse:
		input.HandleMouseEvent(e)
	}
}

func update(time time.Duration, deltaTime time.Duration) {
	for _, gameObject := range gameObjects {
		gameObject.Update(time, deltaTime)
	}
}

func draw() {
	Screen.Clear()
	for _, gameObject := range gameObjects {
		gameObject.Draw(&Screen)
	}
	Screen.Show()
}
