package client

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// 中弘 VRF 集控器 B19
type b19client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

// NewClient creates a new Zhonghong client with given backend handler.
func NewB19Client(handler api.ClientHandler) api.Client {
	return &b19client{packager: handler, transporter: handler}
}

func (mb *b19client) ExecuteCommand(data []uint16, command int, OnOff int) (results *protocol.ProtocolDataUnit, err error) {
	addressLen := AddLen(data)
	var commandsArr []byte
	commandsArr = dataBlockArray(data[2:])
	if OnOff == 1 {
		commandsArr = AppendCommands(data, protocol.ON)
	} else if OnOff == 0 {
		commandsArr = AppendCommands(data, protocol.OFF)
	}
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: byte(command),
		Address:      addressLen,
		Commands:     commandsArr,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// B27Client sends the specified control command to the gateway
func (mb *b19client) send(request *protocol.ProtocolDataUnit) (response *protocol.ProtocolDataUnit, err error) {
	aduRequest, err := mb.packager.Encode(request)
	if err != nil {
		return
	}
	aduResponse, err := mb.transporter.Send(aduRequest)
	if err != nil {
		return
	}
	if err = mb.packager.Verify(aduRequest, aduResponse); err != nil {
		return
	}
	response, err = mb.packager.Decode(aduResponse)
	if err != nil {
		return
	}
	// Check correct function code returned (exception)
	if response.FunctionCode != request.FunctionCode {
		err = fmt.Errorf("zhonghong-b19 client: response function code does not match request")
		return
	}

	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("zhonghong-b19 client: response data is empty")
		return
	}
	return
}
