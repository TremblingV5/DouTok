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
