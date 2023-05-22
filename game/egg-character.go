package game

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
)

type EggCharacterFactory struct {
}

func (factory EggCharacterFactory) Create() *engine.Character {
	egg := &engine.Character{}
	factory.Init(egg)

	return egg
}

func (factory EggCharacterFactory) Init(egg *engine.Character) {
	textures := res.GetTextures()
	animations := res.GetAnimations()
	characterFactory := engine.CharacterFactory{
		IdleTexture:      textures.Egg,
		FallingTexture:   textures.FallingEgg,
		JumpingTexture:   textures.JumpingEgg,
		WalkingAnimation: animations.WalkingAnimation,
	}
	characterFactory.Init(egg)

	egg.FallingTexture = textures.FallingEgg
	egg.Size = res.EggSize
	egg.SetOriginCenter()
}
