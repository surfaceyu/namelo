package chat

import (
	"github.com/surfaceyu/namelo/component"
	"github.com/surfaceyu/namelo/session"
)

var (
	// Services All services in master server
	Services = &component.Components{}

	roomService = newRoomService()
)

func init() {
	Services.Register(roomService)
}

// OnSessionClosed HH
func OnSessionClosed(s *session.Session) {
	roomService.userDisconnected(s)
}
