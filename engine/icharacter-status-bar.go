package engine

import "github.com/nosimplegames/ns-framework/render"

type ICharacterStatusBar interface {
	SetMaxHealth(int)
	SetHealth(int)
	SetAmmo(int)
	SetAmmoTexture(render.Texture)
}
