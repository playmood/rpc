package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderAccessKey = "client-id"
	ClientHeaderSecretKey = "client-secret"
)

func NewClientCredential(ak, sk string) metadata.MD {
	return metadata.MD{
		ClientHeaderAccessKey: []string{ak},
		ClientHeaderSecretKey: []string{sk},
	}
}

func NewGrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return NewGrpcAuther().UnaryServerInterceptor
}

func NewGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

type grpcAuther struct {
}

// request response 认证
func (a *grpcAuther) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 1. 读取凭证， 放在meta信息中，[http2 header]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验合法性
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func (a *grpcAuther) GetClientCredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cidList := md[ClientHeaderAccessKey]
	if len(cidList) > 0 {
		clientId = cidList[0]
	}
	cskList := md[ClientHeaderSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}
	return
}

func (a *grpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if !(clientId == "admin" && clientSecret == "1234567890") {
		// 返回认证错误 直接结束rpc调用
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret not correct")
	}
	return nil
}

func NewGrpcAuthStreamServerInterceptor() grpc.StreamServerInterceptor {
	return NewGrpcAuther().StreamServerInterceptor
}

// server stream 认证
func (a *grpcAuther) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return fmt.Errorf("ctx is not an grpc incoming context")
	}

	clientId, clientSecret := a.GetClientCredentialsFromMeta(md)

	// 校验合法性
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return err
	}

	return handler(srv, ss)
}
