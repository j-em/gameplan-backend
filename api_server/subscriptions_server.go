package api_server

import (
	"context"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/stripe/stripe-go/v81/client"
)

type SubscriptionsServer struct {
	StripeClient *client.API
	DB           *db.Queries
}

func (s *SubscriptionsServer) PostSubscriptionsHandleSuccessUpgrade(ctx context.Context, request api.PostSubscriptionsHandleSuccessUpgradeRequestObject) (api.PostSubscriptionsHandleSuccessUpgradeResponseObject, error) {
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

func (s *SubscriptionsServer) PostSubscriptionsInitUpdatePaymentMethod(ctx context.Context, request api.PostSubscriptionsInitUpdatePaymentMethodRequestObject) (api.PostSubscriptionsInitUpdatePaymentMethodResponseObject, error) {
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

func (s *SubscriptionsServer) PostSubscriptionsUpgradeUserSubscription(ctx context.Context, request api.PostSubscriptionsUpgradeUserSubscriptionRequestObject) (api.PostSubscriptionsUpgradeUserSubscriptionResponseObject, error) {
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

func (s SubscriptionsServer) DeleteUsersUserIdSubscription(ctx context.Context, request api.DeleteUsersUserIdSubscriptionRequestObject) (api.DeleteUsersUserIdSubscriptionResponseObject, error) {
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

func (s *SubscriptionsServer) GetUsersUserIdSubscription(ctx context.Context, request api.GetUsersUserIdSubscriptionRequestObject) (api.GetUsersUserIdSubscriptionResponseObject, error) {
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
