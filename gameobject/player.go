package gameobject

import (
	"mafrans/gorogue/input"
	"time"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/exp/slices"
)

type Player struct {
	MoveSpeed    float64
	X            int
	Y            int
	moveCooldown float64
	isMoving     bool
	movePath     [][2]int
}

func NewPlayer() *Player {
	return &Player{
		MoveSpeed: 4,
	}
}

func (player *Player) Start(id int) {
	input.OnClick(tcell.ButtonPrimary, func() {
		player.movePath = PathFind(
			[2]int{player.X, player.Y},
			[2]int{input.MouseX, input.MouseY},
		)

		player.isMoving = true
	})
}

func (player *Player) Update(time time.Duration, deltaTime time.Duration) {
	player.updateMove(deltaTime)
}

func (player *Player) Draw(screen *tcell.Screen) {
	for _, node := range player.movePath {
		(*screen).SetContent(node[0], node[1], 'x', nil, tcell.StyleDefault)
	}

	playerRune := '🧍'
	if player.isMoving {
		playerRune = '🏃'
	}
	(*screen).SetContent(player.X, player.Y, playerRune, nil, tcell.StyleDefault)
}

func (player *Player) Destroy() {

}

func (player *Player) updateMove(deltaTime time.Duration) {
	if player.moveCooldown <= 0 && player.isMoving {
		if len(player.movePath) == 0 {
			player.isMoving = false

			return
		}

		player.X = player.movePath[len(player.movePath)-1][0]
		player.Y = player.movePath[len(player.movePath)-1][1]
		player.movePath = slices.Delete(player.movePath, len(player.movePath)-1, len(player.movePath))

		player.moveCooldown = 0.5 / player.MoveSpeed
	}
	player.moveCooldown -= deltaTime.Seconds()
}
