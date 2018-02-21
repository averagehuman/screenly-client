package client

const (
	defaultEndpoint = "http://127.0.0.1:8080/api/v1.1/"
	defaultTimeout  = 30
)

type Config struct {
	ENDPOINT string
	TIMEOUT  uint
}

var defaultConfig = Config{
	ENDPOINT: defaultEndpoint,
	TIMEOUT:  defaultTimeout,
}
