package auth

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestAuthModel(t *testing.T) {
	g := Goblin(t)

	//	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Checking valid username", func() {
		g.It("lemon should return true", func() {
			g.Assert(isValidUsername("lemon")).IsTrue()
		})
		g.It("of over 240 character should return false", func() {
			g.Assert(isValidUsername("lemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemonlemon")).IsFalse()
		})
		g.It("containing a character should return false", func() {
			g.Assert(isValidUsername("$now")).IsFalse()
		})

	})


}

