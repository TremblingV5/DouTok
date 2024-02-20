package response

import (
	"fmt"
	"strconv"
)

type Response struct {
	code         int32
	message      string
	detail       string
	nameCode     int32
	nodeCode     int32
	status       int32
	totalMessage string
}

type Option func(*Response)

type Config interface {
	GetNameCode() int32
	GetNodeCode() int32
}

func Success(config Config) *Response {
	return New(
		Code(0),
		NameCode(config.GetNameCode()),
		NodeCode(config.GetNodeCode()),
		Message("Success"),
	)
}

func New(options ...Option) *Response {
	response := &Response{}
	return response.Update(options...)
}

func (r *Response) Update(options ...Option) *Response {
	for _, option := range options {
		option(r)
	}

	if r.code == 0 {
		r.status = 0
	} else {
		code := fmt.Sprintf("%03d%03d%03d", r.nameCode, r.nodeCode, r.code)
		codeI32, _ := strconv.ParseInt(code, 10, 32)
		r.status = int32(codeI32)
	}

	if r.detail != "" {
		r.totalMessage = fmt.Sprintf("%s %s", r.message, r.detail)
	} else {
		r.totalMessage = r.message
	}

	return r
}

func (r *Response) Copy() *Response {
	return &Response{
		code:     r.code,
		message:  r.message,
		detail:   r.detail,
		nameCode: r.nameCode,
		nodeCode: r.nodeCode,
	}
}

func Code(code int32) Option {
	return func(r *Response) {
		r.code = code
	}
}

func Message(message string) Option {
	return func(r *Response) {
		r.message = message
	}
}

func Detail(detail string) Option {
	return func(r *Response) {
		r.detail = detail
	}
}

func NameCode(nameCode int32) Option {
	return func(r *Response) {
		r.nameCode = nameCode
	}
}

func NodeCode(nodeCode int32) Option {
	return func(r *Response) {
		r.nodeCode = nodeCode
	}
}

func (r *Response) Code() int32 {
	return r.status
}

func (r *Response) Message() string {
	return r.message
}

func (r *Response) Detail() string {
	return r.detail
}

func (r *Response) Total() string {
	return r.totalMessage
}
