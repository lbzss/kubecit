// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.6
// source: helloworld/v1/greeter.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGreeterClusterList = "/helloworld.v1.Greeter/ClusterList"
const OperationGreeterNamespaceList = "/helloworld.v1.Greeter/NamespaceList"
const OperationGreeterSayHello = "/helloworld.v1.Greeter/SayHello"
const OperationGreeterUserList = "/helloworld.v1.Greeter/UserList"
const OperationGreeterUserRegister = "/helloworld.v1.Greeter/UserRegister"

type GreeterHTTPServer interface {
	ClusterList(context.Context, *Empty) (*ClusterListResponse, error)
	NamespaceList(context.Context, *NamespaceReq) (*NamespaceResp, error)
	// SayHello Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	UserList(context.Context, *Empty) (*UserListResponse, error)
	// UserRegister Register a user
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
}

func RegisterGreeterHTTPServer(s *http.Server, srv GreeterHTTPServer) {
	r := s.Route("/")
	r.GET("/helloworld/{name}", _Greeter_SayHello0_HTTP_Handler(srv))
	r.POST("/user/register", _Greeter_UserRegister0_HTTP_Handler(srv))
	r.GET("/user", _Greeter_UserList0_HTTP_Handler(srv))
	r.GET("/clusters", _Greeter_ClusterList0_HTTP_Handler(srv))
	r.GET("/namespaces/{cluster}", _Greeter_NamespaceList0_HTTP_Handler(srv))
}

func _Greeter_SayHello0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UserRegister0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRegister(ctx, req.(*UserRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UserList0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserList(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserListResponse)
		return ctx.Result(200, reply)
	}
}

func _Greeter_ClusterList0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterClusterList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ClusterList(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ClusterListResponse)
		return ctx.Result(200, reply)
	}
}

func _Greeter_NamespaceList0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NamespaceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterNamespaceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NamespaceList(ctx, req.(*NamespaceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NamespaceResp)
		return ctx.Result(200, reply)
	}
}

type GreeterHTTPClient interface {
	ClusterList(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *ClusterListResponse, err error)
	NamespaceList(ctx context.Context, req *NamespaceReq, opts ...http.CallOption) (rsp *NamespaceResp, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	UserList(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *UserListResponse, err error)
	UserRegister(ctx context.Context, req *UserRegisterRequest, opts ...http.CallOption) (rsp *UserRegisterResponse, err error)
}

type GreeterHTTPClientImpl struct {
	cc *http.Client
}

func NewGreeterHTTPClient(client *http.Client) GreeterHTTPClient {
	return &GreeterHTTPClientImpl{client}
}

func (c *GreeterHTTPClientImpl) ClusterList(ctx context.Context, in *Empty, opts ...http.CallOption) (*ClusterListResponse, error) {
	var out ClusterListResponse
	pattern := "/clusters"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterClusterList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) NamespaceList(ctx context.Context, in *NamespaceReq, opts ...http.CallOption) (*NamespaceResp, error) {
	var out NamespaceResp
	pattern := "/namespaces/{cluster}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterNamespaceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UserList(ctx context.Context, in *Empty, opts ...http.CallOption) (*UserListResponse, error) {
	var out UserListResponse
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...http.CallOption) (*UserRegisterResponse, error) {
	var out UserRegisterResponse
	pattern := "/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
