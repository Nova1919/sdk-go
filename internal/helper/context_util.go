package helper

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserContext struct {
	UserId      string `json:"userId"`
	TeamId      string `json:"teamId"`
	AccessToken string `json:"accessToken"`
}

func (userCtx *UserContext) GetUserId() string {
	return userCtx.UserId
}

func (userCtx *UserContext) GetTeamId() string {
	return userCtx.TeamId
}

const UserContextKey = "user-context"

func WithUserContext(ctx context.Context, userCtx *UserContext) context.Context {
	return context.WithValue(ctx, UserContextKey, userCtx)
}

func FromContext(ctx context.Context) (*UserContext, error) {
	userCtx, ok := ctx.Value(UserContextKey).(*UserContext)
	if !ok {
		return nil, fmt.Errorf("no user found in context")
	}
	return userCtx, nil
}

func EncodeUserContext(userCtx *UserContext) (string, error) {
	jsonBytes, err := json.Marshal(userCtx)
	if err != nil {
		log.Errorf("Error marshaling user context: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

func DecodeUserContext(encodedValue string) (*UserContext, error) {
	jsonBytes, err := base64.StdEncoding.DecodeString(encodedValue)
	if err != nil {
		log.Errorf("Error decoding user context: %v", err)
		return nil, err
	}

	var userCtx UserContext
	if err := json.Unmarshal(jsonBytes, &userCtx); err != nil {
		log.Errorf("Error unmarshaling user context: %v", err)
		return nil, err
	}
	return &userCtx, nil
}

func ExtractFromIncoming(md metadata.MD) (*UserContext, bool) {
	values := md.Get(UserContextKey)
	if len(values) == 0 {
		return nil, false
	}

	decoded, err := DecodeUserContext(values[0])
	if err != nil {
		return nil, false
	}
	return decoded, true
}

func ClientContextInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		defer func() {
			err = invoker(ctx, method, req, reply, cc, opts...)
		}()

		userContext, err := FromContext(ctx)
		if err != nil {
			return
		}

		encodedValue, err := EncodeUserContext(userContext)
		if err != nil {
			log.Errorf("[Client Interceptor] Failed to encode UserContext: %v", err)
			return
		}

		ctx = metadata.AppendToOutgoingContext(ctx, UserContextKey, encodedValue)
		log.Errorf("[Client Interceptor] Added UserContext to metadata, UserId: %s,TeamId:%s, method: %s", userContext.UserId, userContext.TeamId, method)
		return
	}
}

func ServerContextInterceptor(requireAuth bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			if requireAuth {
				return nil, status.Error(codes.Unauthenticated, "missing metadata")
			}
			return handler(ctx, req)
		}
		userContext, ok := ExtractFromIncoming(md)
		if requireAuth && (!ok || userContext == nil) {
			log.Errorf("[Server Interceptor] Missing required UserId for method: %s", info.FullMethod)
			return nil, status.Error(codes.Unauthenticated, "missing userContext")
		}

		if ok {
			log.Infof("[Server Interceptor] Received request with UserId: %s, TeamId: %s, method: %s", userContext.UserId, userContext.TeamId, info.FullMethod)
			ctx = WithUserContext(ctx, userContext)
		} else {
			log.Infof("[Server Interceptor] Received request without valid UserContext, method: %s", info.FullMethod)
		}
		return handler(ctx, req)
	}
}
