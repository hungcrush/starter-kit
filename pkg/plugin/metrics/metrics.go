package metrics

import (
	"github.com/stack-labs/stack-rpc-plugins/service/stackway/plugin"
)

//NewPlugin of metrics
func NewPlugin(opts ...Option) plugin.Plugin {
	return newPrometheus(opts...)
}
