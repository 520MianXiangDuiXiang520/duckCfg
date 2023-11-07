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
