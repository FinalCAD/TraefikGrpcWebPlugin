package TraefikGrpcWebPlugin

import (
	"context"
	"net/http"

	"github.com/traefik/grpc-web/go/grpcweb"
)

type Config struct {
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		AllowOrigins: []string{"*"},
	}
}

func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {

	return grpcweb.WrapHandler(next, grpcweb.WithCorsForRegisteredEndpointsOnly(false), grpcweb.WithOriginFunc(func(origin string) bool {
		for _, originCfg := range config.AllowOrigins {
			if originCfg == "*" || originCfg == origin {
				return true
			}
		}
		return false
	})), nil
}
