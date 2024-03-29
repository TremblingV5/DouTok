// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userdomainservice

import (
	"context"
	"fmt"
	userDomain "github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userDomainServiceServiceInfo
}

var userDomainServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserDomainService"
	handlerType := (*userDomain.UserDomainService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AddUser":     kitex.NewMethodInfo(addUserHandler, newAddUserArgs, newAddUserResult, false),
		"CheckUser":   kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
		"GetUserInfo": kitex.NewMethodInfo(getUserInfoHandler, newGetUserInfoArgs, newGetUserInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "userDomain",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func addUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userDomain.DoutokAddUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userDomain.UserDomainService).AddUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AddUserArgs:
		success, err := handler.(userDomain.UserDomainService).AddUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddUserResult)
		realResult.Success = success
	}
	return nil
}
func newAddUserArgs() interface{} {
	return &AddUserArgs{}
}

func newAddUserResult() interface{} {
	return &AddUserResult{}
}

type AddUserArgs struct {
	Req *userDomain.DoutokAddUserRequest
}

func (p *AddUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userDomain.DoutokAddUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *AddUserArgs) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokAddUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddUserArgs_Req_DEFAULT *userDomain.DoutokAddUserRequest

func (p *AddUserArgs) GetReq() *userDomain.DoutokAddUserRequest {
	if !p.IsSetReq() {
		return AddUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddUserResult struct {
	Success *userDomain.DoutokAddUserResponse
}

var AddUserResult_Success_DEFAULT *userDomain.DoutokAddUserResponse

func (p *AddUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userDomain.DoutokAddUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *AddUserResult) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokAddUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddUserResult) GetSuccess() *userDomain.DoutokAddUserResponse {
	if !p.IsSetSuccess() {
		return AddUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userDomain.DoutokAddUserResponse)
}

func (p *AddUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userDomain.DoutokCheckUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userDomain.UserDomainService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(userDomain.UserDomainService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *userDomain.DoutokCheckUserRequest
}

func (p *CheckUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userDomain.DoutokCheckUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokCheckUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *userDomain.DoutokCheckUserRequest

func (p *CheckUserArgs) GetReq() *userDomain.DoutokCheckUserRequest {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *userDomain.DoutokCheckUserResponse
}

var CheckUserResult_Success_DEFAULT *userDomain.DoutokCheckUserResponse

func (p *CheckUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userDomain.DoutokCheckUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokCheckUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *userDomain.DoutokCheckUserResponse {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userDomain.DoutokCheckUserResponse)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userDomain.DoutokGetUserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userDomain.UserDomainService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserInfoArgs:
		success, err := handler.(userDomain.UserDomainService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *userDomain.DoutokGetUserInfoRequest
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(userDomain.DoutokGetUserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokGetUserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *userDomain.DoutokGetUserInfoRequest

func (p *GetUserInfoArgs) GetReq() *userDomain.DoutokGetUserInfoRequest {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserInfoResult struct {
	Success *userDomain.DoutokGetUserInfoResponse
}

var GetUserInfoResult_Success_DEFAULT *userDomain.DoutokGetUserInfoResponse

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(userDomain.DoutokGetUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(userDomain.DoutokGetUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *userDomain.DoutokGetUserInfoResponse {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*userDomain.DoutokGetUserInfoResponse)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddUser(ctx context.Context, Req *userDomain.DoutokAddUserRequest) (r *userDomain.DoutokAddUserResponse, err error) {
	var _args AddUserArgs
	_args.Req = Req
	var _result AddUserResult
	if err = p.c.Call(ctx, "AddUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *userDomain.DoutokCheckUserRequest) (r *userDomain.DoutokCheckUserResponse, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *userDomain.DoutokGetUserInfoRequest) (r *userDomain.DoutokGetUserInfoResponse, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
