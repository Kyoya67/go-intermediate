package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Kyoya67/go-intermediate/apperrors"
	"github.com/Kyoya67/go-intermediate/common"
	"google.golang.org/api/idtoken"
)

const (
	googleClientId = "482610659858-imfs79kl9cssqdc9a46c8ds2j0uane8f.apps.googleusercontent.com"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(nil, "authorization header is required")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]

		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(nil, "bearer token is missing or id token is empty")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err := apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientId)
		if err != nil {
			err := apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err := apperrors.Unauthorizated.Wrap(nil, "name claim is missing")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))
		next.ServeHTTP(w, req)
	})
}
