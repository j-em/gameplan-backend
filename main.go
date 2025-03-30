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

// MyApiServer provides a concrete implementation of the generated StrictServerInterface.
type MyApiServer struct {
	StytchClient *stytchapi.API
}

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
func (s *MyApiServer) PostSessions(ctx context.Context, request api.PostSessionsRequestObject) (api.PostSessionsResponseObject, error) {
	params := request.Body

	fmt.Printf("Login attempt: Email=%s\n", params.Email)

	// Authenticate user with Stytch
	resp, err := s.StytchClient.Passwords.Authenticate(ctx, &passwords.AuthenticateParams{
		Email:                  params.Email,
		Password:               params.Password,
		SessionDurationMinutes: 60,
	})
	if err != nil {
		return api.PostSessions200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("INVALID_CREDENTIALS"),
				Message: Ptr("Invalid credentials"),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	responseData := map[string]interface{}{
		"message": "Login successful",
		"token":   resp.SessionToken,
	}

	return api.PostSessions200JSONResponse(api.ApiResult{
		Data:      &responseData,
		IsSuccess: Ptr(true),
	}), nil
}

// PostUsersSignUpUser - Example stub for user sign up
func (s *MyApiServer) PostUsersSignUpUser(ctx context.Context, request api.PostUsersSignUpUserRequestObject) (api.PostUsersSignUpUserResponseObject, error) {
	// --- TODO: Implement user sign up logic ---
	fmt.Println("Received request for POST /users/signUpUser")
	return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{Code: Ptr("NOT_IMPLEMENTED"), Message: Ptr("Sign up not implemented yet")},
		IsSuccess: Ptr(false),
	}), nil
}

// --- Add Stubs for all other methods defined in api.ServerInterface ---
// (Returning NotImplemented for brevity)

func (s *MyApiServer) PostMatches(ctx context.Context, request api.PostMatchesRequestObject) (api.PostMatchesResponseObject, error) {
	return api.PostMatches200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostMatches not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PutMatchesBatches(ctx context.Context, request api.PutMatchesBatchesRequestObject) (api.PutMatchesBatchesResponseObject, error) {
	return api.PutMatchesBatches200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Batch matches update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostMatchesUnassignPlayerFromMatch(ctx context.Context, request api.PostMatchesUnassignPlayerFromMatchRequestObject) (api.PostMatchesUnassignPlayerFromMatchResponseObject, error) {
	return api.PostMatchesUnassignPlayerFromMatch200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostMatchesUnassignPlayerFromMatch not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeleteMatchesMatchId(ctx context.Context, request api.DeleteMatchesMatchIdRequestObject) (api.DeleteMatchesMatchIdResponseObject, error) {
	return api.DeleteMatchesMatchId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeleteMatchesMatchId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PutMatchesMatchId(ctx context.Context, request api.PutMatchesMatchIdRequestObject) (api.PutMatchesMatchIdResponseObject, error) {
	return api.PutMatchesMatchId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Match update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetPlayers(ctx context.Context, request api.GetPlayersRequestObject) (api.GetPlayersResponseObject, error) {
	return api.GetPlayers200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetPlayers not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostPlayers(ctx context.Context, request api.PostPlayersRequestObject) (api.PostPlayersResponseObject, error) {
	return api.PostPlayers200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostPlayers not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeletePlayersPlayerId(ctx context.Context, request api.DeletePlayersPlayerIdRequestObject) (api.DeletePlayersPlayerIdResponseObject, error) {
	return api.DeletePlayersPlayerId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeletePlayersPlayerId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetPlayersPlayerId(ctx context.Context, request api.GetPlayersPlayerIdRequestObject) (api.GetPlayersPlayerIdResponseObject, error) {
	return api.GetPlayersPlayerId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetPlayersPlayerId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PutPlayersPlayerId(ctx context.Context, request api.PutPlayersPlayerIdRequestObject) (api.PutPlayersPlayerIdResponseObject, error) {
	return api.PutPlayersPlayerId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Player update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetPlayersPlayerIdCustomColumns(ctx context.Context, request api.GetPlayersPlayerIdCustomColumnsRequestObject) (api.GetPlayersPlayerIdCustomColumnsResponseObject, error) {
	return api.GetPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetPlayersPlayerIdCustomColumns not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PutPlayersPlayerIdCustomColumns(ctx context.Context, request api.PutPlayersPlayerIdCustomColumnsRequestObject) (api.PutPlayersPlayerIdCustomColumnsResponseObject, error) {
	return api.PutPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Player custom columns update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (极s *MyApiServer) GetPlayersPlayerIdSchedule(ctx context.Context, request api.GetPlayersPlayerIdScheduleRequestObject) (api.GetPlayersPlayerIdScheduleResponseObject, error) {
	return api.GetPlayersPlayerIdSchedule200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT极IMPLEMENTED"),
			Message: Ptr("GetPlayersPlayerIdSchedule not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasons(ctx context.Context, request api.GetSeasonsRequestObject) (api.GetSeasonsResponseObject, error) {
	return api.GetSeasons200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasons not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostSeasons(ctx context.Context, request api.PostSeasonsRequestObject) (api.PostSeasonsResponseObject, error) {
	return api.PostSeasons200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSeasons not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasonsTotalAmount(ctx context.Context, request api.GetSeasonsTotalAmountRequestObject) (api.GetSeasonsTotalAmountResponseObject, error) {
	return api.GetSeasonsTotalAmount200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasonsTotalAmount not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeleteSeasonsSeasonId(ctx context.Context, request api.DeleteSeasonsSeasonIdRequestObject) (api.DeleteSeasonsSeasonIdResponseObject, error) {
	return api.DeleteSeasonsSeasonId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeleteSeasonsSeasonId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasonsSeasonId(ctx context.Context, request api.GetSeasonsSeasonIdRequestObject) (api.GetSeasonsSeasonIdResponseObject, error) {
	return api.GetSeasonsSeasonId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasonsSeasonId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PutSeasonsSeasonId(ctx context.Context, request api.PutSeasonsSeasonIdRequestObject) (api.PutSeasonsSeasonIdResponseObject, error) {
	return api.PutSeasonsSeasonId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Season update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasonsSeasonIdPublicScheduleLink(ctx context.Context, request api.GetSeasonsSeasonIdPublicScheduleLinkRequestObject) (api.GetSeasonsSeasonIdPublicScheduleLinkResponseObject, error) {
	return api.GetSeasonsSeasonIdPublicScheduleLink200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasonsSeasonIdPublicScheduleLink not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasonsSeasonIdScoreboard(ctx context.Context, request api.GetSeasonsSeasonIdScoreboardRequestObject) (api.GetSeasonsSeasonIdScoreboardResponseObject, error) {
	return api.GetSeasonsSeasonIdScoreboard200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasonsSeasonIdScoreboard not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetSeasonsSeasonIdUpcoming(ctx context.Context, request api.GetSeasonsSeasonIdUpcomingRequestObject) (api.GetSeasonsSeasonIdUpcomingResponseObject, error) {
	return api.GetSeasonsSeasonIdUpcoming200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetSeasonsSeasonIdUpcoming not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostSubscriptionsHandleSuccessUpgrade(ctx context.Context, request api.PostSubscriptionsHandleSuccessUpgradeRequestObject) (api.PostSubscriptionsHandleSuccessUpgradeResponseObject, error) {
	return api.PostSubscriptionsHandleSuccessUpgrade200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSubscriptionsHandleSuccessUpgrade not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostSubscriptionsInitUpdatePaymentMethod(ctx context.Context, request api.PostSubscriptionsInitUpdatePaymentMethodRequestObject) (api.PostSubscriptionsInitUpdatePaymentMethodResponseObject, error) {
	return api.PostSubscriptionsInitUpdatePaymentMethod200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSubscriptionsInitUpdatePaymentMethod not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostSubscriptionsUpgradeUserSubscription(ctx context.Context, request api.PostSubscriptionsUpgradeUserSubscriptionRequestObject) (api.PostSubscriptionsUpgradeUserSubscriptionResponseObject, error) {
	return api.PostSubscriptionsUpgradeUserSubscription200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSubscriptionsUpgradeUserSubscription not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostSupportMessages(ctx context.Context, request api.PostSupportMessagesRequestObject) (api.PostSupportMessagesResponseObject, error) {
	return api.PostSupportMessages200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSupportMessages not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersSendResetPasswordLink(ctx context.Context, request api.PostUsersSendResetPasswordLinkRequestObject) (api.PostUsersSendResetPasswordLinkResponseObject, error) {
	return api.PostUsersSendResetPasswordLink200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostUsersSendResetPasswordLink not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersSendVerificationEmail(ctx context.Context, request api.PostUsersSendVerificationEmailRequestObject) (api.PostUsersSendVerificationEmailResponseObject, error) {
	return api.PostUsersSendVerificationEmail200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostUsersSendVerificationEmail not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersUpdateUserPassword(ctx context.Context, request api.PostUsersUpdateUserPasswordRequestObject) (api.PostUsersUpdateUserPasswordResponseObject, error) {
	return api.PostUsersUpdateUserPassword200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Password update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersVerifyMagicLinkToken(ctx context.Context, request api.PostUsersVerifyMagicLinkTokenRequestObject) (api.PostUsersVerifyMagicLinkTokenResponseObject, error) {
	return api.PostUsersVerifyMagicLinkToken200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Magic link verification not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeleteUsersUserId(ctx context.Context, request api.DeleteUsersUserIdRequestObject) (api.DeleteUsersUserIdResponseObject, error) {
	return api.DeleteUsersUserId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeleteUsersUserId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetUsersUserIdAppsettings(ctx context.Context, request api.GetUsersUserIdAppsettingsRequestObject) (api.GetUsersUserIdAppsettingsResponseObject, error) {
	return api.GetUsersUserIdAppsettings200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetUsersUserIdAppsettings not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersUserIdAppsettings(ctx context.Context, request api.PostUsersUserIdAppsettingsRequestObject) (api.PostUsersUserIdAppsettingsResponseObject, error) {
	return api.PostUsersUserIdAppsettings200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("App settings update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetUsersUserIdCustomPlayerColumns(ctx context.Context, request api.GetUsersUserIdCustomPlayerColumnsRequestObject) (api.GetUsersUserIdCustomPlayerColumnsResponseObject, error) {
	return api.GetUsersUserIdCustomPlayerColumns200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetUsersUserIdCustomPlayerColumns not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersUserIdCustomPlayerColumns(ctx context.Context, request api.PostUsersUserIdCustomPlayerColumnsRequestObject) (api.PostUsersUserIdCustomPlayerColumnsResponseObject, error) {
	return api.PostUsersUserIdCustomPlayerColumns200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Custom player columns not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeleteUsersUserIdCustomPlayerColumnsColumnId(ctx context.Context, request api.DeleteUsersUserIdCustomPlayerColumnsColumnIdRequestObject) (api.DeleteUsersUserIdCustomPlayerColumnsColumnIdResponseObject, error) {
	return api.DeleteUsersUserIdCustomPlayerColumnsColumnId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeleteUsersUserIdCustomPlayerColumnsColumnId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersUserIdResetCurrentUserPassword(ctx context.Context, request api.PostUsersUserIdResetCurrentUserPasswordRequestObject) (api.PostUsersUserIdResetCurrentUserPasswordResponseObject, error) {
	return api.PostUsersUserIdResetCurrentUserPassword200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("Password reset not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) DeleteUsersUserIdSubscription(ctx context.Context, request api.DeleteUsersUserIdSubscriptionRequestObject) (api.DeleteUsersUserIdSubscriptionResponseObject, error) {
	return api.DeleteUsersUserIdSubscription200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("DeleteUsersUserIdSubscription not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetUsersUserIdSubscription(ctx context.Context, request api.GetUsersUserIdSubscriptionRequestObject) (api.GetUsersUserIdSubscriptionResponseObject, error) {
	return api.GetUsersUserIdSubscription200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetUsersUserIdSubscription not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) GetUsersUserIdUsersettings(ctx context.Context, request api.GetUsersUserIdUsersettingsRequestObject) (api.GetUsersUserIdUsersettingsResponseObject, error) {
	return api.GetUsersUserIdUsersettings200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("GetUsersUserIdUsersettings not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
func (s *MyApiServer) PostUsersUserIdUsersettings(ctx context.Context, request api.PostUsersUserIdUsersettingsRequestObject) (api.PostUsersUserIdUsersettingsResponseObject, error) {
	return api.PostUsersUserIdUsersettings200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("User settings update not implemented"),
		},
		IsSuccess: Ptr(false),
	}), nil
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

	// Register the strict handlers generated by oapi-codegen
	strictHandler := api.NewStrictHandler(myApi, nil)
	api.RegisterHandlers(e, strictHandler)

	// Start server
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
