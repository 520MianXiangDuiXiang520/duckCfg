# duckCfg

A simpler and more unified way to read configuration files.

[中文](./README_ZH.md)

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
 
