package res

import (
	_ "embed"
)

var (
	//go:embed tileset.png
	tileset []byte

	//go:embed egg.png
	egg []byte
	//go:embed falling-egg.png
	fallingEgg []byte
	//go:embed jumping-egg.png
	jumpingEgg []byte
	//go:embed walking-egg-animation.png
	walkingEggAnimation []byte
	//go:embed stop-falling-egg-animation.png
	stopFallingEggAnimation []byte

	//go:embed grunt-closed.png
	gruntClosed []byte
	//go:embed grunt-low-damage.png
	gruntLowDamage []byte
	//go:embed grunt-medium-damage.png
	gruntMediumDamage []byte
	//go:embed grunt-high-damage.png
	gruntHighDamage []byte
	//go:embed grunt-walking-animation.png
	gruntWalkingAnimation []byte

	//go:embed magnum-item.png
	magnumItem []byte
	//go:embed magnum.png
	magnum []byte
	//go:embed magnum-ammo.png
	magnumAmmo []byte
	//go:embed magnum-flash.png
	magnumFlash []byte
	//go:embed magnum-explosion-animation.png
	magnumExplosionAnimation []byte

	//go:embed action-button-animation.png
	actionButtonAnimation []byte

	//go:embed player-bar.png
	playerBar []byte
	//go:embed egg-character.png
	eggCharacter []byte
	//go:embed filled-heart.png
	filledHeart []byte
	//go:embed empty-heart.png
	emptyHeart []byte
)
