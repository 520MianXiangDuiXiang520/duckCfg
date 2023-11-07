package duckcfg

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

// FConfigLoader 加载器，将配置文件数据反序列化成 interface{} 并返回
type FConfigLoader func(path string) (any, error)

var fileLoader = map[string]FConfigLoader{
	".json": JsonLoader,
	".yaml": YamlLoader,
	".yml":  YamlLoader,
}

// Register 为某种文件类型注册一个加载器
func Register(fileType string, loader FConfigLoader) {
	fileLoader[fileType] = loader
}

func JsonLoader(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.WithMessagef(
			ErrorLoadConfigFail, "err: %s", err.Error())
	}
	var d interface{}
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, errors.WithMessagef(
			ErrorLoadConfigFail, "unmarshal fail err: %s", err.Error())
	}
	return d, nil
}

func YamlLoader(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.WithMessagef(
			ErrorLoadConfigFail, "err: %s", err.Error())
	}
	var d interface{}
	err = yaml.Unmarshal(data, &d)
	if err != nil {
		return nil, errors.WithMessagef(
			ErrorLoadConfigFail, "unmarshal fail err: %s", err.Error())
	}
	return d, nil
}
