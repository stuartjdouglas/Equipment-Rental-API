package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

const ConfigFile = "./../../config.json"

func TestAuthModel(t *testing.T) {
	g := Goblin(t)

	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	testUser := Register{
		Username: "lemontest",
		Password: "testpassword",
		Email: "test@email.com",
	}

	g.Describe("register", func() {
		g.It("should be successful", func() {
			g.Assert(PerformRegister(api, testUser))
		})
	})

	g.Describe("login", func() {
		g.It("should return true", func() {
			g.Assert(PerformLogin(api, testUser.Username, testUser.Password).Success).IsTrue()
		})

		g.It("should return username lemon", func() {
			g.Assert(PerformLogin(api, testUser.Username, testUser.Password).Username == testUser.Username).IsTrue()
		})

		g.It("should return a token", func() {
			g.Assert(len(PerformLogin(api, testUser.Username, testUser.Password).Token) > 2).IsTrue()
		})

		g.It("should not return a token with length 0", func() {
			g.Assert(len(PerformLogin(api, testUser.Username, testUser.Password).Token) != 0).IsTrue()
		})

		g.It("should return false with bad login", func() {
			g.Assert(PerformLogin(api, testUser.Username, "Password123").Success).IsFalse()
		})

		g.It("should return true is password is correct", func() {
			var digest = database.GetDigest(api, testUser.Username)
			g.Assert(authLogin(testUser.Password, digest)).IsTrue()
		})

		g.It("should return false is password is incorrect", func() {
			var digest = database.GetDigest(api, testUser.Username)
			g.Assert(authLogin("Password123", digest)).IsFalse()
		})
	})

	g.Describe("Checking valid username", func() {
		g.It("lemon should return true", func() {
			g.Assert(isValidEntry("lemon")).IsTrue()
		})
		g.It("of over 240 character should return false", func() {
			g.Assert(isValidEntry("lemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemon")).IsFalse()
		})
		g.It("containing a character should return false", func() {
			g.Assert(isValidEntry("$now")).IsFalse()
		})

	})

	g.Describe("password handling", func() {
		g.It("with length 5 should be false", func() {
			g.Assert(secureEntry("sdfjl")).IsFalse()
		})
		g.It("with 6 character should be accepted", func() {
			g.Assert(secureEntry("asdasd")).IsTrue()
		})
		g.It("with more than 6 characters should be true", func() {
			g.Assert(secureEntry("asdasdasd")).IsTrue()
		})
		g.It("should return false if spaces are detected", func() {
			g.Assert(secureEntry("asd asd")).IsFalse()
		})
	})




	g.Describe("removing user", func() {
		g.It("should return true", func() {
			g.Assert(PerformRemoveUser(api, testUser)).IsTrue()
		})

		g.It("should return false", func() {
			g.Assert(PerformRemoveUser(api, testUser)).IsFalse()
		})


	})
}

