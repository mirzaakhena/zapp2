package local

type localGateway struct {
}

// NewLocalGateway ...
func NewLocalGateway() (*localGateway, error) {
	return &localGateway{}, nil
}
