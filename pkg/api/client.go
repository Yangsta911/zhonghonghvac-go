package api

import "github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"

// Client interface for both B27 and B19 controllers
type Client interface {
	// Bit access
	ExecuteCommand(data []uint16, command int, OnOff int) (results *protocol.ProtocolDataUnit, err error)
}
