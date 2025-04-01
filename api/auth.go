package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

// AuthMiddleware authenticates requests using the Authorization header.
func AuthMiddleware(stytchClient *stytchapi.API) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get(BearerAuthScopes) == nil {
				// Public route, skip authentication
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			token := authHeader[len("Bearer "):]
			session, err := stytchClient.Sessions.Authenticate(context.Background(), &sessions.AuthenticateParams{
				SessionToken:           token,
				SessionDurationMinutes: 60,
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid session token")
			}

			c.Set("stytch_session", session)
			c.Set("stytch_client", stytchClient) // Store stytchClient in context
			return next(c)
		}
	}
}
