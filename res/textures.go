package res

import (
	"github.com/nosimplegames/ns-framework/assets"
	"github.com/nosimplegames/ns-framework/render"
)

type Textures struct {
	Tileset render.Texture

	Egg        render.Texture
	FallingEgg render.Texture
	JumpingEgg render.Texture

	WalkingEggAnimation  render.Texture
	StopFallingAnimation render.Texture
}

var textures *Textures = nil

func GetTextures() *Textures {
	needToInitTextures := textures == nil

	if needToInitTextures {
		textures = &Textures{}
		textures.Tileset = assets.LoadTexture(tileset)
		textures.Egg = assets.LoadTexture(egg)
		textures.FallingEgg = assets.LoadTexture(fallingEgg)
		textures.JumpingEgg = assets.LoadTexture(jumpingEgg)

		textures.WalkingEggAnimation = assets.LoadTexture(walkingEggAnimation)
		textures.StopFallingAnimation = assets.LoadTexture(stopFallingEggAnimation)
	}

	return textures
}
