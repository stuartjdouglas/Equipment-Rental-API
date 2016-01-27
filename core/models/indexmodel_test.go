package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func TestIndexModel(t *testing.T) {
	g := Goblin(t)

	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	g.Describe("Index", func() {
		g.It("should return title", func() {
			g.Assert(len(GetIndex(api).Title) > 0).IsTrue()
		})
		g.It("should return description", func() {
			g.Assert(len(GetIndex(api).Description) > 0).IsTrue()
		})
	})
}

