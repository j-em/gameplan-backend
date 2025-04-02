package api_server

import (
	"context"
	"fmt"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// SeasonsServer handles season-related operations
type SeasonsServer struct {
	DB *db.Queries
}

// CreateSeason creates a new season record based on API params
func (s *SeasonsServer) CreateSeason(
	ctx context.Context,
	userId int32,
	name string,
	startDate pgtype.Date,
	seasonType string,
	frequency string,
) (*db.Season, error) {
	params := db.CreateSeasonParams{
		Userid:     pgtype.Int4{Int32: userId, Valid: true},
		Name:       name,
		Startdate:  startDate,
		Seasontype: seasonType,
		Frequency:  frequency,
	}

	season, err := s.DB.CreateSeason(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create season: %w", err)
	}
	return &season, nil
}

// GetSeason retrieves season details by ID with user auth check
func (s *SeasonsServer) GetSeason(
	ctx context.Context,
	userId int32,
	seasonId int32,
) (*db.Season, error) {
	season, err := s.DB.GetSeason(ctx, db.GetSeasonParams{
		ID:     seasonId,
		Userid: pgtype.Int4{Int32: userId, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get season: %w", err)
	}
	return &season, nil
}

// UpdateSeason updates season information via key-value pairs
func (s *SeasonsServer) UpdateSeason(
	ctx context.Context,
	userId int32,
	seasonId int32,
	updates map[string]interface{},
) (*db.Season, error) {
	// Get current season to preserve unchanged fields
	current, err := s.GetSeason(ctx, userId, seasonId)
	if err != nil {
		return nil, err
	}

	params := db.UpdateSeasonParams{
		ID:         seasonId,
		Userid:     pgtype.Int4{Int32: userId, Valid: true},
		Name:       current.Name,
		Startdate:  current.Startdate,
		Seasontype: current.Seasontype,
		Frequency:  current.Frequency,
		Isactive:   current.Isactive,
	}

	// Apply updates from the key-value map
	for key, value := range updates {
		switch key {
		case "name":
			params.Name = value.(string)
		case "startDate":
			params.Startdate = value.(pgtype.Date)
		case "seasonType":
			params.Seasontype = value.(string)
		case "frequency":
			params.Frequency = value.(string)
		case "isActive":
			params.Isactive = value.(bool)
		}
	}

	season, err := s.DB.UpdateSeason(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update season: %w", err)
	}
	return &season, nil
}

// DeleteSeason removes a season record with user auth check
func (s *SeasonsServer) DeleteSeason(
	ctx context.Context,
	userId int32,
	seasonId int32,
) error {
	if err := s.DB.DeleteSeason(ctx, db.DeleteSeasonParams{
		ID:     seasonId,
		Userid: pgtype.Int4{Int32: userId, Valid: true},
	}); err != nil {
		return fmt.Errorf("failed to delete season: %w", err)
	}
	return nil
}

// ListSeasons retrieves all seasons for a user
func (s *SeasonsServer) ListSeasons(
	ctx context.Context,
	userId int32,
) ([]db.Season, error) {
	seasons, err := s.DB.GetSeasons(ctx, pgtype.Int4{Int32: userId, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list seasons: %w", err)
	}
	return seasons, nil
}

// GetSeasonScoreboard retrieves the scoreboard for a season
func (s *SeasonsServer) GetSeasonScoreboard(
	ctx context.Context,
	userId int32,
	seasonId int32,
) ([]db.GetSeasonScoreboardRow, error) {
	scoreboard, err := s.DB.GetSeasonScoreboard(ctx, db.GetSeasonScoreboardParams{
		Seasonid: pgtype.Int4{Int32: seasonId, Valid: true},
		ID:       seasonId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get season scoreboard: %w", err)
	}
	return scoreboard, nil
}

// GetSeasonUpcomingMatches retrieves upcoming matches for a season
func (s *SeasonsServer) GetSeasonUpcomingMatches(
	ctx context.Context,
	seasonId int32,
) ([]db.Match, error) {
	matches, err := s.DB.GetSeasonUpcomingMatches(ctx, pgtype.Int4{Int32: seasonId, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get upcoming matches: %w", err)
	}
	return matches, nil
}

// API endpoint implementations

func (s *SeasonsServer) GetSeasons(ctx context.Context, request api.GetSeasonsRequestObject) (api.GetSeasonsResponseObject, error) {
	// Extract user ID from context (would come from JWT in middleware)
	userID := int32(1) // TODO: Get from auth context

	seasons, err := s.ListSeasons(ctx, userID)
	if err != nil {
		return api.GetSeasons200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to list seasons: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	seasonsMap := map[string]interface{}{
		"seasons": seasons,
	}
	return api.GetSeasons200JSONResponse(api.ApiResult{
		Data:      &seasonsMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) PostSeasons(ctx context.Context, request api.PostSeasonsRequestObject) (api.PostSeasonsResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	season, err := s.CreateSeason(
		ctx,
		userID,
		request.Body.Name,
		pgtype.Date{Time: request.Body.StartDate.Time, Valid: true},
		request.Body.SeasonType,
		"weekly", // Default frequency
	)
	if err != nil {
		return api.PostSeasons200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to create season: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	seasonMap := map[string]interface{}{
		"season": *season,
	}
	return api.PostSeasons200JSONResponse(api.ApiResult{
		Data:      &seasonMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) GetSeasonsTotalAmount(ctx context.Context, request api.GetSeasonsTotalAmountRequestObject) (api.GetSeasonsTotalAmountResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	seasons, err := s.ListSeasons(ctx, userID)
	if err != nil {
		return api.GetSeasonsTotalAmount200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to count seasons: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	countMap := map[string]interface{}{
		"totalSeasons": len(seasons),
	}
	return api.GetSeasonsTotalAmount200JSONResponse(api.ApiResult{
		Data:      &countMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) DeleteSeasonsSeasonId(ctx context.Context, request api.DeleteSeasonsSeasonIdRequestObject) (api.DeleteSeasonsSeasonIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	if err := s.DeleteSeason(ctx, userID, int32(request.SeasonId)); err != nil {
		return api.DeleteSeasonsSeasonId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to delete season: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	return api.DeleteSeasonsSeasonId200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) GetSeasonsSeasonId(ctx context.Context, request api.GetSeasonsSeasonIdRequestObject) (api.GetSeasonsSeasonIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	season, err := s.GetSeason(ctx, userID, int32(request.SeasonId))
	if err != nil {
		return api.GetSeasonsSeasonId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get season: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Get players for the season
	players, err := s.DB.GetPlayers(ctx, pgtype.Int4{Int32: userID, Valid: true})
	if err != nil {
		return api.GetSeasonsSeasonId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get players: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	// Get matches for the season
	matches, err := s.DB.GetSeasonUpcomingMatches(ctx, pgtype.Int4{Int32: int32(request.SeasonId), Valid: true})
	if err != nil {
		return api.GetSeasonsSeasonId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get matches: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	seasonDetails := map[string]interface{}{
		"season":     *season,
		"players":    players,
		"allMatches": matches,
	}
	return api.GetSeasonsSeasonId200JSONResponse(api.ApiResult{
		Data:      &seasonDetails,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) PutSeasonsSeasonId(ctx context.Context, request api.PutSeasonsSeasonIdRequestObject) (api.PutSeasonsSeasonIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	updates := map[string]interface{}{
		"name": request.Body.Name,
	}

	season, err := s.UpdateSeason(ctx, userID, int32(request.SeasonId), updates)
	if err != nil {
		return api.PutSeasonsSeasonId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to update season: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	seasonMap := map[string]interface{}{
		"season": *season,
	}
	return api.PutSeasonsSeasonId200JSONResponse(api.ApiResult{
		Data:      &seasonMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) GetSeasonsSeasonIdPublicScheduleLink(ctx context.Context, request api.GetSeasonsSeasonIdPublicScheduleLinkRequestObject) (api.GetSeasonsSeasonIdPublicScheduleLinkResponseObject, error) {
	// TODO: Implement public schedule link generation
	publicLink := fmt.Sprintf("/public/seasons/%d/schedule", request.SeasonId)
	publicLinkMap := map[string]interface{}{
		"publicLink": publicLink,
	}
	return api.GetSeasonsSeasonIdPublicScheduleLink200JSONResponse(api.ApiResult{
		Data:      &publicLinkMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) GetSeasonsSeasonIdScoreboard(ctx context.Context, request api.GetSeasonsSeasonIdScoreboardRequestObject) (api.GetSeasonsSeasonIdScoreboardResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	scoreboard, err := s.GetSeasonScoreboard(ctx, userID, int32(request.SeasonId))
	if err != nil {
		return api.GetSeasonsSeasonIdScoreboard200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get scoreboard: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	season, err := s.GetSeason(ctx, userID, int32(request.SeasonId))
	if err != nil {
		return api.GetSeasonsSeasonIdScoreboard200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get season name: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	scoreboardData := map[string]interface{}{
		"scoreboard": scoreboard,
		"seasonName": season.Name,
	}
	return api.GetSeasonsSeasonIdScoreboard200JSONResponse(api.ApiResult{
		Data:      &scoreboardData,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *SeasonsServer) GetSeasonsSeasonIdUpcoming(ctx context.Context, request api.GetSeasonsSeasonIdUpcomingRequestObject) (api.GetSeasonsSeasonIdUpcomingResponseObject, error) {
	matches, err := s.GetSeasonUpcomingMatches(ctx, int32(request.SeasonId))
	if err != nil {
		return api.GetSeasonsSeasonIdUpcoming200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get upcoming matches: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	matchesData := map[string]interface{}{
		"matches": matches,
	}
	return api.GetSeasonsSeasonIdUpcoming200JSONResponse(api.ApiResult{
		Data:      &matchesData,
		IsSuccess: Ptr(true),
	}), nil
}
