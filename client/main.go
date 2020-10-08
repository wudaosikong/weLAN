package client

import (
	"weLAN/client/Manager"
	"weLAN/client/tools"
	"fmt"
)

var ch = make(chan []byte, 10)

func Run() {
	fmt.Println("开始启动！")
	LocalIps := tools.GetIntranetIp()
	fmt.Print("你的ID是：")
	for _, LocalIp := range LocalIps {
		fmt.Println(LocalIp)
	}
	fmt.Println("输入 help 以获取更多帮助")

	gui := Manager.GUI{}
	gui.LocalIP = LocalIps
	gui.Render()
}
