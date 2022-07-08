package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type GameObject interface {
	Start(id int)
	Update(time time.Duration, deltaTime time.Duration)
	Draw(screen *tcell.Screen)
	Destroy()
}

var gameObjects map[int]GameObject = make(map[int]GameObject)
var nextId = 0

func Instantiate(gameObject GameObject) {
	gameObject.Start(nextId)
	gameObjects[nextId] = gameObject
	nextId++
}

func Destroy(id int) {
	gameObjects[id].Destroy()
	delete(gameObjects, id)
}
