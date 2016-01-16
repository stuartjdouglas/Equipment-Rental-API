package email
//
//import (
//	"testing"
//	. "github.com/franela/goblin"
//	"github.com/remony/Equipment-Rental-API/core/router"
//	"github.com/remony/Equipment-Rental-API/core/config"
//)
//const CONF_FILE = "./../../config.json"
//
//func TestUtilsEmail(t *testing.T) {
//	g := Goblin(t)
//
//		api := router.API{Context: config.Connection(config.LoadConfig(CONF_FILE, true).Production.DbUrl)}
//
//	g.Describe("Email", func() {
//		g.It("Should send", func() {
//			g.Assert(SendEmail(api, "remonasebi@gmail.com", "Unit test", "<h1>Yay!</h1><br>if you got this then it must have passed"))
//		})
//
//	})
//}
//
