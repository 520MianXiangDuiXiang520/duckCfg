# duckCfg

用更简单，更统一的方法读取配置文件

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

如果你想从 json 中读取，只需要修改文件名即可

## Q&A

> Q: 键中有 . 怎么办？
> * 使用 \ 转义，如你的键名为 `db.mongo` 可以使用 `db\\.mongo` 来读取 键中包含 \ 同理

> Q: 没有支持我的配置文件格式
> * 你可以自定义一个 `FConfigLoader` 并使用 `Register` 函数注册

    