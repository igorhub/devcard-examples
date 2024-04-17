package examples

import (
	"github.com/igorhub/devcard"
)

// code appends a cell with Go code to the end of the devcard.
func code(dc *devcard.Devcard, src string) {
	dc.Mono(devcard.WithHighlighting("go"), src)
}

func DevcardTextCells(dc *devcard.Devcard) {
	dc.SetTitle("Text cells")

	dc.Md(`A devcard is essentialy a list of cells. The most basic type of a cell
is a markdown cell.`)

	dc.Md("### Markdown cell")
	dc.Md("You can add a markdown cell to a devcard with the following code:")
	code(dc, `dc.Md("I am a *markdown* **cell**!")`)
	dc.Md("I am a *markdown* **cell**!")

	dc.Md("### Devcard.Md arguments")
	dc.Md("`dc.Md` accepts variable number of arguments.",
		" The arguments are converted to strings and concatenated:")
	code(dc, `dc.Md("This cell is made from ", 3, " pieces.")`)
	dc.Md("This cell is made from ", 3, " pieces.")

	dc.Md("### Append")
	dc.Md("You may use `dc.Append` to append one or more piece to the bottom cell of the devcard:")
	code(dc,
		`dc.Md("This cell")
dc.Append("is made from ", 4)
dc.Append("pieces.")`)
	dc.Md("This cell")
	dc.Append("is made from ", 4)
	dc.Append("pieces.")

	dc.Md("### HTML cell")
	dc.Md("When markdown doesn't cover your needs, you may drop down to an HTML cell:")
	code(dc, `dc.Html(`+"`"+`<div style="color: purple; font-size: 2rem;">*PURPLE TEXT*</div>`+"`"+`)`)
	dc.Html(`<div style="color: purple; font-size: 2rem;">*PURPLE TEXT*</div><br/>`)
}
