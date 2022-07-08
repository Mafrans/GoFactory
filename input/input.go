package input

import (
	"github.com/gdamore/tcell/v2"
)

type keyListener struct {
	key tcell.Key
	cb  func()
}

type runeListener struct {
	rune rune
	cb   func()
}

type clickListener struct {
	button tcell.ButtonMask
	cb     func()
}

var keyListeners []keyListener = make([]keyListener, 0)
var runeListeners []runeListener = make([]runeListener, 0)
var clickListeners []clickListener = make([]clickListener, 0)

var MouseX = 0
var MouseY = 0

func OnKeyDown(key tcell.Key, cb func()) {
	keyListeners = append(keyListeners, keyListener{
		key, cb,
	})
}

func OnRuneDown(rune rune, cb func()) {
	runeListeners = append(runeListeners, runeListener{
		rune, cb,
	})
}

func OnClick(button tcell.ButtonMask, cb func()) {
	clickListeners = append(clickListeners, clickListener{
		button, cb,
	})
}

func HandleKeyEvent(e *tcell.EventKey) {
	for _, listener := range keyListeners {
		if e.Key() == listener.key {
			listener.cb()
		}
	}

	for _, listener := range runeListeners {
		if e.Rune() == listener.rune {
			listener.cb()
		}
	}
}

func HandleMouseEvent(e *tcell.EventMouse) {
	for _, listener := range clickListeners {
		if e.Buttons() == listener.button {
			listener.cb()
		}
	}

	MouseX, MouseY = e.Position()
}
