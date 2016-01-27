package utils

import (
	"testing"
	. "github.com/franela/goblin"
)

const ConfigFile = "./../../config.json"

func TestUtilsEncryption(t *testing.T) {
	g := Goblin(t)

	//	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Encypting", func() {
		g.It("a password with SHA512 should be 128 characters long", func() {
			g.Assert(len(Sha512Me([]byte("Hello"))) == 128).IsTrue()
			g.Assert(len(Sha512Me([]byte("Hello, World!"))) == 128).IsTrue()
			g.Assert(len(Sha512Me([]byte("This is a password"))) == 128).IsTrue()
			g.Assert(len(Sha512Me([]byte("47 is the magic number"))) == 128).IsTrue()

		})

	})
}

