package duckcfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCfg_AllJson(t *testing.T) {
	err := InitConfig(testPath("01.json"))
	assert.Nil(t, err, "load fail")

	assert.Equal(t, 3.0, Cfg().GetFloatDefault("db.mongo.conn_timeout", 0))

	cs, err := Cfg().GetStrings("db.mongo.compressors")
	assert.Nil(t, err, "load fail")
	assert.Equal(t, []string{"snappy", "zlib", "zstd"}, cs)
}

func TestCfg_AllYaml(t *testing.T) {
	err := InitConfig(testPath("01.yaml"))
	assert.Nil(t, err, "load fail")

	assert.Equal(t, 3, Cfg().GetIntFormatDefault("db.mongo.conn_timeout", 0))

	cs, err := Cfg().GetStrings("db.mongo.compressors")
	assert.Nil(t, err, "load fail")
	assert.Equal(t, []string{"snappy", "zlib", "zstd"}, cs)
}
