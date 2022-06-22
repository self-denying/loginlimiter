package constant

import "time"

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 11:22
**/

const (
	WriteWait      = 10 * time.Second
	PongWait       = 60 * time.Second
	PingPeriod     = (PongWait * 9) / 10
	MaxMessageSize = 512
)

type MessageType int8

const (
	KickOut MessageType = iota + 1
	BroadCast
	OnlyOne
)
