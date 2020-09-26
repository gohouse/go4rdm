package fmcolor

import "image/color"

func GetColorWithAlpha(cor *color.RGBA, a uint8) *color.RGBA {
	return &color.RGBA{R: cor.R, G: cor.G, B: cor.B, A: a}
}
