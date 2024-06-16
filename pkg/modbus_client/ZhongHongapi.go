package modbusclient

type Client interface {
	// Bit access

	ReadGateway() (results []byte, err error)
}
