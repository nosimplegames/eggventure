package res

import (
	nsanimations "github.com/nosimplegames/ns-framework/animations"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type Animations struct {
	EggWalkingAnimation core.IAnimation

	ActionButtonAnimation core.IAnimation
}

var animations *Animations = nil

func GetAnimations() *Animations {
	needToInitanimations := animations == nil

	if needToInitanimations {
		textures := GetTextures()

		animations = &Animations{}
		animations.EggWalkingAnimation = nsanimations.SpriteAnimationFactory{
			Texture:       textures.WalkingEggAnimation,
			FrameDuration: 0.1,
			FrameSize: math.Vector{
				X: 16,
				Y: 16,
			},
			LoopCount: nsanimations.AnimationInfiniteLoop,
		}.Create(nil)

		animations.ActionButtonAnimation = nsanimations.SpriteAnimationFactory{
			Texture:       textures.ActionButtonAnimation,
			FrameDuration: 0.3,
			FrameSize: math.Vector{
				X: 11,
				Y: 12,
			},
			LoopCount: nsanimations.AnimationInfiniteLoop,
		}.Create(nil)
	}

	return animations
}
