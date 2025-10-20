package core

import (
	"fmt"
	"math/rand"
	"mmo-game/pb"
	"sync"

	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection //当前玩家的连接
	X    float32
	Y    float32
	Z    float32
	V    float32 //玩家朝向,旋转0-360度
}

/*
Player ID 生成器
*/
var PidGen int32 = 1  //生成玩家ID的计数器
var IdLock sync.Mutex //保护PidGen的互斥机制

// 创建玩家对象
func NewPlayer(conn ziface.IConnection) *Player {
	//生成玩家PID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(134 + rand.Intn(17)),
		V:    0,
	}

	return p
}

/*
发送消给客户端，
主要是将pb的protobuf数据序列息化之后发送
*/

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	fmt.Printf("before Marshal data = %v\n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err:", err)
	}
	if p.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("player send msg err:", err)
		return
	}
	return
}

// 同步玩家ID给客户端
func (p *Player) SyncPid() {
	//组建MsgId0 proto数据
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	p.SendMsg(1, data)
}

// 广播玩家自己的出生地点
func (p *Player) BroadCastStartPosition() {

	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			&pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}

// 广播玩家聊天
func (p *Player) Talk(content string) {
	//1. 组建Msg200 proto数据
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1, //TP1 代表聊天消息
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	//2.得到当前世界所有玩家
	player := WorldMgrObj.GetAllPlayers()
	//3.循环给所有玩家发送消息
	for _, player := range player {
		player.SendMsg(200, msg)
	}
}
