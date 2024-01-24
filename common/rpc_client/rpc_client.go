package rpc_client

import (
	"context"
	"github.com/chinabobo/learning-go-zero/common/constant"
	"github.com/chinabobo/learning-go-zero/common/jsonx"
	"github.com/chinabobo/learning-go-zero/common/xerr"
	"github.com/chinabobo/learning-go-zero/common/xerr/code"
	"github.com/chinabobo/learning-go-zero/common/xerr/msg"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

func NewRPCClient(c zrpc.RpcClientConf) zrpc.Client {
	return zrpc.MustNewClient(c, zrpc.WithTimeout(2*time.Second),
		zrpc.WithUnaryClientInterceptor(requestIDClientInterceptor()),
		zrpc.WithUnaryClientInterceptor(extractErrorClientInterceptor()))
}

func NewRPCClientWithTimeout(c zrpc.RpcClientConf, timeout time.Duration) zrpc.Client {
	return zrpc.MustNewClient(c, zrpc.WithTimeout(timeout),
		zrpc.WithUnaryClientInterceptor(requestIDClientInterceptor()),
		zrpc.WithUnaryClientInterceptor(extractErrorClientInterceptor()))
}

func requestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value(constant.RequestIDKey)
		if requestID, ok := value.(string); ok && requestID != "" {
			md[constant.RequestIDKey] = []string{requestID}
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
	}
}

func extractErrorClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value(constant.RequestIDKey)
		if requestID, ok := value.(string); ok && requestID != "" {
			md[constant.RequestIDKey] = []string{requestID}
		}

		err = invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
		if err == nil {
			return nil
		}
		errCode := code.CODE_ERROR
		errMsg := msg.GetMsg(errCode)
		s, ok := status.FromError(err)
		if ok && (s.Code() == codes.OK || s.Code() == codes.Unknown) {
			rpcErr := xerr.CodeError{}
			err = jsonx.Unmarshal([]byte(s.Message()), &rpcErr)
			if err == nil {
				errCode = rpcErr.GetCode()
				errMsg = rpcErr.GetMsg()
			}
		} else {
			var e *xerr.CodeError
			if errors.As(err, &e) {
				causeErr := xerr.Cause(err) // err类型
				errCode, errMsg = e.GetCode(), causeErr.Error()
			}
		}

		return xerr.New(errCode, errMsg)
	}
}
