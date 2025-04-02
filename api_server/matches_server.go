package api_server

import (
	"context"
	"fmt"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stripe/stripe-go/v81/client"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

type MatchesServer struct {
	StytchClient *stytchapi.API
	StripeClient *client.API
	DB           *db.Queries
	Emailer      *mailgun.MailgunImpl
}

func (s *MatchesServer) PostMatches(ctx context.Context, request api.PostMatchesRequestObject) (api.PostMatchesResponseObject, error) {
	// Prepare database parameters with proper pgtype conversions
	params := db.CreateMatchParams{
		Seasonid:        pgtype.Int4{Int32: int32(request.Body.SeasonId), Valid: true},
		Playerid1:       pgtype.Int4{Int32: int32(request.Body.PlayerId1), Valid: true},
		Playerid2:       pgtype.Int4{Int32: int32(request.Body.PlayerId2), Valid: true},
		Playerid1points: 0,
		Playerid2points: 0,
		Winnerid:        pgtype.Int4{Valid: false},
		Group:           int32(request.Body.Group),
		Matchdate:       pgtype.Date{Time: request.Body.MatchDate.Time, Valid: true},
	}

	// Create match in database
	_, err := s.DB.CreateMatch(ctx, params)
	if err != nil {
		return api.PostMatches200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to create match: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	return api.PostMatches200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *MatchesServer) PutMatchesBatches(ctx context.Context, request api.PutMatchesBatchesRequestObject) (api.PutMatchesBatchesResponseObject, error) {
	var updatedMatchesCount int32 = 0
	for _, match := range *request.Body {
		params := db.UpdateMatchParams{
			Seasonid:        pgtype.Int4{Int32: int32(*match.SeasonId), Valid: match.SeasonId != nil},
			Playerid1:       pgtype.Int4{Int32: int32(*match.PlayerId1), Valid: match.PlayerId1 != nil},
			Playerid2:       pgtype.Int4{Int32: int32(*match.PlayerId2), Valid: match.PlayerId2 != nil},
			Playerid1points: int32(*match.PlayerId1Points),
			Playerid2points: int32(*match.PlayerId2Points),
			Group:           int32(*match.Group),
			Matchdate:       pgtype.Date{Time: match.MatchDate.Time, Valid: match.MatchDate != nil},
			ID:              int32(*match.Id),
		}

		_, err := s.DB.UpdateMatch(ctx, params)
		if err != nil {
			fmt.Printf("Failed to update match: %v\n", err)
			continue
		}
		updatedMatchesCount++
	}

	return api.PutMatchesBatches200JSONResponse(api.ApiResult{
		Data: &map[string]interface{}{
			"updatedMatchesCount": updatedMatchesCount,
		},
		IsSuccess: Ptr(true),
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
