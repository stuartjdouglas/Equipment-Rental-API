package models

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func TestProductModel(t *testing.T) {
	g := Goblin(t)

	api := router.API{Context:config.Connection(config.LoadConfig(ConfigFile, true).Production.DbUrl)}

	g.Describe("getting products", func() {
		products := GetProducts(api)
		g.It("should contain a total", func() {
			g.Assert(products.Total >= 0).IsTrue()
		})

		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})

	g.Describe("getting paged products", func() {
		products := GetProductsPaging(api, 0, 10, "", true)
		g.It("should return 10 items", func() {
			g.Assert(products.Total == 10).IsTrue()
		})
		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})

	g.Describe("getting paged products sorted by added", func() {
		products := GetProductsPagingSortedByAdded(api, 0, 10, "", true)
		g.It("should return 10 items", func() {
			g.Assert(products.Total == 10).IsTrue()
		})
		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})

	g.Describe("getting paged products sorted by updated", func() {
		products := GetProductsPagingSortedByUpdated(api, 0, 10, "", true)
		g.It("should return 10 items", func() {
			g.Assert(products.Total == 10).IsTrue()
		})
		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})

	g.Describe("getting paged products sorted by likes", func() {
		products := GetProductsPagingSortedByLikes(api, 0, 10, "", true)
		g.It("should return 10 items", func() {
			g.Assert(products.Total == 10).IsTrue()
		})
		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})

	g.Describe("getting paged products sorted by random", func() {
		products := GetProductsPagingRandom(api, 0, 10, "")
		g.It("should return 10 items", func() {
			g.Assert(products.Total == 10).IsTrue()
		})
		if (products.Total > 0) {
			g.It("should contain a title", func() {
				g.Assert(len(products.Items[0].Product_name) > 0).IsTrue()
			})
		}
	})
}

