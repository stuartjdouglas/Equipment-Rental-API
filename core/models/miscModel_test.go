package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func TestMiscModel(t *testing.T) {
	g := Goblin(t)

	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	g.Describe("Liking", func() {
		g.It("should return false with no token", func() {
			g.Assert(Like(api, "123", "567")).IsFalse()
		})
	})
	g.Describe("UnLiking", func() {
		g.It("should return false with no token", func() {
			g.Assert(UnLike(api, "123", "567")).IsFalse()
		})
	})
}

