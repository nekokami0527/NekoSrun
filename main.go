package main

import (
	"flag"
	"log"

	"nekokami0527.com/nekosrun/nekosrun"
)

var (
	username = flag.String("u", "", "学工号")
	password = flag.String("p", "", "密码")
	server   = flag.String("s", "192.168.118.51", "深澜服务器地址")
	clientIp = flag.String("c", "", "客户端地址")
	logout   = flag.Bool("logout", false, "注销当前设备登录状态")
)

func main() {
	flag.Parse()
	if username == nil || password == nil || len(*username) == 0 || len(*password) == 0 || len(*server) == 0 {
		log.Fatalln("请跟参数-h获取使用说明")
	}
	srun := nekosrun.New(*server)

	if *clientIp != "" {
		srun.SetClientIp(*clientIp)
	}
	if *logout {
		log.Println("登出当前设备")
		srun.Logout()
		return
	}
	srun.Login(*username, *password)

}
