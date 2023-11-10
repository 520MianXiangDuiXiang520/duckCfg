# duckCfg

A simpler and more unified way to read configuration files.

[中文](./README_ZH.md)

[![GoDoc](https://camo.githubusercontent.com/ba58c24fb3ac922ec74e491d3ff57ebac895cf2deada3bf1c9eebda4b25d93da/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f67616d6d617a65726f2f776f726b6572706f6f6c3f7374617475732e737667)](https://pkg.go.dev/github.com/520MianXiangDuiXiang520/duckCfg)


## Usage

```shell
go get github.com/520MianXiangDuiXiang520/duckCfg
```

```go
package duckcfg_test

import (
    "fmt"
    "github.com/520MianXiangDuiXiang520/duckCfg"
)

func GetIntFormatDefault() {
    err := duckcfg.InitConfig("./test_data/01.yaml")
    if err != nil {
        panic(err)
    }
    timeout := duckcfg.Cfg().GetIntFormatDefault("db.mongo.conn_timeout", 0)
    fmt.Println(timeout) // 3
}


```

If you want to read from JSON, you only need to modify the file name.

## Q&A

> Q: What if my key contains a dot?
> * Use a backslash to escape it, e.g., if your key is db.mongo, you can use db\\.mongo to read it. The same applies to keys containing a backslash.

> Q: My configuration file format is not supported.
> * You can define a FConfigLoader and use the Register function to register it.
 
