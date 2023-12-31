package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
)

type Player struct {
	EggCharacter
}

func (player *Player) HandleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		player.SetState(engine.Jumping)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		player.Interact()
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		player.Shoot()
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.SetState(engine.Walking)
		player.SetMovingDirection(engine.MovingDirectionRight)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.SetState(engine.Walking)
		player.SetMovingDirection(engine.MovingDirectionLeft)
	} else {
		player.SetState(engine.Idle)
	}
}

type PlayerFactory struct {
}

func (factory PlayerFactory) Create() *Player {
	player := &Player{}
	EggCharacterFactory{}.Init(&player.EggCharacter)

	player.CollisionMask = "player"
	player.SetPosition(res.GameSize.By(0.5))

	return player
}
