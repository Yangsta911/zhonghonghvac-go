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

// ReadGateway returns data related to the gateway
func (mb *b19client) ReadGateway() (results *protocol.ProtocolDataUnit, err error) {
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

// EditGateway edits the data related to the gateway
func (mb *b19client) EditGateway(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	slice := []uint16{0, 0}
	newdata := append(slice, data...)
	senddata := dataBlockArray(newdata)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeReadGateway,
		FunctionCode: protocol.FuncCodeEditGateway,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// On sends the ON command for HVAC to gateway
func (mb *b19client) On(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// Off sends the OFF command for HVAC to gateway
func (mb *b19client) Off(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// TempControl sends the temperature control command for HVAC to gateway
func (mb *b19client) TempControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// Control sends the control mode command for HVAC to gateway
func (mb *b19client) Control(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// WindSpeedControl sends the wind speed control command for HVAC to gateway
func (mb *b19client) WindSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// WindDirControl sends the wind direction control command for HVAC to gateway
func (mb *b19client) WindDirControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// FreshAirOn sends the ON command for Fresh Air ventilation unit to gateway
func (mb *b19client) FreshAirOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// FreshAirOff sends the OFF command for Fresh Air ventilation unit to gateway
func (mb *b19client) FreshAirOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// FreshAirModeControl sends the mode control command for Fresh Air ventilation unit to gateway
func (mb *b19client) FreshAirModeControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// FreshAirSpeedControl sends the speed control command for Fresh Air ventilation unit to gateway
func (mb *b19client) FreshAirSpeedControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
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

// FloorHeatingOn sends the ON command for Floor Heating unit to gateway
func (mb *b19client) FloorHeatingOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.ON)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.FuncCodeGatewayFloorHeatingOnOff,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingOff sends the OFF command for Floor Heating unit to gateway
func (mb *b19client) FloorHeatingOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.OFF)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.FuncCodeGatewayFloorHeatingOnOff,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingTemp sends the temperature control command for Floor Heating unit to gateway
func (mb *b19client) FloorHeatingTemp(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingTemp,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingControl sends the control command for Floor Heating unit to gateway
func (mb *b19client) FloorHeatingControl(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	senddata := dataBlockArray(datalenarr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingControl,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingAntiFreezeOn sends the ON command for Floor Heating unit AntiFreeze function to gateway
func (mb *b19client) FloorHeatingAntiFreezeOn(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.OFF)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingAntiFreezeOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// FloorHeatingAntiFreezeOff sends the OFF command for Floor Heating unit AntiFreeze function to gateway
func (mb *b19client) FloorHeatingAntiFreezeOff(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	len_data := uint16(len(data) + 4)
	datalenarr := PrependUint16(data, len_data)
	newArr := PrependUint16(datalenarr, protocol.OFF)
	senddata := dataBlockArray(newArr)
	request := protocol.ProtocolDataUnit{
		Header:       protocol.HeadCodeGateway,
		FunctionCode: protocol.FuncCodeGatewayFloorHeatingAntiFreezeOnOff,
		Data:         senddata,
	}
	resp, err := mb.send(&request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (mb *b19client) ErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
}

func (mb *b19client) PerformanceCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
}

func (mb *b19client) FreshAirErrorCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
}

func (mb *b19client) StatusCheck(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
}

func (mb *b19client) FreshAirPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")

}

func (mb *b19client) FreshAirStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")

}

func (mb *b19client) FloorHeatingPerformance(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
}

func (mb *b19client) FloorHeatingStatus(data []uint16) (results *protocol.ProtocolDataUnit, err error) {
	return nil, fmt.Errorf("zhonghong-b19 client: does not support following protocol")
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
