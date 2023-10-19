package main

import (
	"fmt"
	"github.com/linx93/chat-demo/init/conf"
)

// 启动入口
func main() {
	//加载配置
	c := conf.LoadConf()

	fmt.Printf("配置信息=>%v", c)
}
