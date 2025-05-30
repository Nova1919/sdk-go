package helper

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

func GrpcDialCredentials(ctx context.Context, host string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	var tlsEnable = false
	if strings.Contains(host, "tls://") {
		tlsEnable = true
	}
	host = strings.ReplaceAll(host, "tls://", "")
	if tlsEnable {
		cer := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		opts = append(opts, grpc.WithTransportCredentials(cer))
		return grpc.DialContext(ctx, host, grpc.WithTransportCredentials(cer))
	}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, host, opts...)
}
