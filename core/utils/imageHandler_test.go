package utils

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestUtilsImageHandler(t *testing.T) {
	g := Goblin(t)

	//	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Extracting extension", func() {
		g.It("image.png should return .png", func() {
			g.Assert(getFileExt("image.png") == "png").IsTrue()
		})
		g.It("image.jpg.png should return .png", func() {
			g.Assert(getFileExt("image.jpg.png") == "png").IsTrue()
		})
		g.It("image/image.png should return .png", func() {
			g.Assert(getFileExt("image.jpg.png") == "png").IsTrue()
		})
		g.It("image/image.jpg.png should return .png", func() {
			g.Assert(getFileExt("image.jpg.png") == "png").IsTrue()
		})
		g.It("image.jpg should return .jpg", func() {
			g.Assert(getFileExt("image.jpg") == "jpg").IsTrue()
		})
		g.It("image.jpg.jpg should return .jpg", func() {
			g.Assert(getFileExt("image.jpg.jpg") == "jpg").IsTrue()
		})
		g.It("image/image.jpg should return .jpg", func() {
			g.Assert(getFileExt("image.jpg.jpg") == "jpg").IsTrue()
		})
		g.It("image/image.jpg.jpg should return .jpg", func() {
			g.Assert(getFileExt("image.jpg.jpg") == "jpg").IsTrue()
		})
	})

	g.Describe("Extracting filename", func() {
		g.It("image.png should return image", func() {
			g.Assert(getFilename("image.png") == "image").IsTrue()
		})
		g.It("image/image.png should return image", func() {
			g.Assert(getFilename("image/image.png") == "image").IsTrue()
		})
		g.It("image/image/image.png should return image", func() {
			g.Assert(getFilename("image/image/image.png") == "image").IsTrue()
		})
		g.It("image/image/image/image.png should return image", func() {
			g.Assert(getFilename("image/image/image/image.png") == "image").IsTrue()
		})
	})

	g.Describe("Extracting full filename", func() {
		g.It("image.png should return image.png", func() {
			g.Assert(getFullFilename("image.png") == "image.png").IsTrue()
		})
		g.It("image/image.png should return image.png", func() {
			g.Assert(getFullFilename("image/image.png") == "image.png").IsTrue()
		})
		g.It("image/image/image.png should return image.png", func() {
			g.Assert(getFullFilename("image/image/image.png") == "image.png").IsTrue()
		})
		g.It("image/image/image/image.png should return image.png", func() {
			g.Assert(getFullFilename("image/image/image/image.png") == "image.png").IsTrue()
		})
	})
}

