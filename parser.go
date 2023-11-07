package duckcfg

import (
	"bytes"
	"github.com/pkg/errors"
	"strconv"
)

func handler(data any, key string) (res any, err error) {
	switch v := data.(type) {
	case map[string]any:
		res, ok := v[key]
		if !ok {
			return nil,
				errors.WithMessagef(ErrorKeyNotFound, "key: %s", key)
		}
		return res, nil
	case []any:
		idx, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, errors.WithMessagef(ErrorBadPathStr,
				"object is a slice, key must is a number, got: %s", key)
		}
		if int(idx) >= len(v) {
			return nil, errors.WithMessagef(ErrorKeyNotFound,
				"out of range, max: %d, got: %d", len(v), idx)
		}
		return v[idx], nil
	default:
		return nil,
			errors.WithMessagef(ErrorKeyNotFound, "key: %s", key)
	}
}

func read(data any, path string) (res any, err error) {
	keyBuffer := bytes.NewBuffer([]byte{})
	pathSlice := []byte(path)
	pathSize := len(pathSlice)
	for i := 0; i < pathSize; i++ {
		s := pathSlice[i]
		switch s {
		case '\\':
			if i == pathSize-1 {
				return nil, errors.WithMessage(ErrorBadPathStr,
					"after the transfer symbol, only allow itself or '.' ")
			}
			next := pathSlice[i+1]
			switch next {
			case '\\':
				keyBuffer.WriteByte('\\')
				i++
				continue
			case '.':
				keyBuffer.WriteByte('.')
				i++
				continue
			default:
				return nil, errors.WithMessage(ErrorBadPathStr,
					"after the transfer symbol, only allow itself or '.' ")
			}
		case '.':
			key := keyBuffer.String()
			keyBuffer.Reset()
			data, err = handler(data, key)
			if err != nil {
				return nil, err
			}
		default:
			keyBuffer.WriteByte(s)
		}
	}
	if keyBuffer.Len() > 0 {
		key := keyBuffer.String()
		data, err = handler(data, key)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return data, nil
}
