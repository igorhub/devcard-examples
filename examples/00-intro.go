package examples

import (
	"github.com/igorhub/devcard"
)

func DevcardIntro(dc *devcard.Devcard) {
	dc.SetTitle("Getting started")

	dc.Md("One good way to get started with devcards is to clone the `devcard-examples` repo and to tinker with it. ",
		"This guide describes the process step by step.")

	dc.Md("First, clone the repo:")
	sh(dc, "git clone https://github.com/igorhub/devcard-examples")

	dc.Md("Second, install the devcards web app:")
	sh(dc, "go install github.com/igorhub/devcard/cmd/devcards@latest")

	dc.Md("Now, cd into `devcard-examples` and start the app:")
	sh(dc, `cd devcard-examples
devcards`)

	dc.Md("Put your browser alongside your editor, and go to http://127.0.0.1:50051. ",
		"Click [devcard-examples](/dc/devcard-examples) and inspect the example cards one by one.")

	dc.Md("Note: on the first run, the app will suggest to create an initial config file. ",
		"While not required, having a config file will open up some possibilities, ",
		"such as remembering your projects or changing the color scheme.")
}

// sh appends a cell with shell code to the end of the devcard.
func sh(dc *devcard.Devcard, src string) {
	dc.Mono(devcard.WithHighlighting("sh"), src)
}
