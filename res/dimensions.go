package res

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/ui"
)

var (
	TileSize = math.Vector{
		X: 8,
		Y: 8,
	}
	GameSize = math.Vector{
		X: 400,
		Y: 200,
	}
	WindowSize = math.Vector{
		X: 1200,
		Y: 600,
	}

	EggAnimationFrameSize = math.Vector{
		X: 16,
		Y: 16,
	}
	EggSize = math.Vector{
		X: 16,
		Y: 16,
	}
	EggWeaponAnchorPoint = math.Vector{
		X: 19,
		Y: 9,
	}

	InteractionButtonSize = math.Vector{
		X: 11,
		Y: 12,
	}
	InteractionButtonSpace = 20.0

	PlayerStatusBarPosition = math.Vector{
		X: 54,
		Y: 14,
	}
	CharacterIconPosition = math.Vector{
		X: 12,
		Y: 9,
	}
	HealthBarPosition = math.Vector{
		X: 28,
		Y: 5,
	}
	WeaponStatusBarPosition = math.Vector{
		X: 28,
		Y: 13,
	}
	HeartSize = math.Vector{
		X: 6,
		Y: 6,
	}
	HeartMargin = 1.0

	HUDPadding                      = ui.SamePadding(5)
	PlayerBarPadding                = ui.HorizontalVerticalPadding(4, 3)
	HealthBarGap                    = 1.0
	SpaceBetweenHealthAndWeaponBars = 2.0
	WeaponStatusBarGap              = 2.0
	AmmoTextureSize                 = math.Vector{
		X: 10,
		Y: 6,
	}
)
