package master

import (
	"github.com/yongjuntang/nanoV0.5.0-dialect/component"
	"github.com/yongjuntang/nanoV0.5.0-dialect/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	// Topic service
	topicService = newTopicService()
	// ... other services
)

func init() {
	Services.Register(topicService)
}

func OnSessionClosed(s *session.Session) {
	topicService.userDisconnected(s)
}
