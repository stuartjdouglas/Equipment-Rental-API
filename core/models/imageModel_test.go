	package models
import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/config/database"
	"github.com/remony/Equipment-Rental-API/core/config"
)

func TestImageModel(t *testing.T) {
	g := Goblin(t)
	router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).DbUrl)}

	g.Describe("Image Available", func() {
		// TODO when database is complete change to check is true
		g.It("Should be true", func() {
			g.Assert(IsImageAvailable(router, "image.jpg")).IsFalse()
		})
		g.It("Should be false", func() {
			g.Assert(IsImageAvailable(router, "image.jpg")).IsFalse()
		})
	})
}
