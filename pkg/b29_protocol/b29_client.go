package b29protocol

import (
	"encoding/binary"
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/clientinterface"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

// ClientHandler is the interface that groups the Packager and Transporter methods.
type ClientHandler interface {
	protocol.Packager
	protocol.Transporter
}

type client struct {
	packager    protocol.Packager
	transporter protocol.Transporter
}

// NewClient creates a new Zhonghonh client with given backend handler.
func NewClient(handler ClientHandler) clientinterface.Client {
	return &client{packager: handler, transporter: handler}
}

func (mb *client) FunctionCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeFunctionCheck,
		FunctionCode: protocol.FuncCodeFunctionCheck,
		CommandType:  "remote",
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeStatusCheck,
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

func (mb *client) On(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.ON)
	commandsOn := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeOnOff,
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

func (mb *client) Off(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.OFF)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeOnOff,
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

func (mb *client) ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeErrorCheck,
		FunctionCode: protocol.FuncCodeErrorCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) FreshAirCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeFreshAirCheck,
		FunctionCode: protocol.FuncCodeFreshAirCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) NewAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	commands := data[2:]
	addressLen := dataBlockArray(newArr)
	newArr = PrependUint16(commands, protocol.ON)
	commandsOn := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeOnOff,
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

func (mb *client) NewAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.ON)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeOnOff,
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

func (mb *client) NewAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	address := data[:2]
	len_data := uint16(len(address) + 4)
	newArr := PrependUint16(address, len_data)
	addressLen := dataBlockArray(newArr)
	commands := data[2:]
	newArr = PrependUint16(commands, protocol.OFF)
	commandsOff := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeOnOff,
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

func (mb *client) NewAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	newArr := PrependUint16(data, len_data)
	addressLen := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeFreshAirErrorCheck,
		FunctionCode: protocol.FuncCodeFreshAirErrorCheck,
		Address:      addressLen,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) Control(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) NewAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b29 does not support following protocol")
}

func (mb *client) send(request *protocol.ProtocolDataUnit) (response *protocol.ProtocolDataUnit, err error) {
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
		err = responseError(response)
		return
	}
	if response.Data == nil || len(response.Data) == 0 {
		// Empty response
		err = fmt.Errorf("zhonghong: response data is empty")
		return
	}
	return
}

func dataBlock(value ...uint16) []byte {
	data := make([]byte, 2*len(value))
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return data
}

func dataBlockArray(arr []uint16) []byte {
	byteSlice := make([]byte, len(arr)*2)
	for i, v := range arr {
		binary.BigEndian.PutUint16(byteSlice[i*2:], v)
	}

	return byteSlice
}

func PrependUint16(slice []uint16, element uint16) []uint16 {
	newSlice := append([]uint16{element}, slice...)
	return newSlice
}

func responseError(response *protocol.ProtocolDataUnit) error {
	mbError := &protocol.ZhonghongError{FunctionCode: response.FunctionCode}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.Data[0]
	}
	return mbError
}
