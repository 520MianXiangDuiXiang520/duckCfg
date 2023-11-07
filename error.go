package duckcfg

import "errors"

var (
	ErrorFailToLoadConfig = errors.New("fail to load config")
	ErrorLoadConfigFail   = errors.New("load config fail")
	ErrorBadPathStr       = errors.New("bad path str")
	ErrorKeyNotFound      = errors.New("key not found")
	ErrorTypeMismatch     = errors.New("type mismatch")
)
