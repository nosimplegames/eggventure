package game

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
)

type EggCharacter struct {
	engine.InteractiveCharacter
}

type EggCharacterFactory struct {
}

func (factory EggCharacterFactory) Init(egg *EggCharacter) {
	textures := res.GetTextures()
	animations := res.GetAnimations()
	engine.InteractiveCharacterFactory{
		GunnerCharacterFactory: engine.GunnerCharacterFactory{
			LivingCharacterFactory: engine.LivingCharacterFactory{
				MaxLife: 3,
				Life:    3,

				CharacterFactory: engine.CharacterFactory{
					IdleTexture:      textures.Egg,
					FallingTexture:   textures.FallingEgg,
					JumpingTexture:   textures.JumpingEgg,
					WalkingAnimation: animations.EggWalkingAnimation,
				},
			},
			WeaponAnchorPoint: res.EggWeaponAnchorPoint,
		},
	}.Init(&egg.InteractiveCharacter)
	egg.CollisionMask = "character"

	egg.FallingTexture = textures.FallingEgg
	egg.SetSize(res.EggSize)
	egg.SetOriginCenter()
}
