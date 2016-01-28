package models

import (
	"testing"
	. "github.com/franela/goblin"
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

	//g.Describe("Tags extracted from JSON", func() {
	//	var json = "tag, eletronic, stuff, anothers";
	//	g.It("should be returned in struct", func() {
	//		log.Println(parseJSArrayTags(json))
	//	})
	//})

	g.Describe("Count", func() {
		g.It("below 6 should return 6", func() {
			g.Assert(parseCount(0)).Equal(6)
			g.Assert(parseCount(1)).Equal(6)
			g.Assert(parseCount(2)).Equal(6)
			g.Assert(parseCount(3)).Equal(6)
			g.Assert(parseCount(4)).Equal(6)
			g.Assert(parseCount(5)).Equal(6)
		})
		g.It("value 6 should be 6", func() {
			g.Assert(parseCount(6)).Equal(6)
		})
		g.It("above 6 should return there own value", func() {
			g.Assert(parseCount(7) == 6).IsFalse()
			g.Assert(parseCount(8) == 6).IsFalse()
			g.Assert(parseCount(9) == 6).IsFalse()
		})
	})

	g.Describe("String to int", func() {
		g.It("string 5 should be int 5", func() {
			g.Assert(parseStringToInt("5")).Equal(5)
		})
	})

}

