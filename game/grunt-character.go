package game

import (
	"github.com/nosimplegames/eggventure/engine"
	"github.com/nosimplegames/eggventure/res"
)

type GruntCharacter struct {
	engine.LivingCharacter
}

type GruntCharacterFactory struct {
}

func (factory GruntCharacterFactory) Init(grunt *GruntCharacter) {
	textures := res.GetTextures()
	animations := res.GetAnimations()
	engine.LivingCharacterFactory{
		Size: res.GruntClosedSize,

		MaxHealth: 4,
		Health:    4,

		DamageMaskSize:      res.GruntClosedSize,
		LowDamageTexture:    textures.GruntLowDamage,
		MediumDamageTexture: textures.GruntMediumDamage,
		HighDamageTexture:   textures.GruntHighDamage,

		CharacterFactory: engine.CharacterFactory{
			IdleTexture:      textures.GruntClosed,
			FallingTexture:   textures.GruntClosed,
			WalkingAnimation: animations.GruntWalkingAnimation,
		},
	}.Init(&grunt.LivingCharacter)
	grunt.CollisionMask = "enemy"

	grunt.SetSize(res.GruntClosedSize)
	grunt.SetOrigin(res.GruntClosedSize.By(0.5))
}
