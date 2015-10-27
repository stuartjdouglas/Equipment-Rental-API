package models

import (
	"testing"
	. "github.com/franela/goblin"
	"../config"
	"../router"
	"../config/database"
)
const CONF_FILE = "./../../config.json"

func TestAuthModel(t *testing.T) {
	g := Goblin(t)

	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Sessions", func() {
//		TODO get token from a login
		g.It("Should be valid", func() {
			g.Assert(IsSessionValid(router, "f4fa5c51-0b58-4884-8f8d-ea832728ac8e")).IsTrue()
		})

		g.It("Should be invalid", func() {
			g.Assert(IsSessionValid(router, "master-code")).IsFalse()
		})
	})
}

