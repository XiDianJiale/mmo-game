package core

import (
	"sync"
)

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

type WorldManager struct {
	AoiMgr  *AOIManager
	Players map[int32]*Player //这个动态map和上一个项目同理
	pLock   sync.RWMutex
}

// 提供对外的世界管理模块句柄
var WorldMgrObj *WorldManager

// WorldMgr初始化方法
func init() {
	WorldMgrObj = &WorldManager{
		Players: make(map[int32]*Player),
		AoiMgr: NewAOIManager(AOI_MIN_X, AOI_MAX_X,
			AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}
}

// 添加玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Y)

}

// 删除玩家
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

// 通过玩家ID获取玩家信息
func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pid]
}

// 获取世界中所有玩家信息
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player, 0)

	for _, v := range wm.Players {
		players = append(players, v)
	}
	return players
}
