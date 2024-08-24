package client

type ConfClient struct {
	addresses           []string
	healthcheckEndpoint string
}

func NewConfClient(addresses []string, healthcheckEndpoint string) ConfClient {
	return ConfClient{addresses: addresses, healthcheckEndpoint: healthcheckEndpoint}
}
