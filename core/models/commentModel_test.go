package models

import (
	"testing"
	. "github.com/franela/goblin"
	//"github.com/remony/Equipment-Rental-API/core/router"
	//"github.com/remony/Equipment-Rental-API/core/config"
	"gitlab.com/remon/lemon-swear-detector"
)

//const CONF_FILE = "./../../config.json"

func TestCommentModel(t *testing.T) {
	g := Goblin(t)
	//router := router.API{Context: config.Connection(config.LoadConfig(CONF_FILE, true).Production.DbUrl)}

	g.Describe("Text", func() {
		// TODO when database is complete change to check is true
		g.It("Should 0x0 should be detected as naughty", func() {
			g.Assert(checkIfNaughtyWord("0x0")).IsTrue()
		})

		g.It("empty should be detected as naughty", func() {
			g.Assert(checkIfNaughtyWord("")).IsTrue()
		})
	})

	g.Describe("library lemon swear detector", func() {
		g.It("Should detect word tit as a swear word", func() {
			g.Assert(lemon_swear_detector.CheckWord("tit")).IsTrue()
		})
		g.It("should detect word twat in sentence", func() {
			g.Assert((lemon_swear_detector.CheckSentence("x is a twat"))).IsTrue()
		})
	})
}
