package b17protocol

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

// NewClient creates a new Zhonghong client with given backend handler.
func NewClient(handler ClientHandler) clientinterface.Client {
	return &client{packager: handler, transporter: handler}
}

func (mb *client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeReadGateway,
		FunctionCode: protocol.FuncCodeReadGateway,
		Data:         []byte{0x00, 0x00, 0x00, 0x00},
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	slice := []uint16{0, 0}
	newdata := append(slice, data...)
	senddata := dataBlockArray(newdata)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeReadGateway,
		FunctionCode: protocol.FuncCodeReadGateway,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) On(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.ON)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) Off(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.ON)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) Control(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayControl,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayWindSpeed,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayWindDir,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) NewAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.ON)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayNewAirOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) NewAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.OFF)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayNewAirOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) NewAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayNewAirMode,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) NewAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayNewAirSpeed,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *client) ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b17 does not support following protocol")
}

func (mb *client) FunctionCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b17 does not support following protocol")
}

func (mb *client) NewAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b17 does not support following protocol")
}

func (mb *client) StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong: b17 does not support following protocol")
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

func responseError(response *protocol.ProtocolDataUnit) error {
	mbError := &protocol.ZhonghongError{FunctionCode: response.FunctionCode}
	if response.Data != nil && len(response.Data) > 0 {
		mbError.ExceptionCode = response.Data[0]
	}
	return mbError
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
