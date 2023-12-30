package errs

import "github.com/TremblingV5/DouTok/pkg/response"

const (
	ParamsError      = "params error"
	ParamsErrorCode  = 101
	BindingError     = "binding error"
	BindingErrorCode = 102
)

var (
	Success    *response.Response
	ParamsErr  *response.Response
	BindingErr *response.Response
)

func Init(config response.Config) {
	Success = response.Success(config)

	ParamsErr = Success.Copy().Update(
		response.Message(ParamsError), response.Code(ParamsErrorCode),
	)
	BindingErr = Success.Copy().Update(
		response.Message(BindingError), response.Code(BindingErrorCode),
	)
}
