package examples

import "github.com/igorhub/devcard"

func DevcardFoobar(dc *devcard.Devcard) {
	dc.SetTitle("*scratch*")

	dc.Md("This devcard is for you to play with.")

	dc.Md("...")

	dc.Md("Note the icon next to the devcard's title. ",
		"Clicking it opens the devcard-producing function in your editor (defaults to vscode).")
}
