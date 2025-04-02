package api_server

import (
	"context"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stripe/stripe-go/v78/client"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

type AuthServer struct {
	StytchClient *stytchapi.API
	StripeClient *client.API
	DB           *db.Queries
	Emailer      *mailgun.MailgunImpl
}

func (s *AuthServer) PostSessions(ctx context.Context, request api.PostSessionsRequestObject) (api.PostSessionsResponseObject, error) {
	return api.PostSessions200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostSessions not implemented yet"),
		},
		IsSuccess: Ptr(false),
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
	return api.PostUsersSignUpUser200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PostUsersSignUpUser not implemented yet"),
		},
		IsSuccess: Ptr(false),
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
