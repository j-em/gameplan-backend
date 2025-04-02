package api_server

import (
	"context"

	"github.com/gameplan-backend/api"
	"github.com/gameplan-backend/db"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stripe/stripe-go/v78/client"
	"github.com/stytchauth/stytch-go/v16/stytch/consumer/stytchapi"
)

// MyApiServer provides a concrete implementation of the generated StrictServerInterface.
type MyApiServer struct {
	StytchClient        *stytchapi.API
	StripeClient        *client.API
	DB                  *db.Queries
	Emailer             *mailgun.MailgunImpl
	SubscriptionsServer *SubscriptionsServer
	AuthServer          *AuthServer
	MatchesServer       *MatchesServer
	PlayersServer       *PlayersServer
	SeasonsServer       *SeasonsServer
}

func (s MyApiServer) PostMatches(ctx context.Context, request api.PostMatchesRequestObject) (api.PostMatchesResponseObject, error) {
	return s.MatchesServer.PostMatches(ctx, request)
}

func (s MyApiServer) PutMatchesBatches(ctx context.Context, request api.PutMatchesBatchesRequestObject) (api.PutMatchesBatchesResponseObject, error) {
	return s.MatchesServer.PutMatchesBatches(ctx, request)
}

func (s MyApiServer) PostMatchesUnassignPlayerFromMatch(ctx context.Context, request api.PostMatchesUnassignPlayerFromMatchRequestObject) (api.PostMatchesUnassignPlayerFromMatchResponseObject, error) {
	return s.MatchesServer.PostMatchesUnassignPlayerFromMatch(ctx, request)
}

func (s MyApiServer) DeleteMatchesMatchId(ctx context.Context, request api.DeleteMatchesMatchIdRequestObject) (api.DeleteMatchesMatchIdResponseObject, error) {
	return s.MatchesServer.DeleteMatchesMatchId(ctx, request)
}

func (s MyApiServer) PutMatchesMatchId(ctx context.Context, request api.PutMatchesMatchIdRequestObject) (api.PutMatchesMatchIdResponseObject, error) {
	return s.MatchesServer.PutMatchesMatchId(ctx, request)
}

func (s MyApiServer) GetPlayers(ctx context.Context, request api.GetPlayersRequestObject) (api.GetPlayersResponseObject, error) {
	return s.PlayersServer.GetPlayers(ctx, request)
}

func (s MyApiServer) PostPlayers(ctx context.Context, request api.PostPlayersRequestObject) (api.PostPlayersResponseObject, error) {
	return s.PlayersServer.PostPlayers(ctx, request)
}

func (s MyApiServer) DeletePlayersPlayerId(ctx context.Context, request api.DeletePlayersPlayerIdRequestObject) (api.DeletePlayersPlayerIdResponseObject, error) {
	return s.PlayersServer.DeletePlayersPlayerId(ctx, request)
}

func (s MyApiServer) GetPlayersPlayerId(ctx context.Context, request api.GetPlayersPlayerIdRequestObject) (api.GetPlayersPlayerIdResponseObject, error) {
	return s.PlayersServer.GetPlayersPlayerId(ctx, request)
}

func (s MyApiServer) PutPlayersPlayerId(ctx context.Context, request api.PutPlayersPlayerIdRequestObject) (api.PutPlayersPlayerIdResponseObject, error) {
	return s.PlayersServer.PutPlayersPlayerId(ctx, request)
}

func (s MyApiServer) GetPlayersPlayerIdCustomColumns(ctx context.Context, request api.GetPlayersPlayerIdCustomColumnsRequestObject) (api.GetPlayersPlayerIdCustomColumnsResponseObject, error) {
	return s.PlayersServer.GetPlayersPlayerIdCustomColumns(ctx, request)
}

func (s MyApiServer) PutPlayersPlayerIdCustomColumns(ctx context.Context, request api.PutPlayersPlayerIdCustomColumnsRequestObject) (api.PutPlayersPlayerIdCustomColumnsResponseObject, error) {
	return s.PlayersServer.PutPlayersPlayerIdCustomColumns(ctx, request)
}

func (s MyApiServer) GetPlayersPlayerIdSchedule(ctx context.Context, request api.GetPlayersPlayerIdScheduleRequestObject) (api.GetPlayersPlayerIdScheduleResponseObject, error) {
	return s.PlayersServer.GetPlayersPlayerIdSchedule(ctx, request)
}

func (s MyApiServer) GetSeasons(ctx context.Context, request api.GetSeasonsRequestObject) (api.GetSeasonsResponseObject, error) {
	return s.SeasonsServer.GetSeasons(ctx, request)
}

func (s MyApiServer) PostSeasons(ctx context.Context, request api.PostSeasonsRequestObject) (api.PostSeasonsResponseObject, error) {
	return s.SeasonsServer.PostSeasons(ctx, request)
}

func (s MyApiServer) GetSeasonsTotalAmount(ctx context.Context, request api.GetSeasonsTotalAmountRequestObject) (api.GetSeasonsTotalAmountResponseObject, error) {
	return s.SeasonsServer.GetSeasonsTotalAmount(ctx, request)
}

func (s MyApiServer) DeleteSeasonsSeasonId(ctx context.Context, request api.DeleteSeasonsSeasonIdRequestObject) (api.DeleteSeasonsSeasonIdResponseObject, error) {
	return s.SeasonsServer.DeleteSeasonsSeasonId(ctx, request)
}

func (s MyApiServer) GetSeasonsSeasonId(ctx context.Context, request api.GetSeasonsSeasonIdRequestObject) (api.GetSeasonsSeasonIdResponseObject, error) {
	return s.SeasonsServer.GetSeasonsSeasonId(ctx, request)
}

func (s MyApiServer) PutSeasonsSeasonId(ctx context.Context, request api.PutSeasonsSeasonIdRequestObject) (api.PutSeasonsSeasonIdResponseObject, error) {
	return s.SeasonsServer.PutSeasonsSeasonId(ctx, request)
}

func (s MyApiServer) GetSeasonsSeasonIdPublicScheduleLink(ctx context.Context, request api.GetSeasonsSeasonIdPublicScheduleLinkRequestObject) (api.GetSeasonsSeasonIdPublicScheduleLinkResponseObject, error) {
	return s.SeasonsServer.GetSeasonsSeasonIdPublicScheduleLink(ctx, request)
}

func (s MyApiServer) GetSeasonsSeasonIdScoreboard(ctx context.Context, request api.GetSeasonsSeasonIdScoreboardRequestObject) (api.GetSeasonsSeasonIdScoreboardResponseObject, error) {
	return s.SeasonsServer.GetSeasonsSeasonIdScoreboard(ctx, request)
}

func (s MyApiServer) GetSeasonsSeasonIdUpcoming(ctx context.Context, request api.GetSeasonsSeasonIdUpcomingRequestObject) (api.GetSeasonsSeasonIdUpcomingResponseObject, error) {
	return s.SeasonsServer.GetSeasonsSeasonIdUpcoming(ctx, request)
}

func (s MyApiServer) PostSessions(ctx context.Context, request api.PostSessionsRequestObject) (api.PostSessionsResponseObject, error) {
	return s.AuthServer.PostSessions(ctx, request)
}

func (s MyApiServer) PostSubscriptionsHandleSuccessUpgrade(ctx context.Context, request api.PostSubscriptionsHandleSuccessUpgradeRequestObject) (api.PostSubscriptionsHandleSuccessUpgradeResponseObject, error) {
	return s.SubscriptionsServer.PostSubscriptionsHandleSuccessUpgrade(ctx, request)
}

func (s MyApiServer) PostSubscriptionsInitUpdatePaymentMethod(ctx context.Context, request api.PostSubscriptionsInitUpdatePaymentMethodRequestObject) (api.PostSubscriptionsInitUpdatePaymentMethodResponseObject, error) {
	return s.SubscriptionsServer.PostSubscriptionsInitUpdatePaymentMethod(ctx, request)
}

func (s MyApiServer) PostSubscriptionsUpgradeUserSubscription(ctx context.Context, request api.PostSubscriptionsUpgradeUserSubscriptionRequestObject) (api.PostSubscriptionsUpgradeUserSubscriptionResponseObject, error) {
	return s.SubscriptionsServer.PostSubscriptionsUpgradeUserSubscription(ctx, request)
}

func (s MyApiServer) PostSupportMessages(ctx context.Context, request api.PostSupportMessagesRequestObject) (api.PostSupportMessagesResponseObject, error) {
	return s.AuthServer.PostSupportMessages(ctx, request)
}

func (s MyApiServer) DeleteUsersUserId(ctx context.Context, request api.DeleteUsersUserIdRequestObject) (api.DeleteUsersUserIdResponseObject, error) {
	return s.AuthServer.DeleteUsersUserId(ctx, request)
}

func (s MyApiServer) GetUsersUserIdAppsettings(ctx context.Context, request api.GetUsersUserIdAppsettingsRequestObject) (api.GetUsersUserIdAppsettingsResponseObject, error) {
	return s.AuthServer.GetUsersUserIdAppsettings(ctx, request)
}

func (s MyApiServer) PostUsersUserIdAppsettings(ctx context.Context, request api.PostUsersUserIdAppsettingsRequestObject) (api.PostUsersUserIdAppsettingsResponseObject, error) {
	return s.AuthServer.PostUsersUserIdAppsettings(ctx, request)
}

func (s MyApiServer) GetUsersUserIdCustomPlayerColumns(ctx context.Context, request api.GetUsersUserIdCustomPlayerColumnsRequestObject) (api.GetUsersUserIdCustomPlayerColumnsResponseObject, error) {
	return s.AuthServer.GetUsersUserIdCustomPlayerColumns(ctx, request)
}

func (s MyApiServer) PostUsersUserIdCustomPlayerColumns(ctx context.Context, request api.PostUsersUserIdCustomPlayerColumnsRequestObject) (api.PostUsersUserIdCustomPlayerColumnsResponseObject, error) {
	return s.AuthServer.PostUsersUserIdCustomPlayerColumns(ctx, request)
}

func (s MyApiServer) DeleteUsersUserIdCustomPlayerColumnsColumnId(ctx context.Context, request api.DeleteUsersUserIdCustomPlayerColumnsColumnIdRequestObject) (api.DeleteUsersUserIdCustomPlayerColumnsColumnIdResponseObject, error) {
	return s.AuthServer.DeleteUsersUserIdCustomPlayerColumnsColumnId(ctx, request)
}

func (s MyApiServer) PostUsersUserIdResetCurrentUserPassword(ctx context.Context, request api.PostUsersUserIdResetCurrentUserPasswordRequestObject) (api.PostUsersUserIdResetCurrentUserPasswordResponseObject, error) {
	return s.AuthServer.PostUsersUserIdResetCurrentUserPassword(ctx, request)
}

func (s MyApiServer) DeleteUsersUserIdSubscription(ctx context.Context, request api.DeleteUsersUserIdSubscriptionRequestObject) (api.DeleteUsersUserIdSubscriptionResponseObject, error) {
	return s.SubscriptionsServer.DeleteUsersUserIdSubscription(ctx, request)
}

func (s MyApiServer) GetUsersUserIdSubscription(ctx context.Context, request api.GetUsersUserIdSubscriptionRequestObject) (api.GetUsersUserIdSubscriptionResponseObject, error) {
	return s.SubscriptionsServer.GetUsersUserIdSubscription(ctx, request)
}

func (s MyApiServer) GetUsersUserIdUsersettings(ctx context.Context, request api.GetUsersUserIdUsersettingsRequestObject) (api.GetUsersUserIdUsersettingsResponseObject, error) {
	return s.AuthServer.GetUsersUserIdUsersettings(ctx, request)
}

func (s MyApiServer) PostUsersUserIdUsersettings(ctx context.Context, request api.PostUsersUserIdUsersettingsRequestObject) (api.PostUsersUserIdUsersettingsResponseObject, error) {
	return s.AuthServer.PostUsersUserIdUsersettings(ctx, request)
}

func (s MyApiServer) PostUsersSendResetPasswordLink(ctx context.Context, request api.PostUsersSendResetPasswordLinkRequestObject) (api.PostUsersSendResetPasswordLinkResponseObject, error) {
	return s.AuthServer.PostUsersSendResetPasswordLink(ctx, request)
}

func (s MyApiServer) PostUsersSendVerificationEmail(ctx context.Context, request api.PostUsersSendVerificationEmailRequestObject) (api.PostUsersSendVerificationEmailResponseObject, error) {
	return s.AuthServer.PostUsersSendVerificationEmail(ctx, request)
}

func (s MyApiServer) PostUsersSignUpUser(ctx context.Context, request api.PostUsersSignUpUserRequestObject) (api.PostUsersSignUpUserResponseObject, error) {
	return s.AuthServer.PostUsersSignUpUser(ctx, request)
}

func (s MyApiServer) PostUsersUpdateUserPassword(ctx context.Context, request api.PostUsersUpdateUserPasswordRequestObject) (api.PostUsersUpdateUserPasswordResponseObject, error) {
	return s.AuthServer.PostUsersUpdateUserPassword(ctx, request)
}

func (s MyApiServer) PostUsersVerifyMagicLinkToken(ctx context.Context, request api.PostUsersVerifyMagicLinkTokenRequestObject) (api.PostUsersVerifyMagicLinkTokenResponseObject, error) {
	return s.AuthServer.PostUsersVerifyMagicLinkToken(ctx, request)
}
