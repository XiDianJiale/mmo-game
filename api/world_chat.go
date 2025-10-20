package api

import (
	"fmt"
	"mmo-game/core"
	"mmo-game/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (*WorldChatApi) Handle(request ziface.IRequest) {
	//1.将客户端传来的proto解码
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("WorldChatApi Handle Unmarshal err:", err)
		return
	}
	//2.从 连接的属性pid 中获取是哪个玩家传来的
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("WorldChatApi Handle GetProperty err:", err)
		request.GetConnection().Stop()
		return
	}
	//3.更具pid获取player对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	//4.让player对象发起广播请求
	player.Talk(msg.Content)
}
