package api_server

import (
	"context"
	"fmt"

	"github.com/gameplan-backend/api"
)

// Define handler functions for user and app settings
func (s *MyApiServer) GetUsersUserIdAppsettingsInternal(ctx context.Context, request api.GetUsersUserIdAppsettingsRequestObject) (api.GetUsersUserIdAppsettingsResponseObject, error) {
	appSettings, err := s.DB.GetUserAppSettings(ctx, int32(request.UserId))
	if err != nil {
		return api.GetUsersUserIdAppsettings200JSONResponse(api.ApiResult{
			Error: &struct {
				Code    *string `json:"code,omitempty"`
				Message *string `json:"message,omitempty"`
			}{
				Code:    Ptr("DATABASE_ERROR"),
				Message: Ptr(fmt.Sprintf("Failed to get app settings: %v", err)),
			},
			IsSuccess: Ptr(false),
		}), nil
	}

	responseData := map[string]interface{}{
		"settings": appSettings,
	}

	return api.GetUsersUserIdAppsettings200JSONResponse(api.ApiResult{
		Data:      &responseData,
		IsSuccess: Ptr(true),
	}), nil
}
