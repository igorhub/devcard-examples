package examples

import (
	"image"

	"github.com/igorhub/devcard"
)

func DevcardMonospacedCells(dc *devcard.Devcard) {
	dc.SetTitle("Monospaced cells")

	dc.Md("There are 4 types of monospace cells: `MonospaceCell`, `ValueCell`, `AnnotatedValueCell`, and `SourceCell`.")

	// MonospaceCell
	dc.Md("### Monospace cell")
	dc.Md("With `dc.Mono`, you can create a basic monospace cell. ")
	dc.Mono(devcard.WithHighlighting("go"), `dc.Mono("I am a monospace Cell.")`)
	dc.Mono("I am a monospace Cell.")

	dc.Md("Monospace cells support syntax highlighting with `devcard.WithHighlighting` option:")

	code(dc, "dc.Mono(devcard.WithHighlighting(\"clojure\"), \"(defn lookup [kw]\\n  (get @*registry kw))\")")

	dc.Mono(devcard.WithHighlighting("clojure"), "(defn lookup [kw]\n  (get @*registry kw))")

	dc.Md("The behavior of `dc.Append` for a monospace cell is slightly different from a markdown cell: ",
		"rather than adding text to the end of the line, it appends a new line. ",
		"I find this behavior both intuitive and the most common use case.")

	code(dc, `dc.Mono()
for i := range 7 {
	dc.Append("Processing item #", i+1, "...")
}
dc.Append("Done.")`)

	dc.Mono()
	for i := range 7 {
		dc.Append("Processing item #", i+1, "...")
	}
	dc.Append("Done.")

	// ValueCell
	dc.Md("### Value cell")
	dc.Md("Value cells contain pretty-printed Go values:")

	code(dc, "rect := image.Rectangle{}\ndc.Val(v)")

	rect := image.Rectangle{}
	dc.Val(rect)

	// AnnotatedValueCell
	dc.Md("### Annotated value cell")
	dc.Md("Annotated value cells are similar, but each value has a comment attached to it:")

	code(dc, `dc.Ann("Small", image.Rect(5, 5, 10, 10), "Large", image.Rect(10, 10, 60, 60))`)
	dc.Ann("Small", image.Rect(5, 5, 10, 10), "Large", image.Rect(10, 10, 60, 60))

	// SourceCell
	dc.Md("### Source cell")

	dc.Md("Source cell contain source code of a function declared in the module:")

	code(dc, `dc.Source("examples.code")`)
	dc.Source("examples.code")

	dc.Md("The name of the function should always be prefixed by the package it's declared in.")

	dc.Md("Keep in mind that the function must be declared in the same module as your devcard. ",
		"If you try to reach outside of it, you'll get an error:")

	code(dc, `dc.Source("image.Rect")`)
	dc.Source("image.Rect")
}
