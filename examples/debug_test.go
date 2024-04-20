package examples

import (
	"testing"

	"github.com/igorhub/devcard"
)

func TestDebugDevcardIntro(t *testing.T) {
	devcard.Debug(DevcardIntro)
}

func TestDebugDevcardAnatomy(t *testing.T) {
	devcard.Debug(DevcardAnatomy)
}

func TestDebugDevcardTextCells(t *testing.T) {
	devcard.Debug(DevcardTextCells)
}

func TestDebugDevcardMonospacedCells(t *testing.T) {
	devcard.Debug(DevcardMonospacedCells)
}

func TestDebugDevcardImageCell(t *testing.T) {
	devcard.Debug(DevcardImageCell)
}

func TestDebugDevcardCustomCell(t *testing.T) {
	devcard.Debug(DevcardCustomCell)
}

func TestDebugDevcardCurrent(t *testing.T) {
	devcard.Debug(DevcardCurrent)
}

func TestDebugDevcardFoobar(t *testing.T) {
	devcard.Debug(DevcardFoobar)
}
