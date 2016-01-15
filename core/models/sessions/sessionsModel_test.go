package sessions

import (
"testing"
. "github.com/franela/goblin"
)
const CONF_FILE = "./../../config.json"

func TestAuthModel(t *testing.T) {
g := Goblin(t)

//router := router.API{Context: database.Connection(config.LoadConfig(CONF_FILE, true).Production.DbUrl)}

g.Describe("Sessions", func() {
//		TODO get token from a login
//g.It("Should be valid", func() {
//g.Assert(IsSessionValid(router, "ba36e0c4-7c30-42cf-8d29-4bb57442ef9c")).IsTrue()
//})

////g.It("Should be invalid", func() {
////g.Assert(IsSessionValid(router, "master-code")).IsFalse()
//})
})
}


