package examples

import (
	"github.com/igorhub/devcard"
)

func DevcardAnatomy(dc *devcard.Devcard) {
	dc.SetTitle("Anatomy of a devcard")

	dc.Md("A devcard is produced by running a devcard-producing function. ",
		"Such a function must follow the following criteria:\n",
		"* Its name begins with `Devcard`.\n",
		"* It takes a single argument, which is a pointer to `devcard.Devcard`.\n",
		"* It returns nothing.\n")

	dc.Mono(devcard.WithHighlighting("go"),
		`func DevcardAnatomy(dc *devcard.Devcard) {
	dc.SetTitle("Anatomy of a devcard")

	// ...
}`)

	dc.Md("All devcard-producing functions defined in your code will be collected in the [devcards list](/dc/devcard-examples) in order of appearance.")

	dc.Md("### Devcard's content")
	dc.Md("A devcard is essentialy a list of cells. ",
		"Some cells are rendered as paragraphs of text, others as code blocks or images. ",
		"Here are the first three cells of this very devcard:")
	dc.Mono(devcard.WithHighlighting("go"), sample())

	dc.Md("The next four devcards describe various types of cells in detail.")
}

func sample() string {
	return "func DevcardAnatomy(dc *devcard.Devcard) {\n" +
		"	dc.SetTitle(\"Anatomy of a devcard\")\n" +
		"\n" +
		"	dc.Md(\"A devcard is produced by running a devcard-producing function. \",\n" +
		"	\"Such a function must follow the following criteria:\\n\",\n" +
		"	\"* Its name begins with `Devcard`.\\n\",\n" +
		"	\"* It takes a single argument, which is a pointer to `devcard.Devcard`.\\n\",\n" +
		"	\"* It returns nothing.\\n\")\n" +
		"	\n" +
		"	dc.Mono(devcard.WithHighlighting(\"go\"),\n" +
		"	`func DevcardAnatomy(dc *devcard.Devcard) {\n" +
		"		dc.SetTitle(\"Anatomy of a devcard\")\n" +
		"	\n" +
		"	// ...\n" +
		"}`)\n" +
		"\n" +
		"	dc.Md(\"All devcard-producing functions defined in your code will be collected in the [devcards list](/dc/devcard-examples) in order of appearance.\")\n"
}
