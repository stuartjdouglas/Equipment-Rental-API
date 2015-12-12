package models

import (
	"testing"
	. "github.com/franela/goblin"

	"github.com/remony/Equipment-Rental-API/core/config/database"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/config"
)
const CONF_FILE = "./../../config.json"

func TestAuthModel(t *testing.T) {
	g := Goblin(t)

	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Sessions", func() {
//		TODO get token from a login
		g.It("Should be valid", func() {
			g.Assert(IsSessionValid(router, "a17a095f-3db6-4635-bf5c-071bc26089e6")).IsTrue()
		})

		g.It("Should be invalid", func() {
			g.Assert(IsSessionValid(router, "master-code")).IsFalse()
		})
	})
}

