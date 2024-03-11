package errs

import "github.com/TremblingV5/DouTok/pkg/response"

const (
	bindingError     = "binding error"
	bindingErrorCode = 101
)

var (
	Success    *response.Response
	BindingErr *response.Response
)

func Init(config response.Config) {
	Success = response.Success(config)

	BindingErr = Success.Copy().Update(
		response.Message(bindingError), response.Code(bindingErrorCode),
	)
}

type Config struct {
}

func (c *Config) GetNodeCode() int32 {
	return 200
}

func (c *Config) GetNameCode() int32 {
	return 200
}