package main

import (
	"fmt"

	"mmo-game/core"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

func main() {
	// 创建服务器句柄
	s := znet.NewServer()

	// 注册连接创建时的 Hook 函数（客户端链接建立和丢失）
	s.SetOnConnStart(OnConnectionAdd)

	// 启动服务器
	s.Serve()
}

func OnConnectionAdd(conn ziface.IConnection) {
	// 创建一个玩家
	player := core.NewPlayer(conn)
	// 同步玩家ID给客户端, MsgID = 1
	player.SyncPid()
	// 同步玩家的初始化坐标给客户端, MsgID = 200
	player.BroadCastStartPosition()
	fmt.Println("=====> Player pid =", player.Pid, "arrived =====")
}
