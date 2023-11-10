package duckcfg

import (
	"encoding/json"
	"github.com/pkg/errors"
	path2 "path"
	"path/filepath"
)

//type IConfig interface {
//	GetString(k string) (r string, err error)
//	GetStringDefault(k, d string) string
//
//	GetInt(k string) (r int, err error)
//	GetIntDefault(k string, d int) int
//
//	GetInt64(k string) (r int64, err error)
//	GetInt64Default(k string, d int64) int64
//
//	GetFloat(k string) (r float64, err error)
//	GetFloatDefault(k string, d float64) float64
//
//	// GetIntFormat 如果 key 对应的 value 是 float64 int64, int 类型，
//	// 强制转换成 int 类型并返回，浮点向下取整
//	GetIntFormat(k string) (r int, err error)
//	GetIntFormatDefault(k string, d int) (r int)
//
//	// GetFloatFormat 如果 key 对应的 value 是 float64 int64, int 类型，
//	// 强制转换成 float64 类型并返回
//	GetFloatFormat(k string) (r float64, err error)
//	GetFloatFormatDefault(k string, d float64) (r float64)
//
//	GetBool(k string) (r bool, err error)
//	GetBoolDefault(k string, d bool) bool
//
//	GetVal(k string, to any) (r any, err error)
//
//	GetStrings(k string) (r []string, err error)
//	GetInts(k string) (r []int64, err error)
//	GetFloats(k string) (r []float64, err error)
//	GetBooleans(k string) (r []bool, err error)
//	GetMap(k string) (r map[string]any, err error)
//}

// IConfig interface defines a set of generic methods
// for getting values from configuration files.
type IConfig interface {

	// GetString retrieve the string value for the
	// specified key from the configuration file.
	GetString(k string) (r string, err error)

	// GetStringDefault retrieve the string value for the
	// specified key from the configuration file,
	// or return the default value if it doesn't exist.
	GetStringDefault(k, d string) string

	// GetInt Retrieve the integer value for the
	// specified key from the configuration file.
	GetInt(k string) (r int, err error)

	// GetIntDefault Retrieve the integer value for the
	// specified key from the configuration file,
	// or return the default value if it doesn't exist.
	GetIntDefault(k string, d int) int

	// GetInt64 Retrieve the int64 value for the specified
	// key from the configuration file.
	GetInt64(k string) (r int64, err error)

	// GetInt64Default Retrieve the int64 value for the
	// specified key from the configuration file,
	// or return the default value if it doesn't exist.
	GetInt64Default(k string, d int64) int64

	// GetFloat Retrieve the float value for the specified
	// key from the configuration file.
	GetFloat(k string) (r float64, err error)

	// GetFloatDefault Retrieve the float value for the
	// specified key from the configuration file,
	// or return the default value if it doesn't exist.
	GetFloatDefault(k string, d float64) float64

	// GetIntFormat If the value for the specified key
	// in the configuration file is of type float64, int64, or int,
	// force conversion to int and return.
	GetIntFormat(k string) (r int, err error)

	// GetIntFormatDefault similar to GetIntFormat,
	// except that if the key does not exist or the conversion fails,
	// the default value is returned
	GetIntFormatDefault(k string, d int) (r int)

	// GetFloatFormat If the value for the specified key in the
	// configuration file is of type float64, int64, or int,
	// force conversion to float64 and return.
	GetFloatFormat(k string) (r float64, err error)

	// GetFloatFormatDefault similar to GetFloatFormat,
	// except that if the key does not exist or the conversion fails,
	// the default value is returned
	GetFloatFormatDefault(k string, d float64) (r float64)

	// GetBool Retrieve the boolean value for the
	// specified key from the configuration file.
	GetBool(k string) (r bool, err error)

	// GetBoolDefault Retrieve the boolean value for the
	// specified key from the configuration file,
	// or return the default value if it doesn't exist.
	GetBoolDefault(k string, d bool) bool

	// GetVal Read the value corresponding to key from
	//the configuration file and deserialize it to the
	// parameter to,to should be a pointer
	GetVal(k string, to any) (r any, err error)

	// GetStrings Read the string list for the
	// specified key from the configuration file.
	GetStrings(k string) (r []string, err error)

	// GetIntegers Read the integers list for the
	// specified key from the configuration file.
	GetIntegers(k string) (r []int64, err error)

	// GetFloats Read the float list for the
	// specified key from the configuration file.
	GetFloats(k string) (r []float64, err error)

	// GetBooleans Read the boolean list for the
	// specified key from the configuration file.
	GetBooleans(k string) (r []bool, err error)

	// GetMap Read the mapt for the
	// specified key from the configuration file.
	GetMap(k string) (r map[string]any, err error)
}

func Cfg() IConfig {
	return _cfg
}

var _cfg = &cfg{}

func InitConfig(path string) error {
	_, fileName := path2.Split(path)
	tp := filepath.Ext(fileName)
	loader, ok := fileLoader[tp]
	if !ok {
		return errors.WithMessagef(ErrorFailToLoadConfig,
			"load func not found, type: %s", tp)
	}
	data, err := loader(path)
	if err != nil {
		return err
	}
	_cfg.data = data
	return nil
}

type cfg struct {
	data any
}

func get[T any](c *cfg, k string) (r T, err error) {
	res, err := read(c.data, k)
	if err != nil {
		return r, err
	}
	r, ok := res.(T)
	if !ok {
		return r, ErrorTypeMismatch
	}
	return r, nil
}

func getDefault[T any](c *cfg, k string, d T) (r T) {
	v, err := get[T](c, k)
	if err != nil {
		return d
	}
	return v
}

func gets[S ~[]T, T any](c *cfg, k string) (r S, err error) {
	res, err := get[[]any](c, k)
	if err != nil {
		return
	}
	r = make([]T, len(res))
	for i, re := range res {
		d, ok := re.(T)
		if !ok {
			err = ErrorTypeMismatch
			return
		}
		r[i] = d
	}
	return
}

func (c *cfg) GetString(k string) (r string, err error) {
	return get[string](c, k)
}

func (c *cfg) GetStringDefault(k, d string) string {
	return getDefault(c, k, d)
}

func (c *cfg) GetIntFormat(k string) (r int, err error) {
	v, err := get[any](c, k)
	if err != nil {
		return
	}
	switch t := v.(type) {
	case int:
		r = t
		return
	case int64:
		r = int(t)
		return
	case float64:
		r = int(t)
		return
	default:
		err = ErrorTypeMismatch
		return
	}
}

func (c *cfg) GetIntFormatDefault(k string, d int) (r int) {
	v, err := c.GetIntFormat(k)
	if err != nil {
		return d
	}
	return v
}

func (c *cfg) GetFloatFormat(k string) (r float64, err error) {
	v, err := get[any](c, k)
	if err != nil {
		return
	}
	switch t := v.(type) {
	case int:
		r = float64(t)
		return
	case int64:
		r = float64(t)
		return
	case float64:
		r = t
		return
	default:
		err = ErrorTypeMismatch
		return
	}
}

func (c *cfg) GetFloatFormatDefault(k string, d float64) (r float64) {
	v, err := c.GetFloatFormat(k)
	if err != nil {
		return d
	}
	return v
}

func (c *cfg) GetInt(k string) (r int, err error) {
	return get[int](c, k)
}

func (c *cfg) GetIntDefault(k string, d int) int {
	return getDefault(c, k, d)
}

func (c *cfg) GetInt64(k string) (r int64, err error) {
	return get[int64](c, k)
}

func (c *cfg) GetInt64Default(k string, d int64) int64 {
	return getDefault(c, k, d)
}

func (c *cfg) GetFloat(k string) (r float64, err error) {
	return get[float64](c, k)
}

func (c *cfg) GetFloatDefault(k string, d float64) float64 {
	return getDefault(c, k, d)
}

func (c *cfg) GetBool(k string) (r bool, err error) {
	return get[bool](c, k)
}

func (c *cfg) GetBoolDefault(k string, d bool) bool {
	return getDefault(c, k, d)
}

func (c *cfg) GetVal(k string, to any) (r any, err error) {
	source, err := get[any](c, k)
	if err != nil {
		return
	}
	data, err := json.Marshal(source)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, to)
	if err != nil {
		return
	}
	return to, nil
}

func (c *cfg) GetStrings(k string) (r []string, err error) {
	return gets[[]string](c, k)
}

func (c *cfg) GetIntegers(k string) (r []int64, err error) {
	return gets[[]int64](c, k)
}

func (c *cfg) GetFloats(k string) (r []float64, err error) {
	return gets[[]float64](c, k)
}

func (c *cfg) GetBooleans(k string) (r []bool, err error) {
	return gets[[]bool](c, k)
}

func (c *cfg) GetMap(k string) (r map[string]any, err error) {
	return get[map[string]any](c, k)
}
