package config
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"../core/config"
)

func TestLoadConfig(t *testing.T) {
	assert.NotNil(t, config.LoadConfig(), "should not be empty")
}

func TestHasTitle(t *testing.T) {
	assert.NotNil(t, config.LoadConfig().Title, "should be a title")
}

func TestHasMongoDB(t *testing.T) {
	assert.NotNil(t, config.LoadConfig().MongoDb, "should be a Mongo Url")
}

func TestHasPort(t *testing.T) {
	assert.NotNil(t, config.LoadConfig().Port, "should be a port number")
}



//
//func TestPortIsInt(t *testing.T) {
//	assert.IsType(t, config.LoadConfig().Port, types.int)
//}