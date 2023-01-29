package pack
//响应
// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *notedemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *notedemo.BaseResp {
	return &notedemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
