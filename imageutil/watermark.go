package imageutil

import (
	"image"
	"image/draw"
	// dr "golang.org/x/image/draw"
)

func Watermark(src image.Image, wmk image.Image, pt image.Point) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0, src.Bounds().Dx(), src.Bounds().Dy()))
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)

	// height := int(float64(src.Bounds().Dy()/9))
	// width := int(float64(height) / float64(wmk.Bounds().Dy()) * float64(wmk.Bounds().Dx()))
	// rect := image.Rect(src.Bounds().Dx() - width, src.Bounds().Dy() - height, width, height)
	// dr.BiLinear.Scale(dst, rect, wmk, wmk.Bounds(), dr.Over, nil)
	// dr.BiLinear.Scale(dst, rect.Add(src.Bounds().Max.Sub(rect.Size())), wmk, wmk.Bounds(), dr.Over, nil)
	// dr.BiLinear.Scale(dst, rect, wmk, wmk.Bounds(), dr.Over, nil)
	mark := Resize(wmk, 0, src.Bounds().Dy()/9)

	// temp := imageutil.Resize(wmk, 0, wmk.Bounds().Dy()/2)
	// draw.Draw(dst, temp.Bounds().Add(dst.Bounds().Max.Sub(temp.Bounds().Max)), temp, temp.Bounds().Min.Add(pt), draw.Over)
	draw.Draw(dst, mark.Bounds().Add(dst.Bounds().Max.Sub(mark.Bounds().Max)), mark, mark.Bounds().Min.Add(pt), draw.Over)
	return dst
}
