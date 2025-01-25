package imageutil

import (
	"image"
	"image/draw"
)

func Watermark(src image.Image, wmk image.Image, pt image.Point) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0, src.Bounds().Dx(), src.Bounds().Dy()))
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)

	height = wmk.Bounds().Dy() * 0.55
	// width = int(float64(height) / float64(wmk.Bounds().Dy()) * float64(wmk.Bounds().Dx()))
	// rect := image.Rect(0, 0, width, height)
	// temp := image.NewRGBA(rect)
	// sc := scdraw.BiLinear /* default interpolator */
	// sc.Scale(temp, rect, wmk, wmk.Bounds(), scdraw.Over, nil)

	temp := imageutil.Resize(wmk, 0, height)

	draw.Draw(dst, temp.Bounds().Add(dst.Bounds().Max.Sub(temp.Bounds().Max)), temp, temp.Bounds().Min.Add(pt), draw.Over)
	return dst
}
