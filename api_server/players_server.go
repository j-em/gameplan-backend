package api_server

import (
	"context"
	"fmt"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// PlayersServer handles player-related operations
type PlayersServer struct {
	DB *db.Queries
}

// CreatePlayer creates a new player record based on API params
func (s *PlayersServer) CreatePlayer(
	ctx context.Context,
	userId int32,
	name string,
	email pgtype.Text,
	emailNotificationsEnabled bool,
) (*db.Player, error) {
	params := db.CreatePlayerParams{
		Userid:                    pgtype.Int4{Int32: userId, Valid: true},
		Name:                      name,
		Email:                     email,
		Preferredmatchgroup:       pgtype.Int4{Valid: false},
		Emailnotificationsenabled: emailNotificationsEnabled,
	}

	player, err := s.DB.CreatePlayer(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create player: %w", err)
	}
	return &player, nil
}

// GetPlayer retrieves player details by ID with user auth check
func (s *PlayersServer) GetPlayer(
	ctx context.Context,
	userId int32,
	playerId int32,
) (*db.Player, error) {
	player, err := s.DB.GetPlayer(ctx, db.GetPlayerParams{
		ID:     playerId,
		Userid: pgtype.Int4{Int32: userId, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get player: %w", err)
	}
	return &player, nil
}

// UpdatePlayer updates player information via key-value pairs
func (s *PlayersServer) UpdatePlayer(
	ctx context.Context,
	userId int32,
	playerId int32,
	updates map[string]interface{},
) (*db.Player, error) {
	// Get current player to preserve unchanged fields
	current, err := s.GetPlayer(ctx, userId, playerId)
	if err != nil {
		return nil, err
	}

	params := db.UpdatePlayerParams{
		ID:                        playerId,
		Userid:                    pgtype.Int4{Int32: userId, Valid: true},
		Name:                      current.Name,
		Email:                     current.Email,
		Preferredmatchgroup:       current.Preferredmatchgroup,
		Emailnotificationsenabled: current.Emailnotificationsenabled,
		Isactive:                  true,
	}

	// Apply updates from the key-value map
	for key, value := range updates {
		switch key {
		case "name":
			params.Name = value.(string)
		case "email":
			params.Email = pgtype.Text{String: value.(string), Valid: value.(string) != ""}
		case "emailNotificationsEnabled":
			params.Emailnotificationsenabled = value.(bool)
		}
	}

	player, err := s.DB.UpdatePlayer(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update player: %w", err)
	}
	return &player, nil
}

// DeletePlayer removes a player record with user auth check
func (s *PlayersServer) DeletePlayer(
	ctx context.Context,
	userId int32,
	playerId int32,
) error {
	if err := s.DB.DeletePlayer(ctx, db.DeletePlayerParams{
		ID:     playerId,
		Userid: pgtype.Int4{Int32: userId, Valid: true},
	}); err != nil {
		return fmt.Errorf("failed to delete player: %w", err)
	}
	return nil
}

// ListPlayers retrieves all players for a user
func (s *PlayersServer) ListPlayers(
	ctx context.Context,
	userId int32,
) ([]db.Player, error) {
	players, err := s.DB.GetPlayers(ctx, pgtype.Int4{Int32: userId, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list players: %w", err)
	}
	return players, nil
}

// GetPlayerCustomColumns retrieves active custom columns
func (s *PlayersServer) GetPlayerCustomColumns(
	ctx context.Context,
) ([]db.PlayerCustomColumn, error) {
	columns, err := s.DB.GetPlayerCustomColumns(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get player custom columns: %w", err)
	}
	return columns, nil
}

// GetPlayerCustomValues retrieves custom values for a player
func (s *PlayersServer) GetPlayerCustomValues(
	ctx context.Context,
	playerId int32,
) ([]db.GetPlayerCustomValuesRow, error) {
	values, err := s.DB.GetPlayerCustomValues(ctx, pgtype.Int4{Int32: playerId, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get player custom values: %w", err)
	}
	return values, nil
}

// UpsertPlayerCustomValue updates or inserts a custom value
func (s *PlayersServer) UpsertPlayerCustomValue(
	ctx context.Context,
	playerId int32,
	columnId int32,
	value string,
) (*db.PlayerCustomValue, error) {
	val, err := s.DB.UpsertPlayerCustomValue(ctx, db.UpsertPlayerCustomValueParams{
		PlayerID: pgtype.Int4{Int32: playerId, Valid: true},
		ColumnID: pgtype.Int4{Int32: columnId, Valid: true},
		Value:    pgtype.Text{String: value, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upsert player custom value: %w", err)
	}
	return &val, nil
}

// API endpoint implementations

func (s *PlayersServer) GetPlayers(ctx context.Context, request api.GetPlayersRequestObject) (api.GetPlayersResponseObject, error) {
	// Extract user ID from context (would come from JWT in middleware)
	userID := int32(1) // TODO: Get from auth context

	players, err := s.ListPlayers(ctx, userID)
	if err != nil {
		return api.GetPlayers200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to list players: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	playersMap := map[string]interface{}{
		"players": players,
	}
	return api.GetPlayers200JSONResponse(api.ApiResult{
		Data:      &playersMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) PostPlayers(ctx context.Context, request api.PostPlayersRequestObject) (api.PostPlayersResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	player, err := s.CreatePlayer(
		ctx,
		userID,
		request.Body.Name,
		pgtype.Text{String: *request.Body.Email, Valid: request.Body.Email != nil},
		request.Body.EmailNotificationsEnabled,
	)
	if err != nil {
		return api.PostPlayers200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to create player: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	playerMap := map[string]interface{}{
		"player": *player,
	}
	return api.PostPlayers200JSONResponse(api.ApiResult{
		Data:      &playerMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) DeletePlayersPlayerId(ctx context.Context, request api.DeletePlayersPlayerIdRequestObject) (api.DeletePlayersPlayerIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	if err := s.DeletePlayer(ctx, userID, int32(request.PlayerId)); err != nil {
		return api.DeletePlayersPlayerId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to delete player: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	return api.DeletePlayersPlayerId200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) GetPlayersPlayerId(ctx context.Context, request api.GetPlayersPlayerIdRequestObject) (api.GetPlayersPlayerIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	player, err := s.GetPlayer(ctx, userID, int32(request.PlayerId))
	if err != nil {
		return api.GetPlayersPlayerId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get player: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	customValues, err := s.GetPlayerCustomValues(ctx, int32(request.PlayerId))
	if err != nil {
		// Log but continue since custom values are optional
		fmt.Printf("Failed to get custom values: %v\n", err)
	}

	customMap := make(map[string]string)
	for _, cv := range customValues {
		if cv.Value.Valid {
			customMap[cv.ColumnName] = cv.Value.String
		}
	}

	playerDetails := map[string]interface{}{
		"player":       *player,
		"customValues": customMap,
	}
	return api.GetPlayersPlayerId200JSONResponse(api.ApiResult{
		Data:      &playerDetails,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) PutPlayersPlayerId(ctx context.Context, request api.PutPlayersPlayerIdRequestObject) (api.PutPlayersPlayerIdResponseObject, error) {
	userID := int32(1) // TODO: Get from auth context

	updates := map[string]interface{}{
		request.Body.Key: request.Body.Value,
	}

	player, err := s.UpdatePlayer(ctx, userID, int32(request.PlayerId), updates)
	if err != nil {
		return api.PutPlayersPlayerId200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to update player: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	playerMap := map[string]interface{}{
		"player": *player,
	}
	return api.PutPlayersPlayerId200JSONResponse(api.ApiResult{
		Data:      &playerMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) GetPlayersPlayerIdSchedule(ctx context.Context, request api.GetPlayersPlayerIdScheduleRequestObject) (api.GetPlayersPlayerIdScheduleResponseObject, error) {
	// TODO: Implement schedule retrieval
	return api.GetPlayersPlayerIdSchedule200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) GetPlayersPlayerIdCustomColumns(ctx context.Context, request api.GetPlayersPlayerIdCustomColumnsRequestObject) (api.GetPlayersPlayerIdCustomColumnsResponseObject, error) {
	columns, err := s.GetPlayerCustomColumns(ctx)
	if err != nil {
		return api.GetPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get custom columns: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	columnsMap := map[string]interface{}{
		"columns": columns,
	}
	return api.GetPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
		Data:      &columnsMap,
		IsSuccess: Ptr(true),
	}), nil
}

func (s *PlayersServer) PutPlayersPlayerIdCustomColumns(ctx context.Context, request api.PutPlayersPlayerIdCustomColumnsRequestObject) (api.PutPlayersPlayerIdCustomColumnsResponseObject, error) {
	// Update custom column value
	_, err := s.UpsertPlayerCustomValue(ctx,
		int32(request.Body.PlayerId),
		int32(request.Body.ColumnId),
		request.Body.Value)
	if err != nil {
		return api.PutPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DB_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to update custom column: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	return api.PutPlayersPlayerIdCustomColumns200JSONResponse(api.ApiResult{
		IsSuccess: Ptr(true),
	}), nil
}
