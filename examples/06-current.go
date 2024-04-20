package examples

import (
	"github.com/igorhub/devcard"
)

func DevcardCurrent(dc *devcard.Devcard) {
	dc.SetTitle("Current devcard")

	dc.Md("`devcard` package provides `Current` function, which returns the devcard currently being filled with cells.")

	dc.Md("It point to the same devcard as the entry function parameter `dc`, as you can see:")

	code(dc, `dc.Ann("dc == devcard.Current()", dc == devcard.Current())`)
	dc.Ann("dc == devcard.Current()", dc == devcard.Current())

	dc.Md("The reason it's exposed is to make it available to any function in our program, no matter how deeply nested. ",
		"This can be useful for occasional printf-style debugging or for some quick experiments.")

	todo()
}

func todo() {
}
