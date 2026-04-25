package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames []*ebiten.Image // Animation frames
	FPS    float64         // Animation playback speed (Frames Per Second).
}

func NewAnimation(img *ebiten.Image, rects []image.Rectangle) *Animation {
	// TODO: implement NewAnimation - build Frames by extracting SubImages
	return &Animation{}
}

func (anim *Animation) Update() {
	// TODO: implement Update
}

func (anim *Animation) Draw(screen *ebiten.Image, opts DrawOptions) {
	// TODO: implement Draw
}
