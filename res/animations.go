package res

import (
	nsanimations "github.com/nosimplegames/ns-framework/animations"
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
)

type Animations struct {
	WalkingAnimation core.IAnimation
}

var animations *Animations = nil

func GetAnimations() *Animations {
	needToInitanimations := animations == nil

	if needToInitanimations {
		textures := GetTextures()

		animations = &Animations{}
		animations.WalkingAnimation = nsanimations.SpriteAnimationFactory{
			Texture:       textures.WalkingEggAnimation,
			FrameDuration: 0.1,
			FrameSize: math.Vector{
				X: 16,
				Y: 16,
			},
			LoopCount: nsanimations.AnimationInfiniteLoop,
		}.Create(nil)
	}

	return animations
}
