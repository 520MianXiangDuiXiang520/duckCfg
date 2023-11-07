package duckcfg

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var pwd = "./test_data/"

func testPath(fileName string) string {
	return pwd + fileName
}

func testHelper(t *testing.T, path string, expect any, key string) {
	sourceData, err := JsonLoader(path)
	assert.Nil(t, err, "load source file fail")
	res, err := read(sourceData, key)
	assert.Nil(t, err, "load source file fail")
	assert.Truef(t, reflect.DeepEqual(res, expect), "got: %#v", res)
}

func Test_read(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			100.0, "db.mongo.max_connecting")
	})
	t.Run("t2", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			[]any{"snappy", "zlib", "zstd"}, "db.mongo.compressors")
	})
	t.Run("t3", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			"zstd", "db.mongo.compressors.2")
	})
	t.Run("t4", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			"local", "db.mongo\\\\\\.url")
	})
	t.Run("t5", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			"local2", "db.mongo\\\\\\.")
	})
	t.Run("t6", func(t *testing.T) {
		testHelper(t, testPath("01.json"),
			"local3", "db.\\\\\\.mongo")
	})
}
