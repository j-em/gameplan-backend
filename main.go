package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gameplan-backend/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/passwords"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/sessions"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

// MyApiServer provides a concrete implementation of the generated ServerInterface.
type MyApiServer struct {
	StytchClient *stytchapi.API
}

// Ensure MyApiServer implements the interface at compile time
var _ api.ServerInterface = (*MyApiServer)(nil)

// AuthMiddleware authenticates requests using the Authorization header.
func AuthMiddleware(stytchClient *stytchapi.API) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get(api.BearerAuthScopes) == nil {
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

// PostSessions - Example implementation for login
func (s *MyApiServer) PostSessions(ctx echo.Context) error {
	var params api.LoginUserParams
	if err := ctx.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid request format: %s", err))
	}

	fmt.Printf("Login attempt: Email=%s\n", params.Email)

	// Authenticate user with Stytch
	resp, err := s.StytchClient.Passwords.Authenticate(context.Background(), &passwords.AuthenticateParams{
		Email:                  params.Email,
		Password:               params.Password,
		SessionDurationMinutes: 60,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	responseData := map[string]interface{}{
		"message": "Login successful",
		"token":   resp.SessionToken,
	}
	apiResult := api.ApiResult{
		Data:      &responseData,
		IsSuccess: Ptr(true),
	}

	return ctx.JSON(http.StatusOK, apiResult)
}

// PostUsersSignUpUser - Example stub for user sign up
func (s *MyApiServer) PostUsersSignUpUser(ctx echo.Context) error {
	// --- TODO: Implement user sign up logic ---
	fmt.Println("Received request for POST /users/signUpUser")
	return ctx.JSON(http.StatusNotImplemented, api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{Code: Ptr("NOT_IMPLEMENTED"), Message: Ptr("Sign up not implemented yet")},
		IsSuccess: Ptr(false),
	})
}

// --- Add Stubs for all other methods defined in api.ServerInterface ---
// (Returning NotImplemented for brevity)

func (s *MyApiServer) PostMatches(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PutMatchesBatches(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostMatchesUnassignPlayerFromMatch(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeleteMatchesMatchId(ctx echo.Context, matchId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PutMatchesMatchId(ctx echo.Context, matchId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetPlayers(ctx echo.Context, params api.GetPlayersParams) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostPlayers(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeletePlayersPlayerId(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetPlayersPlayerId(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PutPlayersPlayerId(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetPlayersPlayerIdCustomColumns(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PutPlayersPlayerIdCustomColumns(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetPlayersPlayerIdSchedule(ctx echo.Context, playerId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasons(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostSeasons(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasonsTotalAmount(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeleteSeasonsSeasonId(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasonsSeasonId(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PutSeasonsSeasonId(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasonsSeasonIdPublicScheduleLink(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasonsSeasonIdScoreboard(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetSeasonsSeasonIdUpcoming(ctx echo.Context, seasonId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostSubscriptionsHandleSuccessUpgrade(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostSubscriptionsInitUpdatePaymentMethod(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostSubscriptionsUpgradeUserSubscription(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostSupportMessages(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersSendResetPasswordLink(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersSendVerificationEmail(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersUpdateUserPassword(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersVerifyMagicLinkToken(ctx echo.Context) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeleteUsersUserId(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetUsersUserIdAppsettings(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersUserIdAppsettings(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetUsersUserIdCustomPlayerColumns(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersUserIdCustomPlayerColumns(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeleteUsersUserIdCustomPlayerColumnsColumnId(ctx echo.Context, userId int, columnId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersUserIdResetCurrentUserPassword(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) DeleteUsersUserIdSubscription(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetUsersUserIdSubscription(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) GetUsersUserIdUsersettings(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}
func (s *MyApiServer) PostUsersUserIdUsersettings(ctx echo.Context, userId int) error {
	return echo.ErrNotImplemented
}

// Helper function to get pointers for optional fields in structs
func Ptr[T any](v T) *T {
	return &v
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Stytch client
	stytchProjectID := os.Getenv("STYTCH_PROJECT_ID")
	stytchSecret := os.Getenv("STYTCH_SECRET")
	if stytchProjectID == "" || stytchSecret == "" {
		panic("STYTCH_PROJECT_ID and STYTCH_SECRET environment variables must be set")
	}

	stytchClient, err := stytchapi.NewClient(
		stytchProjectID,
		stytchSecret,
	)
	if err != nil {
		panic(err)
	}

	// Create the API implementation
	myApi := &MyApiServer{StytchClient: stytchClient}

	// Authentication middleware
	e.Use(AuthMiddleware(stytchClient))

	// Register the handlers generated by oapi-codegen
	// This uses the ServerInterfaceWrapper internally
	api.RegisterHandlers(e, myApi)

	// Start server
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}

// --- Dummy types needed to satisfy interface if not using actual implementation ---
// (These might be needed if you haven't implemented all request/response bodies fully)
// type AddMatchParams = api.AddMatchParams
// type GetPlayersParams = api.GetPlayersParams
// ... and so on for all parameter/body types used in the interface methods
// Note: Since we are importing the 'api' package, these explicit type aliases
// are usually not necessary unless there's a naming conflict or specific reason.
// The stubs above directly use types like api.GetPlayersParams.

// --- Placeholder for openapi_types if needed ---
// (Usually provided by oapi-codegen/runtime)
// type Date = openapi_types.Date
