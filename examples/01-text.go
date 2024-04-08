package examples

import (
	"github.com/igorhub/devcard"
)

// code appends a cell with Go code to the end of the devcard.
func code(c *devcard.Devcard, src string) {
	c.Mono(devcard.WithHighlighting("go"), src)
}

func DevcardTextCells(c *devcard.Devcard) {
	c.SetTitle("1. Text cells")

	c.Md(`A devcard is essentialy a list of cells. The most basic type of a cell
is a markdown cell.`)

	c.Md("### Markdown cell")
	c.Md("We may add a markdown cell to our devcard with the following code:")
	code(c, `c.Md("I am a *markdown* **cell**!")`)
	c.Md("I am a *markdown* **cell**!")

	c.Md("### Devcard.Md arguments")
	c.Md("`c.Md` accepts variable number of arguments.",
		" The arguments are converted to strings and concatenated:")
	code(c, `c.Md("This cell is made from ", 3, " pieces.")`)
	c.Md("This cell is made from ", 3, " pieces.")

	c.Md("### Append")
	c.Md("We may use `c.Append` to append one or more piece to the bottom cell of our devcard:")
	code(c,
		`c.Md("This cell")
c.Append("is made from ", 4)
c.Append("pieces.")`)
	c.Md("This cell")
	c.Append("is made from ", 4)
	c.Append("pieces.")

	c.Md("### HTML cell")
	c.Md("When markdown doesn't cover our needs, we may drop down to an HTML cell:")
	code(c, `c.Html(`+"`"+`<div style="color: purple; font-size: 2rem;">*PURPLE TEXT*</div>`+"`"+`)`)
	c.Html(`<div style="color: purple; font-size: 2rem;">*PURPLE TEXT*</div><br/>`)
}
