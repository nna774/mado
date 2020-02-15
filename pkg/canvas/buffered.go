package canvas

import (
	"image"
	"image/color"
	"image/draw"
)

// BufferedCanvas is Canvas with buffer
type BufferedCanvas struct {
	c      Canvas
	buffer *image.NRGBA
}

// NewBufferedCanvas creates new BufferedCanvas
func NewBufferedCanvas(id string) BufferedCanvas {
	canvas := GetCanvas(id)
	rectangle := canvas.rectangle
	return BufferedCanvas{c: canvas, buffer: image.NewNRGBA(rectangle)}
}

// Update syncs real canvas and buffer
func (b BufferedCanvas) Update() {
	ctx := b.c.context()
	bound := b.c.rectangle.Max
	imageData := ctx.Call("createImageData", bound.X, bound.Y)
	data := imageData.Get("data")
	for i := 0; i < bound.Y; i++ {
		for j := 0; j < bound.X; j++ {
			pix := b.buffer.At(j, i)
			r, g, b, a := pix.RGBA()
			base := (i*bound.Y + j) * 4
			data.SetIndex(base+0, r>>8)
			data.SetIndex(base+1, g>>8)
			data.SetIndex(base+2, b>>8)
			data.SetIndex(base+3, a>>8)
		}
	}

	ctx.Call("putImageData", imageData, 0, 0)
}

// At is part of Image interface
func (b BufferedCanvas) At(x, y int) color.Color {
	return b.buffer.At(x, y)
}

// ColorModel is part of Image interface
func (b BufferedCanvas) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds is part of Image interface
func (b BufferedCanvas) Bounds() image.Rectangle {
	return b.c.rectangle
}

// Set is part of image/draw Image interface
func (b BufferedCanvas) Set(x, y int, color color.Color) {
	b.buffer.Set(x, y, color)
}

// FillTestRect is test pattern
func (b BufferedCanvas) FillTestRect() {
	bound := b.c.rectangle.Max
	c := color.NRGBA{G: 128, A: 255}

	for i := bound.X / 2; i < bound.X; i++ {
		for j := bound.Y / 2; j < bound.Y; j++ {
			b.Set(i, j, c)
		}
	}

}

// https://golang.org/doc/effective_go.html#blank_implements
var _ image.Image = new(BufferedCanvas)
var _ draw.Image = new(BufferedCanvas)
