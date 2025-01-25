package imageutil

import (
	"image"
	"image/draw"
	"golang.org/x/image/draw"
)

func Watermark(src image.Image, wmk image.Image, pt image.Point) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0,
		src.Bounds().Dx(), src.Bounds().Dy()))
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)

	height = wmk.Bounds().Dy() * 0.56
	width = int(float64(height) / float64(wmk.Bounds().Dy()) * float64(wmk.Bounds().Dx()))
	rect := image.Rect(0, 0, width, height)
	sc := draw.BiLinear /* default interpolator */
	sc.Scale(dst, rect.Bounds().Add(dst.Bounds().Max.Sub(rect.Bounds().Max)), wmk, wmk.Bounds(), draw.Over, nil)
	return dst
}
