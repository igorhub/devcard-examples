package examples

import (
	"image"

	"github.com/igorhub/devcard"
)

func DevcardImageCell(c *devcard.Devcard) {
	c.SetTitle("3. Image cell")

	c.Md("Image cells allow us to use images.")

	code(c, `c.Image("img/wallhaven-0qy15n.jpg")`)

	c.Image("img/wallhaven-0qy15n.jpg")

	c.Md("An image can be defined in two ways: by using a path to an image file, or by using an `image.Image` instance.")
	c.Append("Similarly to annotated value cells, an annotation may be attached to an image:")

	code(c, `var img image.Image = stripes()
c.Image("Path to an image", "img/red-pandas.jpg", "image.Image instance", img)`)

	var img image.Image = stripes()
	c.Image("Path to an image", "img/red-pandas.jpg", "image.Image instance", img)
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
