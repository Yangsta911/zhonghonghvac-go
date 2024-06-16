package zhonghongprotocol

type Client interface {
	// Bit access

	ReadGateway() (results []byte, err error)
}
