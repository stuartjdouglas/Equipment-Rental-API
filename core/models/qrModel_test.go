package models

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestqrModel(t *testing.T) {
	g := Goblin(t)

	//api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	g.Describe("qr image", func() {
		image := GenerateQR("hello", 500, 500)
		g.It("should retain dimentions", func() {
			g.Assert(image.Bounds().Size().X == 500).IsTrue()
			g.Assert(image.Bounds().Size().Y == 500).IsTrue()
		})
	})

}

