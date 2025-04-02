package api_server

import (
	"context"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stripe/stripe-go/v78/client"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

type MatchesServer struct {
	StytchClient *stytchapi.API
	StripeClient *client.API
	DB           *db.Queries
	Emailer      *mailgun.MailgunImpl
}

func (s *MatchesServer) PostMatches(ctx context.Context, request api.PostMatchesRequestObject) (api.PostMatchesResponseObject, error) {
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

func (s *MatchesServer) PutMatchesBatches(ctx context.Context, request api.PutMatchesBatchesRequestObject) (api.PutMatchesBatchesResponseObject, error) {
	return api.PutMatchesBatches200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PutMatchesBatches not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}

func (s *MatchesServer) PostMatchesUnassignPlayerFromMatch(ctx context.Context, request api.PostMatchesUnassignPlayerFromMatchRequestObject) (api.PostMatchesUnassignPlayerFromMatchResponseObject, error) {
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

func (s *MatchesServer) DeleteMatchesMatchId(ctx context.Context, request api.DeleteMatchesMatchIdRequestObject) (api.DeleteMatchesMatchIdResponseObject, error) {
	// TODO: Implement the logic to delete the match from the database.
	// For now, return a successful response.
	return api.DeleteMatchesMatchId200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *MatchesServer) PutMatchesMatchId(ctx context.Context, request api.PutMatchesMatchIdRequestObject) (api.PutMatchesMatchIdResponseObject, error) {
	return api.PutMatchesMatchId200JSONResponse(api.ApiResult{
		Error: &struct {
			Code    *string `json:"code,omitempty"`
			Message *string `json:"message,omitempty"`
		}{
			Code:    Ptr("NOT_IMPLEMENTED"),
			Message: Ptr("PutMatchesMatchId not implemented yet"),
		},
		IsSuccess: Ptr(false),
	}), nil
}
