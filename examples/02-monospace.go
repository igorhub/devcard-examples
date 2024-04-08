package examples

import (
	"image"

	"github.com/igorhub/devcard"
)

func DevcardMonospacedCells(c *devcard.Devcard) {
	c.SetTitle("2. Monospaced cells")

	c.Md("There are 4 types of monospace cells: `MonospaceCell`, `ValueCell`, `AnnotatedValueCell`, and `SourceCell`.")

	// MonospaceCell
	c.Md("### Monospace cell")
	c.Md("With `c.Mono` we create a basic monospace cell. ")
	c.Mono(devcard.WithHighlighting("go"), `c.Mono("I am a monospace Cell.")`)
	c.Mono("I am a monospace Cell.")

	c.Md("Monospace cells support syntax highlighting with `devcard.WithHighlighting` option:")

	code(c, "c.Mono(devcard.WithHighlighting(\"clojure\"), \"(defn lookup [kw]\\n  (get @*registry kw))\")")

	c.Mono(devcard.WithHighlighting("clojure"), "(defn lookup [kw]\n  (get @*registry kw))")

	c.Md("The behavior of `c.Append` for a monospace cell is slightly different from a markdown cell: rather than adding text to the end of the line, it appends a new line.")
	c.Append("I find this behavior both intuitive and the most common use case.")

	code(c, `c.Mono()
for i := range 7 {
	c.Append("Processing item #", i+1, "...")
}
c.Append("Done.")`)

	c.Mono()
	for i := range 7 {
		c.Append("Processing item #", i+1, "...")
	}
	c.Append("Done.")

	// ValueCell
	c.Md("### Value cell")
	c.Md("Value cells contain pretty-printed Go values:")

	code(c, "rect := image.Rectangle{}\nc.Val(v)")

	rect := image.Rectangle{}
	c.Val(rect)

	// AnnotatedValueCell
	c.Md("### Annotated value cell")
	c.Md("Annotated value cells are similar, but each value has a comment attached to it:")

	code(c, `c.Aval("Small", image.Rect(5, 5, 10, 10), "Large", image.Rect(10, 10, 60, 60))`)
	c.Aval("Small", image.Rect(5, 5, 10, 10), "Large", image.Rect(10, 10, 60, 60))

	// SourceCell
	c.Md("### Source cell")

	c.Md("Source cell contain source code of a function declared in our module:")

	code(c, `c.Source("examples.code")`)
	c.Source("examples.code")

	c.Md("The name of the function should always be prefixed by the package it's declared in.")
	c.Md("Keep in mind that the function must be declared in the same module as our devcard.")
	c.Append("If you try to reach outside of it, you'll get an error:")

	code(c, `c.Source("image.Rect")`)
	c.Source("image.Rect")
}
