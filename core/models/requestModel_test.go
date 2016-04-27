package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func TestRequestModel(t *testing.T) {
	g := Goblin(t)

	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	g.Describe("getting product requests", func() {
		products := GetProductsRequests(api, "", "0", "10")
		g.It("should be empty with invalid token", func() {
			g.Assert(products.Total == 0).IsTrue()
			g.Assert(products.Total > 0).IsFalse()
			g.Assert(products.Total).Equal(0)
		})
	})

	g.Describe("getting product requests", func() {
		products := GetUserRequests(api, "", "0", "10")
		g.It("should return empty with invalid token", func() {
			g.Assert(products.Total == 0).IsTrue()
			g.Assert(products.Total > 0).IsFalse()
			g.Assert(products.Total).Equal(0)
		})
	})

}

