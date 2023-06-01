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

	WalkingEggAnimation     render.Texture
	StopFallingEggAnimation render.Texture

	Magnum                   render.Texture
	MagnumAmmo               render.Texture
	MagnumItem               render.Texture
	MagnumFlash              render.Texture
	MagnumExplosionAnimation render.Texture

	ActionButtonAnimation render.Texture

	PlayerBar    render.Texture
	EggCharacter render.Texture
	FilledHeart  render.Texture
	EmptyHeart   render.Texture
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
		textures.StopFallingEggAnimation = assets.LoadTexture(stopFallingEggAnimation)

		textures.Magnum = assets.LoadTexture(magnum)
		textures.MagnumAmmo = assets.LoadTexture(magnumAmmo)
		textures.MagnumItem = assets.LoadTexture(magnumItem)
		textures.MagnumFlash = assets.LoadTexture(magnumFlash)
		textures.MagnumExplosionAnimation = assets.LoadTexture(magnumExplosionAnimation)

		textures.ActionButtonAnimation = assets.LoadTexture(actionButtonAnimation)

		textures.PlayerBar = assets.LoadTexture(playerBar)
		textures.EggCharacter = assets.LoadTexture(eggCharacter)
		textures.EmptyHeart = assets.LoadTexture(emptyHeart)
		textures.FilledHeart = assets.LoadTexture(filledHeart)
	}

	return textures
}
