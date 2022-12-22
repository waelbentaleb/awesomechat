package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/waelbentaleb/awesomechat/domain/user"
)

type AuthInterceptor struct {
	service *user.Service
}

func NewAuthInterceptor(service *user.Service) *AuthInterceptor {
	return &AuthInterceptor{
		service: service,
	}
}

// UnaryAuthInterceptor intercept and authorize grpc unary requests
// In case when the request is valid it injects the current user to context
func (interceptor *AuthInterceptor) UnaryAuthInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(
		func(
			ctx context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler,
		) (interface{}, error) {

			// Ignore create user method
			if info.FullMethod == "/ChatCore/CreateUser" {
				h, err := handler(ctx, req)
				return h, err
			}

			validUser, err := interceptor.authorize(ctx)
			if err != nil {
				return nil, err
			}

			// Add the current user to context
			ctx = context.WithValue(ctx, "user", validUser)

			h, err := handler(ctx, req)
			return h, err
		},
	)
}

// StreamAuthInterceptor intercept and authorize grpc stream requests
// In case when the request is valid it injects the current user to context
// This requires a custom server stream object that is described bellow
func (interceptor *AuthInterceptor) StreamAuthInterceptor() grpc.ServerOption {
	return grpc.StreamInterceptor(
		func(
			srv interface{},
			ss grpc.ServerStream,
			info *grpc.StreamServerInfo,
			handler grpc.StreamHandler,
		) error {

			// Ignore reflections requests
			if info.FullMethod == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo" {
				return handler(srv, ss)
			}

			ctx := ss.Context()
			validUser, err := interceptor.authorize(ctx)
			if err != nil {
				return err
			}

			// Add the current user to context
			ctx = context.WithValue(ctx, "user", validUser)

			return handler(srv, &serverStream{
				ServerStream: ss,
				ctx:          ctx,
			})
		},
	)
}

// authorize function authorizes the token received from Metadata
func (interceptor *AuthInterceptor) authorize(ctx context.Context) (*user.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("retrieving metadata is failed")
	}

	// Extract authorization token from metadata
	authHeader, ok := md["authorization"]
	if !ok {
		return nil, errors.New("authorization token is not supplied")
	}

	// Validate the given token
	token := authHeader[0]
	validUser, err := interceptor.service.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	if validUser.Username == "" {
		return nil, errors.New("invalid associated username")
	}

	return validUser, nil
}

// serverStream is used as a wrapper to update context in the stream interceptor
type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

// Context represent an override of Context() and it's used to return a custom context
func (s *serverStream) Context() context.Context {
	return s.ctx
}
