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
			g.Assert(IsSessionValid(router, "6cebef82-ac30-4492-b63f-f7aa0c249c20")).IsTrue()
		})

		g.It("Should be invalid", func() {
			g.Assert(IsSessionValid(router, "master-code")).IsFalse()
		})
	})
}

