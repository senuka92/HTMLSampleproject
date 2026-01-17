package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type Validator struct {
	jwks *keyfunc.JWKS
}

func NewValidator(ctx context.Context, jwksURL string) (*Validator, error) {
	if jwksURL == "" {
		return nil, fmt.Errorf("jwks url is required")
	}
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{RefreshInterval: 5 * time.Minute})
	if err != nil {
		return nil, fmt.Errorf("jwks init: %w", err)
	}
	return &Validator{jwks: jwks}, nil
}

func (v *Validator) ValidateRequest(r *http.Request) (*jwt.Token, error) {
	header := r.Header.Get("Authorization")
	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, fmt.Errorf("missing bearer token")
	}
	return jwt.Parse(parts[1], v.jwks.Keyfunc)
}
