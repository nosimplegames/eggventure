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
)
