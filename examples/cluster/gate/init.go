package gate

import "github.com/yongjuntang/nanoV0.5.0-dialect/component"

var (
	// All services in master server
	Services = &component.Components{}

	bindService = newBindService()
)

func init() {
	Services.Register(bindService)
}
