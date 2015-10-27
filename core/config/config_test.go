package config
import (
	"testing"
	. "github.com/franela/goblin"
)

const CONF_FILE = "./../../config.json"


func TestConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("Development config", func() {
		g.It("Should have title ", func() {
			g.Assert(LoadConfig(CONF_FILE, true).Title == "").IsFalse()
		})
		g.It("Should have database url", func() {
			g.Assert(LoadConfig(CONF_FILE, true).DbUrl == "").IsFalse()
		})
		g.It("Should have port number", func() {
			g.Assert(LoadConfig(CONF_FILE, true).Port == 0).IsFalse()
		})
//		g.It("port number should be int", func() {
//			g.Assert(LoadConfig(CONF_FILE, true).Port == reflect.Int ).IsTrue()
//		})
	})

	g.Describe("Production config", func() {
		g.It("Should have title ", func() {
			g.Assert(LoadConfig(CONF_FILE, false).Title == "").IsFalse()
		})
		g.It("Should have database url", func() {
			g.Assert(LoadConfig(CONF_FILE, false).DbUrl == "").IsFalse()
		})
		g.It("Should have port number", func() {
			g.Assert(LoadConfig(CONF_FILE, false).Port == 0).IsFalse()
		})
	})
}
