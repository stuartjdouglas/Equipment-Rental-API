package config_test
import (
	"github.com/stretchr/testify/assert"
	"testing"
	"../core/config"
)


func TestTest(t *testing.T) {
	assert.Equal(t, 1, 1, "should be equal")
}

func TestDatabaseConnectionGivenEmptyString(t *testing.T) {
	assert.Nil(t, config.Connection("").Session, "Should return nil session")
}

func TestDatabaseConnectionGivenWrongDetails(t *testing.T) {
	assert.Nil(t, config.Connection("").Session, "Should return nil session")
}

func TestDatabaseConnectionGivenWrongUrl(t *testing.T) {
	assert.Nil(t, config.Connection("").Session, "Should return nil session")
}

// This test will get the mongodb url from the config method
func TestDatabaseConnectionGivenCorrectDetails(t *testing.T) {
	assert.NotNil(t, config.Connection(config.LoadConfig().MongoDb).Session, "Should return session")
}
