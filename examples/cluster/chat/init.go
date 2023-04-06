package chat

import (
	"github.com/yongjuntang/nanoV0.5.0-dialect/component"
	"github.com/yongjuntang/nanoV0.5.0-dialect/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	roomService = newRoomService()
)

func init() {
	Services.Register(roomService)
}

func OnSessionClosed(s *session.Session) {
	roomService.userDisconnected(s)
}
