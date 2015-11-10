package utils

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestUtilsString(t *testing.T) {
	g := Goblin(t)

	g.Describe("Strings", func() {
		g.It("generate a random set of 10 characters", func() {
			g.Assert(len(RandomString(10)) == 10).IsTrue()
			g.Assert(len(RandomString(10)) == 10).IsTrue()
			g.Assert(len(RandomString(10)) == 10).IsTrue()
			g.Assert(len(RandomString(10)) == 10).IsTrue()

			g.Assert(len(RandomString(10)) == 11).IsFalse()
			g.Assert(len(RandomString(10)) == 11).IsFalse()
			g.Assert(len(RandomString(10)) == 11).IsFalse()
			g.Assert(len(RandomString(10)) == 11).IsFalse()
		})

		g.It("generate a random set of 100 characters", func() {
			g.Assert(len(RandomString(100)) == 100).IsTrue()
			g.Assert(len(RandomString(100)) == 100).IsTrue()
			g.Assert(len(RandomString(100)) == 100).IsTrue()
			g.Assert(len(RandomString(100)) == 100).IsTrue()
			g.Assert(len(RandomString(100)) == 11).IsFalse()
			g.Assert(len(RandomString(100)) == 11).IsFalse()
			g.Assert(len(RandomString(100)) == 11).IsFalse()
			g.Assert(len(RandomString(100)) == 11).IsFalse()
		})

	})
}
