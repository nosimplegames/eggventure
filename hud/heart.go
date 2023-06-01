package hud

import (
	"github.com/nosimplegames/eggventure/res"
	"github.com/nosimplegames/ns-framework/entities"
	"github.com/nosimplegames/ns-framework/render"
)

type Heart struct {
	entities.Sprite
}

func (heart *Heart) Empty() {
	heart.SetTexture(res.GetTextures().EmptyHeart)
}

func (heart *Heart) Fill() {
	heart.SetTexture(res.GetTextures().FilledHeart)
}

type HeartFactory struct {
	IsHeartEmpty bool
}

func (factory HeartFactory) Create() *Heart {
	texture := factory.getTexture()
	heart := &Heart{}
	entities.SpriteFactory{
		Texture: texture,
	}.Init(&heart.Sprite)

	return heart
}

func (factory HeartFactory) getTexture() render.Texture {
	textures := res.GetTextures()
	texture := textures.FilledHeart

	mustUseEmptyHeartTexture := factory.IsHeartEmpty

	if mustUseEmptyHeartTexture {
		texture = textures.EmptyHeart
	}

	return texture
}
