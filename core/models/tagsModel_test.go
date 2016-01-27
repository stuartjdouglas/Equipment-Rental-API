package models

import (
	"testing"
	. "github.com/franela/goblin"
	"log"
)

//const ConfigFile = "./../../config.json"

func TestTagModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("Strings", func() {
		g.It("should have whitespace removed", func() {
			g.Assert(removeWhitespace(" Hello, World! ")).Equal("Hello, World!")
			g.Assert(removeWhitespace("        something,     ")).Equal("something,")
		})
	})

	g.Describe("Tags extracted from JSON", func() {
		var json = "tag, eletronic, stuff, anothers";
		g.It("should be returned in struct", func() {
			log.Println(parseJSArrayTags(json))
		})

	})
}

