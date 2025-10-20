package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int
	MinX      int
	MinY      int
	MaxX      int
	MaxY      int
	playerIDs map[int]bool
	pIDLock   sync.RWMutex
}

// 初始化一个格子 (Grid struct的实例）
func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gID,
		MinX:      minX,
		MinY:      minY,
		MaxX:      maxX,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

// 添加玩家ID到格子中
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playerIDs[playerID] = true
}

// 从格子中删除玩家ID
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	delete(g.playerIDs, playerID)
}

// 获取格子中的所有玩家ID
func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()

	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}

	return
}

// 打印信息方法
func (g *Grid) String() string {
	return fmt.Sprintf("Grid ID: %d, MinX: %d, MaxX: %d, MinY: %d, MaxY: %d, PlayerIDs: %v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.GetPlayerIDs())
}
