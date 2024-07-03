package client

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/api"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}

// 中弘线控器 B27（小超人）
type b27client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

// NewB27Client creates a new Zhonghong client with given backend handler.
func NewB27Client(handler ClientHandler) api.Client {
	return &b27client{packager: handler, transporter: handler}
}

// PerformanceCheck returns performances statistics of the specified HVAC device
func (mb *b27client) PerformanceCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodePerformanceCheck,
		CommandType:  "remote",
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// StatusCheck returns status of the specified HVAC device
func (mb *b27client) StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeStatusCheck,
		CommandType:  "remote",
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// On turns on specified HVAC device
func (mb *b27client) On(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.ON)
	commandsOn := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeOnOff,
		Address:      addressLen,
		Commands:     commandsOn,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Off turns off specified HVAC device
func (mb *b27client) Off(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.OFF)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ErrorCheck returns the error status code of the specified HVAC device
func (mb *b27client) ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeErrorCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirPerformance returns the performance statistics of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFreshAirPerformance,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirStatus returns the status of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFreshAirStatus,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirModeControl selects the mode of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFreshAirControl,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns on the specified Fresh Air ventilation device
func (mb *b27client) FreshAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.ON)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirOn turns off the specified Fresh Air ventilation device
func (mb *b27client) FreshAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.OFF)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FreshAirErrorCheck returns the error status of the specified Fresh Air ventilation device
func (mb *b27client) FreshAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFreshAirErrorCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingPerformance returns the performance statistics of the specified Floor Heating device
func (mb *b27client) FloorHeatingPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFloorHeatingPerformance,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingStatus returns the status of the specified Floor Heating device
func (mb *b27client) FloorHeatingStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FuncCodeFloorHeatingStatusCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOn turns on the specified Floor Heating device
func (mb *b27client) FloorHeatingOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.ON)
	commandsOn := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FloorHeatingOnOff,
		Address:      addressLen,
		Commands:     commandsOn,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOff turns off the specified Floor Heating device
func (mb *b27client) FloorHeatingOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.OFF)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCode,
		FunctionCode: protocol.FloorHeatingOnOff,
		Address:      addressLen,
		Commands:     commandsOff,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *b27client) Control(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FreshAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingAntiFreezeOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingAntiFreezeOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

func (mb *b27client) FloorHeatingTemp(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b27 client: does not support following protocol")
}

// B27Client sends the specified control command to the specified device
func (mb *b27client) send(request *protocol.ProtocolDataUnit) (response *protocol.ProtocolDataUnit, err error) {
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
		err = fmt.Errorf("zhonghong-b27 client: response function code does not match request")
		return
	}
	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("zhonghong-b27 client: response data is empty")
		return
	}
	return
}
