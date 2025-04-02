package api_server

import (
	"context"
	"errors"
	"fmt"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
	"github.com/stytchauth/stytch-go/stytch"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/passwords"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/users"
)

type AuthServer struct {
	StytchClient *stytchapi.API
	StripeClient *client.API
	DB           *db.Queries
	Emailer      *mailgun.MailgunImpl
}

func (s *AuthServer) PostSessions(ctx context.Context, request api.PostSessionsRequestObject) (api.PostSessionsResponseObject, error) {
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

	responseData := map[string]any{
		"message": "Login successful",
		"token":   resp.SessionToken,
	}

	return api.PostSessions200JSONResponse(api.ApiResult{
		Data:      &responseData,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *AuthServer) PostSupportMessages(ctx context.Context, request api.PostSupportMessagesRequestObject) (api.PostSupportMessagesResponseObject, error) {
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

func (s *AuthServer) DeleteUsersUserId(ctx context.Context, request api.DeleteUsersUserIdRequestObject) (api.DeleteUsersUserIdResponseObject, error) {
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

func (s *AuthServer) GetUsersUserIdAppsettings(ctx context.Context, request api.GetUsersUserIdAppsettingsRequestObject) (api.GetUsersUserIdAppsettingsResponseObject, error) {
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

func (s *AuthServer) PostUsersUserIdAppsettings(ctx context.Context, request api.PostUsersUserIdAppsettingsRequestObject) (api.PostUsersUserIdAppsettingsResponseObject, error) {
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

func (s *AuthServer) GetUsersUserIdCustomPlayerColumns(ctx context.Context, request api.GetUsersUserIdCustomPlayerColumnsRequestObject) (api.GetUsersUserIdCustomPlayerColumnsResponseObject, error) {
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

func (s *AuthServer) PostUsersUserIdCustomPlayerColumns(ctx context.Context, request api.PostUsersUserIdCustomPlayerColumnsRequestObject) (api.PostUsersUserIdCustomPlayerColumnsResponseObject, error) {
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

func (s *AuthServer) DeleteUsersUserIdCustomPlayerColumnsColumnId(ctx context.Context, request api.DeleteUsersUserIdCustomPlayerColumnsColumnIdRequestObject) (api.DeleteUsersUserIdCustomPlayerColumnsColumnIdResponseObject, error) {
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

func (s *AuthServer) PostUsersUserIdResetCurrentUserPassword(ctx context.Context, request api.PostUsersUserIdResetCurrentUserPasswordRequestObject) (api.PostUsersUserIdResetCurrentUserPasswordResponseObject, error) {
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

func (s *AuthServer) GetUsersUserIdUsersettings(ctx context.Context, request api.GetUsersUserIdUsersettingsRequestObject) (api.GetUsersUserIdUsersettingsResponseObject, error) {
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

func (s *AuthServer) PostUsersUserIdUsersettings(ctx context.Context, request api.PostUsersUserIdUsersettingsRequestObject) (api.PostUsersUserIdUsersettingsResponseObject, error) {
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

func (s *AuthServer) PostUsersSendResetPasswordLink(ctx context.Context, request api.PostUsersSendResetPasswordLinkRequestObject) (api.PostUsersSendResetPasswordLinkResponseObject, error) {
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

func (s *AuthServer) PostUsersSendVerificationEmail(ctx context.Context, request api.PostUsersSendVerificationEmailRequestObject) (api.PostUsersSendVerificationEmailResponseObject, error) {
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

func (s *AuthServer) PostUsersSignUpUser(ctx context.Context, request api.PostUsersSignUpUserRequestObject) (api.PostUsersSignUpUserResponseObject, error) {
	params := request.Body

	// Validate password length
	if len(params.Password) < 8 {
		return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("WEAK_PASSWORD"),
				Message: Ptr("Password must be at least 8 characters"),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Create Stytch user
	stytchResp, err := s.StytchClient.Passwords.Create(ctx, &passwords.CreateParams{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		var stytchErr *stytch.Error
		if errors.As(err, &stytchErr) {
			switch stytchErr.ErrorType {
			case "duplicate_email":
				return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
					Error: &struct {
						Code    *string `json:"code,omitempty"`
						Message *string `json:"message,omitempty"`
					}{
						Code:    Ptr("DUPLICATE_EMAIL"),
						Message: Ptr("Email address already exists"),
					},
					IsSuccess: Ptr(false),
				}), nil
			case "invalid_email":
				return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
					Error: &struct {
						Code    *string `json:"code,omitempty"`
						Message *string `json:"message,omitempty"`
					}{
						Code:    Ptr("INVALID_EMAIL"),
						Message: Ptr("Invalid email address format"),
					},
					IsSuccess: Ptr(false),
				}), nil
			}
		}
		return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("STYTCH_ERROR"),
				Message: Ptr("Failed to create user account"),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Create Stripe customer
	customer, err := s.StripeClient.Customers.New(&stripe.CustomerParams{
		Email: stripe.String(params.Email),
		Name:  stripe.String(params.Name),
		Params: stripe.Params{
			Metadata: map[string]string{
				"stytchId": stytchResp.UserID,
			},
		},
	})
	if err != nil {
		return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("STRIPE_ERROR"),
				Message: Ptr("Failed to create billing profile"),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Create database user
	birthday := pgtype.Int4{}
	if params.Birthday != nil {
		birthday.Int32 = int32(params.Birthday.Unix())
		birthday.Valid = true
		birthday.Valid = true
	}

	dbUser, err := s.DB.CreateUser(ctx, db.CreateUserParams{
		Stytchid:   stytchResp.UserID,
		Stripeid:   customer.ID,
		Name:       params.Name,
		Email:      params.Email,
		Phone:      pgtype.Text{Valid: false},
		Country:    pgtype.Text{Valid: false},
		Birthday:   birthday,
		Lang:       "en",
		Isverified: false,
	})
	if err != nil {
		return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DATABASE_ERROR"),
				Message: Ptr("Failed to save user information"),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Update Stytch metadata with our internal user ID
	_, err = s.StytchClient.Users.Update(ctx, &users.UpdateParams{
		UserID: stytchResp.UserID,
		TrustedMetadata: map[string]interface{}{
			"userId": dbUser.ID,
		},
	})
	if err != nil {
		// Log but continue since user is already created
		fmt.Printf("Failed to update Stytch metadata: %v\n", err)
	}

	// Send confirmation email
	_, _, err = s.Emailer.Send(ctx, s.Emailer.NewMessage(
		"no-reply@gameplan.com", // From (should be configured)
		"Welcome to Gameplan",   // Subject
		"",                      // Text body
		params.Email,            // To
	))
	if err != nil {
		// Log but continue since email is non-critical
		fmt.Printf("Failed to send confirmation email: %v\n", err)
	}

	return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *AuthServer) PostUsersUpdateUserPassword(ctx context.Context, request api.PostUsersUpdateUserPasswordRequestObject) (api.PostUsersUpdateUserPasswordResponseObject, error) {
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

func (s *AuthServer) PostUsersVerifyMagicLinkToken(ctx context.Context, request api.PostUsersVerifyMagicLinkTokenRequestObject) (api.PostUsersVerifyMagicLinkTokenResponseObject, error) {
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

// Helper function to get pointers for optional fields in structs
func Ptr[T any](v T) *T {
	return &v
}
