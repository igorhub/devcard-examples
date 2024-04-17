package examples

import (
	"image"

	"github.com/igorhub/devcard"
)

func DevcardImageCell(dc *devcard.Devcard) {
	dc.SetTitle("Image cell")

	dc.Md("Image cells allow you to use images.")

	code(dc, `dc.Image("img/wallhaven-0qy15n.jpg")`)

	dc.Image("img/wallhaven-0qy15n.jpg")

	dc.Md("An image can be defined in two ways: by using a path to an image file, or by using an `image.Image` instance.")
	dc.Append("Similarly to annotated value cells, an annotation may be attached to an image:")

	code(dc, `var img image.Image = stripes()
dc.Image("Path to an image", "img/red-pandas.jpg", "image.Image instance", img)`)

	var img image.Image = stripes()
	dc.Image("Path to an image", "img/red-pandas.jpg", "image.Image instance", img)
}

func stripes() image.Image {
	const (
		width  = 800
		height = 300
	)
	img := image.NewGray(image.Rect(0, 0, width, height))
	for x := range width {
		for y := range height {
			img.Pix[y*width+x] = uint8(x + y)
		}
	}
	return img
}
