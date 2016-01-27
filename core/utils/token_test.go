package utils

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestUtilsToken(t *testing.T) {
	g := Goblin(t)

	g.Describe("Tokens", func() {
		g.It("should be generated with a fixed length of 36 characters", func() {
			g.Assert(len(GenerateToken("Hello")) == 36).IsTrue()
			g.Assert(len(GenerateToken("Hello, World!")) == 36).IsTrue()
			g.Assert(len(GenerateToken("0123456789")) == 36).IsTrue()
			g.Assert(len(GenerateToken("qwertyuiopasdfghjklzxcvbnm")) == 36).IsTrue()
		})

		//		g.It("should be of type string", func() {
		//			str := GenerateToken("Hello, World!")
		//			str, boolean := str.(string)
		//			g.Assert(boolean).IsTrue()
		//		})
	})
}
