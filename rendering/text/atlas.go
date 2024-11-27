package text

import (
	"image"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Atlas struct {
	Face  font.Face
	Image *image.Alpha
}

const numCols = 16

func NewAtlas(face font.Face, alphabet Alphabet) *Atlas {
	type gylphInfo struct {
		pos     image.Point
		dims    image.Point
		bounds  fixed.Rectangle26_6
		advance fixed.Int26_6
	}

	gylphTable := make(map[rune]gylphInfo)
	currPos := image.Point{0, 0}
	nextLine := 0
	maxWidth := 0
	col := 0
	for _, r := range alphabet.Runes() {
		bounds, advance, ok := face.GlyphBounds(r)
		if !ok {
			println("Rune not in face:", r)
		}
		boundSize := bounds.Max.Add(bounds.Min.Mul(fixed.I(-1)))
		dims := image.Point{
			X: boundSize.X.Ceil(),
			Y: boundSize.Y.Ceil(),
		}
		gylphTable[r] = gylphInfo{
			pos:     currPos,
			dims:    dims,
			bounds:  bounds,
			advance: advance,
		}
		nextLine = max(nextLine, currPos.Y+dims.Y)
		maxWidth = max(maxWidth, currPos.X+dims.X)
		currPos.X += dims.X + 4
		col++
		if col >= numCols {
			col = 0
			currPos.X = 0
			currPos.Y = nextLine + 4
		}
	}

	atlas := image.NewAlpha(image.Rectangle{Max: image.Point{X: maxWidth, Y: nextLine + 3}})
	inv := &FlippedImage{Image: atlas}
	white := image.White

	m1 := fixed.I(-1)
	for r, info := range gylphTable {
		atlasPos := image.Rectangle{
			Min: info.pos,
			Max: info.dims.Add(info.pos),
		}
		_, m, mp, _, _ := face.Glyph(info.bounds.Min.Mul(m1), r)
		// if r == 'a' {
		// 	println(atlasPos.Max)
		// }
		draw.DrawMask(inv, atlasPos, white, image.Point{}, m, mp, draw.Over)
	}

	return &Atlas{
		Face:  face,
		Image: atlas,
	}
}
