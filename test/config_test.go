package config
import (
	"testing"
	"../core/config"
	"github.com/stretchr/testify/assert"
)
const CONF_FILE = "./../config.json"

func TestLoadConfig(t *testing.T) {

	assert.NotNil(t, config.LoadConfig(CONF_FILE), "should not be empty")
}

func TestHasTitle(t *testing.T) {
	assert.NotNil(t, config.LoadConfig(CONF_FILE).Title, "should be a title")
}

func TestHasMongoDB(t *testing.T) {
	assert.NotNil(t, config.LoadConfig(CONF_FILE).MongoDb, "should be a Mongo Url")
}

func TestHasPort(t *testing.T) {
	assert.NotNil(t, config.LoadConfig(CONF_FILE).Port, "should be a port number")
}



//
//func TestPortIsInt(t *testing.T) {
//	assert.IsType(t, config.LoadConfig().Port, types.int)
//}