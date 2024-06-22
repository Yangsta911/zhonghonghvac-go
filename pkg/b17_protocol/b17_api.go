package b17protocol

import (
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

type Clientb17 interface {
	// Bit access
	//todo change result to struct instead of byte
	ReadGateway() (results *protocol.ProtocolDataUnit, err error)
	EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayControlOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayControlOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayTempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayWindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayWindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayNewAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayNewAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayNewAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
	GatewayNewAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error)
}
