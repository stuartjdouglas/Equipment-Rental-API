package utils

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestUtilsQrCodeHandler(t *testing.T) {
	g := Goblin(t)

	//	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Extracting extension", func() {

		g.It("qr code should retain the given dimentions given", func() {
			image := GenerateQRCode("hello", 500, 500);
			g.Assert(image.Bounds().Size().X == 500).IsTrue()
			g.Assert(image.Bounds().Size().Y == 500).IsTrue()
		})
		g.It("qr code should default to 300 x 300", func() {
			image := GenerateQRCode("hello", 0, 0);
			g.Assert(image.Bounds().Size().X == 300).IsTrue()
			g.Assert(image.Bounds().Size().Y == 300).IsTrue()
		})
	})
}

