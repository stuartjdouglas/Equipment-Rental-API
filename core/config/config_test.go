package config
import (
	"testing"
	. "github.com/franela/goblin"
)

const CONF_FILE = "./../../../config.json"


func TestConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("Development config", func() {
		g.It("Should have title ", func() {
			g.Assert(LoadConfig(CONF_FILE, true).Development.Title == "").IsFalse()
		})
		g.It("Should have database url", func() {
			g.Assert(LoadConfig(CONF_FILE, true).Development.DbUrl == "").IsFalse()
		})
		g.It("Should have port number", func() {
			g.Assert(LoadConfig(CONF_FILE, true).Development.Port == 0).IsFalse()
		})
	})

	g.Describe("Production config", func() {
		g.It("Should have title ", func() {
			g.Assert(LoadConfig(CONF_FILE, false).Production.Title == "").IsFalse()
		})
		g.It("Should have database url", func() {
			g.Assert(LoadConfig(CONF_FILE, false).Production.DbUrl == "").IsFalse()
		})
		g.It("Should have port number", func() {
			g.Assert(LoadConfig(CONF_FILE, false).Production.Port == 0).IsFalse()
		})
	})
}
