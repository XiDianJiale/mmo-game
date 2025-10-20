```bash
└── mmo_game
    ├── api
    ├── conf
    │   └── zinx.json
    ├── core
    │   ├── aoi.go
    │   ├── aoi_test.go
    │   ├── grid.go
    ├── game_client
    │   └── client.exe
    ├── pb
    │   ├── build.sh
    │   └── msg.proto
    ├── README.md
    └── server.go
```


目前使用win可以进行内网连接，检测到上下线，但是无法使用player中的SyncPid()  BroadCastStartPosition() Talk(content string) 全局广播更没有实现


11.8-11.11功能未实现

### 参考链接
https://www.yuque.com/aceld/npyr8s/om3nf2