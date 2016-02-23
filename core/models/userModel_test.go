package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
)

//const ConfigFile = "./../../config.json"

func TestUserModel(t *testing.T) {
	g := Goblin(t)

	testUser := Register{
		Username: "lemontest",
		Password: "testpassword",
		Email: "test@email.com",
	}
	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}
	PerformRegister(api, testUser, true)

	g.Describe("User", func() {
		//g.It("exists should return true", func() {
		//	g.Assert(CheckIfUserExists(api, testUser.Username)).IsTrue()
		//})
		g.It("exits should return false", func() {
			g.Assert(CheckIfUserExists(api, testUser.Username + "aaaaaa")).IsFalse()
		})
	})
}

