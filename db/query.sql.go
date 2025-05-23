// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMatch = `-- name: CreateMatch :one
INSERT INTO matches (
    seasonId, playerId1, playerId1Points, playerId2, playerId2Points, matchDate, winnerId, "group"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, seasonid, playerid1, playerid1points, playerid2, playerid2points, matchdate, winnerid, createdat, updatedat, isactive, "group"
`

type CreateMatchParams struct {
	Seasonid        pgtype.Int4
	Playerid1       pgtype.Int4
	Playerid1points int32
	Playerid2       pgtype.Int4
	Playerid2points int32
	Matchdate       pgtype.Date
	Winnerid        pgtype.Int4
	Group           int32
}

func (q *Queries) CreateMatch(ctx context.Context, arg CreateMatchParams) (Match, error) {
	row := q.db.QueryRow(ctx, createMatch,
		arg.Seasonid,
		arg.Playerid1,
		arg.Playerid1points,
		arg.Playerid2,
		arg.Playerid2points,
		arg.Matchdate,
		arg.Winnerid,
		arg.Group,
	)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.Seasonid,
		&i.Playerid1,
		&i.Playerid1points,
		&i.Playerid2,
		&i.Playerid2points,
		&i.Matchdate,
		&i.Winnerid,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Group,
	)
	return i, err
}

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
    userId, name, email, preferredMatchGroup, emailNotificationsEnabled
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, userid, name, email, createdat, updatedat, preferredmatchgroup, isactive, emailnotificationsenabled
`

type CreatePlayerParams struct {
	Userid                    pgtype.Int4
	Name                      string
	Email                     pgtype.Text
	Preferredmatchgroup       pgtype.Int4
	Emailnotificationsenabled bool
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, createPlayer,
		arg.Userid,
		arg.Name,
		arg.Email,
		arg.Preferredmatchgroup,
		arg.Emailnotificationsenabled,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Email,
		&i.Createdat,
		&i.Updatedat,
		&i.Preferredmatchgroup,
		&i.Isactive,
		&i.Emailnotificationsenabled,
	)
	return i, err
}

const createPlayerCustomColumn = `-- name: CreatePlayerCustomColumn :one
INSERT INTO player_custom_columns (
    name, field_type, description, is_required, display_order
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, name, field_type, description, is_required, is_active, display_order, createdat, updatedat
`

type CreatePlayerCustomColumnParams struct {
	Name         string
	FieldType    string
	Description  pgtype.Text
	IsRequired   pgtype.Bool
	DisplayOrder pgtype.Int4
}

func (q *Queries) CreatePlayerCustomColumn(ctx context.Context, arg CreatePlayerCustomColumnParams) (PlayerCustomColumn, error) {
	row := q.db.QueryRow(ctx, createPlayerCustomColumn,
		arg.Name,
		arg.FieldType,
		arg.Description,
		arg.IsRequired,
		arg.DisplayOrder,
	)
	var i PlayerCustomColumn
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FieldType,
		&i.Description,
		&i.IsRequired,
		&i.IsActive,
		&i.DisplayOrder,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const createSeason = `-- name: CreateSeason :one
INSERT INTO seasons (
    userId, name, startDate, seasonType, frequency
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, userid, name, startdate, createdat, updatedat, isactive, seasontype, frequency
`

type CreateSeasonParams struct {
	Userid     pgtype.Int4
	Name       string
	Startdate  pgtype.Date
	Seasontype string
	Frequency  string
}

func (q *Queries) CreateSeason(ctx context.Context, arg CreateSeasonParams) (Season, error) {
	row := q.db.QueryRow(ctx, createSeason,
		arg.Userid,
		arg.Name,
		arg.Startdate,
		arg.Seasontype,
		arg.Frequency,
	)
	var i Season
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Startdate,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Seasontype,
		&i.Frequency,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    stytchId, stripeId, name, email, phone, country, birthday, lang, isVerified
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, stytchid, stripeid, name, email, phone, country, birthday, lang, createdat, updatedat, isactive, isverified, subscriptiontier, jsonsettings
`

type CreateUserParams struct {
	Stytchid   string
	Stripeid   string
	Name       string
	Email      string
	Phone      pgtype.Text
	Country    pgtype.Text
	Birthday   pgtype.Int4
	Lang       string
	Isverified bool
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Stytchid,
		arg.Stripeid,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Country,
		arg.Birthday,
		arg.Lang,
		arg.Isverified,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Stytchid,
		&i.Stripeid,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Country,
		&i.Birthday,
		&i.Lang,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Isverified,
		&i.Subscriptiontier,
		&i.Jsonsettings,
	)
	return i, err
}

const deleteMatch = `-- name: DeleteMatch :exec
DELETE FROM matches
WHERE id = $1
`

func (q *Queries) DeleteMatch(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteMatch, id)
	return err
}

const deletePlayer = `-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1 AND userId = $2
`

type DeletePlayerParams struct {
	ID     int32
	Userid pgtype.Int4
}

func (q *Queries) DeletePlayer(ctx context.Context, arg DeletePlayerParams) error {
	_, err := q.db.Exec(ctx, deletePlayer, arg.ID, arg.Userid)
	return err
}

const deletePlayerCustomColumn = `-- name: DeletePlayerCustomColumn :exec
DELETE FROM player_custom_columns
WHERE id = $1
`

func (q *Queries) DeletePlayerCustomColumn(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePlayerCustomColumn, id)
	return err
}

const deleteSeason = `-- name: DeleteSeason :exec
DELETE FROM seasons
WHERE id = $1 AND userId = $2
`

type DeleteSeasonParams struct {
	ID     int32
	Userid pgtype.Int4
}

func (q *Queries) DeleteSeason(ctx context.Context, arg DeleteSeasonParams) error {
	_, err := q.db.Exec(ctx, deleteSeason, arg.ID, arg.Userid)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users SET isActive = false, updatedAt = CURRENT_TIMESTAMP WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const deleteUserSubscription = `-- name: DeleteUserSubscription :exec
UPDATE users
SET subscriptionTier = 'free',
    stripeId = '',
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DeleteUserSubscription(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUserSubscription, id)
	return err
}

const getMatch = `-- name: GetMatch :one
SELECT id, seasonid, playerid1, playerid1points, playerid2, playerid2points, matchdate, winnerid, createdat, updatedat, isactive, "group" FROM matches
WHERE id = $1
`

func (q *Queries) GetMatch(ctx context.Context, id int32) (Match, error) {
	row := q.db.QueryRow(ctx, getMatch, id)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.Seasonid,
		&i.Playerid1,
		&i.Playerid1points,
		&i.Playerid2,
		&i.Playerid2points,
		&i.Matchdate,
		&i.Winnerid,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Group,
	)
	return i, err
}

const getPlayer = `-- name: GetPlayer :one
SELECT id, userid, name, email, createdat, updatedat, preferredmatchgroup, isactive, emailnotificationsenabled FROM players
WHERE id = $1 AND userId = $2
`

type GetPlayerParams struct {
	ID     int32
	Userid pgtype.Int4
}

func (q *Queries) GetPlayer(ctx context.Context, arg GetPlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayer, arg.ID, arg.Userid)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Email,
		&i.Createdat,
		&i.Updatedat,
		&i.Preferredmatchgroup,
		&i.Isactive,
		&i.Emailnotificationsenabled,
	)
	return i, err
}

const getPlayerCustomColumns = `-- name: GetPlayerCustomColumns :many
SELECT id, name, field_type, description, is_required, is_active, display_order, createdat, updatedat FROM player_custom_columns
WHERE is_active = true
ORDER BY display_order
`

func (q *Queries) GetPlayerCustomColumns(ctx context.Context) ([]PlayerCustomColumn, error) {
	rows, err := q.db.Query(ctx, getPlayerCustomColumns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlayerCustomColumn
	for rows.Next() {
		var i PlayerCustomColumn
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.FieldType,
			&i.Description,
			&i.IsRequired,
			&i.IsActive,
			&i.DisplayOrder,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayerCustomValues = `-- name: GetPlayerCustomValues :many
SELECT pcv.id, pcv.player_id, pcv.column_id, pcv.value, pcv.createdat, pcv.updatedat, pcc.name as column_name, pcc.field_type
FROM player_custom_values pcv
JOIN player_custom_columns pcc ON pcv.column_id = pcc.id
WHERE pcv.player_id = $1
`

type GetPlayerCustomValuesRow struct {
	ID         int32
	PlayerID   pgtype.Int4
	ColumnID   pgtype.Int4
	Value      pgtype.Text
	Createdat  pgtype.Timestamp
	Updatedat  pgtype.Timestamp
	ColumnName string
	FieldType  string
}

func (q *Queries) GetPlayerCustomValues(ctx context.Context, playerID pgtype.Int4) ([]GetPlayerCustomValuesRow, error) {
	rows, err := q.db.Query(ctx, getPlayerCustomValues, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPlayerCustomValuesRow
	for rows.Next() {
		var i GetPlayerCustomValuesRow
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.ColumnID,
			&i.Value,
			&i.Createdat,
			&i.Updatedat,
			&i.ColumnName,
			&i.FieldType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayers = `-- name: GetPlayers :many
SELECT id, userid, name, email, createdat, updatedat, preferredmatchgroup, isactive, emailnotificationsenabled FROM players
WHERE userId = $1 AND isActive = true
`

func (q *Queries) GetPlayers(ctx context.Context, userid pgtype.Int4) ([]Player, error) {
	rows, err := q.db.Query(ctx, getPlayers, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Name,
			&i.Email,
			&i.Createdat,
			&i.Updatedat,
			&i.Preferredmatchgroup,
			&i.Isactive,
			&i.Emailnotificationsenabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSeason = `-- name: GetSeason :one
SELECT id, userid, name, startdate, createdat, updatedat, isactive, seasontype, frequency FROM seasons
WHERE id = $1 AND userId = $2
`

type GetSeasonParams struct {
	ID     int32
	Userid pgtype.Int4
}

func (q *Queries) GetSeason(ctx context.Context, arg GetSeasonParams) (Season, error) {
	row := q.db.QueryRow(ctx, getSeason, arg.ID, arg.Userid)
	var i Season
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Startdate,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Seasontype,
		&i.Frequency,
	)
	return i, err
}

const getSeasonScoreboard = `-- name: GetSeasonScoreboard :many
SELECT p.id as player_id, p.name as player_name, COUNT(m.winnerId) as wins
FROM players p
LEFT JOIN matches m ON p.id = m.winnerId AND m.seasonId = $1
WHERE p.userId = (SELECT s.userId FROM seasons s WHERE s.id = $2)
GROUP BY p.id, p.name
ORDER BY wins DESC
`

type GetSeasonScoreboardParams struct {
	Seasonid pgtype.Int4
	ID       int32
}

type GetSeasonScoreboardRow struct {
	PlayerID   int32
	PlayerName string
	Wins       int64
}

func (q *Queries) GetSeasonScoreboard(ctx context.Context, arg GetSeasonScoreboardParams) ([]GetSeasonScoreboardRow, error) {
	rows, err := q.db.Query(ctx, getSeasonScoreboard, arg.Seasonid, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSeasonScoreboardRow
	for rows.Next() {
		var i GetSeasonScoreboardRow
		if err := rows.Scan(&i.PlayerID, &i.PlayerName, &i.Wins); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSeasonUpcomingMatches = `-- name: GetSeasonUpcomingMatches :many
SELECT id, seasonid, playerid1, playerid1points, playerid2, playerid2points, matchdate, winnerid, createdat, updatedat, isactive, "group" FROM matches
WHERE seasonId = $1 AND matchDate > CURRENT_TIMESTAMP
ORDER BY matchDate ASC
LIMIT 5
`

func (q *Queries) GetSeasonUpcomingMatches(ctx context.Context, seasonid pgtype.Int4) ([]Match, error) {
	rows, err := q.db.Query(ctx, getSeasonUpcomingMatches, seasonid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Match
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.Seasonid,
			&i.Playerid1,
			&i.Playerid1points,
			&i.Playerid2,
			&i.Playerid2points,
			&i.Matchdate,
			&i.Winnerid,
			&i.Createdat,
			&i.Updatedat,
			&i.Isactive,
			&i.Group,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSeasons = `-- name: GetSeasons :many
SELECT id, userid, name, startdate, createdat, updatedat, isactive, seasontype, frequency FROM seasons
WHERE userId = $1 AND isActive = true
`

func (q *Queries) GetSeasons(ctx context.Context, userid pgtype.Int4) ([]Season, error) {
	rows, err := q.db.Query(ctx, getSeasons, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Season
	for rows.Next() {
		var i Season
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Name,
			&i.Startdate,
			&i.Createdat,
			&i.Updatedat,
			&i.Isactive,
			&i.Seasontype,
			&i.Frequency,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserAppSettings = `-- name: GetUserAppSettings :one
SELECT jsonSettings FROM users
WHERE id = $1
`

func (q *Queries) GetUserAppSettings(ctx context.Context, id int32) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getUserAppSettings, id)
	var jsonsettings pgtype.Text
	err := row.Scan(&jsonsettings)
	return jsonsettings, err
}

const getUserByEmail = `-- name: GetUserByEmail :one

SELECT id, email, isVerified
FROM users
WHERE email = $1
`

type GetUserByEmailRow struct {
	ID         int32
	Email      string
	Isverified bool
}

// query.sql with PostgreSQL parameter syntax
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(&i.ID, &i.Email, &i.Isverified)
	return i, err
}

const getUserSubscription = `-- name: GetUserSubscription :one
SELECT subscriptionTier, stripeId
FROM users
WHERE id = $1
`

type GetUserSubscriptionRow struct {
	Subscriptiontier string
	Stripeid         string
}

func (q *Queries) GetUserSubscription(ctx context.Context, id int32) (GetUserSubscriptionRow, error) {
	row := q.db.QueryRow(ctx, getUserSubscription, id)
	var i GetUserSubscriptionRow
	err := row.Scan(&i.Subscriptiontier, &i.Stripeid)
	return i, err
}

const getUserUserSettings = `-- name: GetUserUserSettings :one
SELECT jsonSettings FROM users
WHERE id = $1
`

func (q *Queries) GetUserUserSettings(ctx context.Context, id int32) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, getUserUserSettings, id)
	var jsonsettings pgtype.Text
	err := row.Scan(&jsonsettings)
	return jsonsettings, err
}

const updateMatch = `-- name: UpdateMatch :one
UPDATE matches
SET seasonId = $1,
    playerId1 = $2,
    playerId1Points = $3,
    playerId2 = $4,
    playerId2Points = $5,
    matchDate = $6,
    winnerId = $7,
    "group" = $8,
    isActive = $9,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $10
RETURNING id, seasonid, playerid1, playerid1points, playerid2, playerid2points, matchdate, winnerid, createdat, updatedat, isactive, "group"
`

type UpdateMatchParams struct {
	Seasonid        pgtype.Int4
	Playerid1       pgtype.Int4
	Playerid1points int32
	Playerid2       pgtype.Int4
	Playerid2points int32
	Matchdate       pgtype.Date
	Winnerid        pgtype.Int4
	Group           int32
	Isactive        bool
	ID              int32
}

func (q *Queries) UpdateMatch(ctx context.Context, arg UpdateMatchParams) (Match, error) {
	row := q.db.QueryRow(ctx, updateMatch,
		arg.Seasonid,
		arg.Playerid1,
		arg.Playerid1points,
		arg.Playerid2,
		arg.Playerid2points,
		arg.Matchdate,
		arg.Winnerid,
		arg.Group,
		arg.Isactive,
		arg.ID,
	)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.Seasonid,
		&i.Playerid1,
		&i.Playerid1points,
		&i.Playerid2,
		&i.Playerid2points,
		&i.Matchdate,
		&i.Winnerid,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Group,
	)
	return i, err
}

const updateMatchesBatch = `-- name: UpdateMatchesBatch :exec
UPDATE matches
SET seasonId = m.seasonId,
    playerId1 = m.playerId1,
    playerId1Points = m.playerId1Points,
    playerId2 = m.playerId2,
    playerId2Points = m.playerId2Points,
    matchDate = m.matchDate,
    winnerId = m.winnerId,
    "group" = m."group",
    isActive = m.isActive,
    updatedAt = CURRENT_TIMESTAMP
FROM (SELECT unnest FROM UNNEST ($1::matches[])) AS m
WHERE matches.id = m.id
`

func (q *Queries) UpdateMatchesBatch(ctx context.Context, dollar_1 []interface{}) error {
	_, err := q.db.Exec(ctx, updateMatchesBatch, dollar_1)
	return err
}

const updatePlayer = `-- name: UpdatePlayer :one
UPDATE players
SET name = $1,
    email = $2,
    preferredMatchGroup = $3,
    emailNotificationsEnabled = $4,
    isActive = $5,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $6 AND userId = $7
RETURNING id, userid, name, email, createdat, updatedat, preferredmatchgroup, isactive, emailnotificationsenabled
`

type UpdatePlayerParams struct {
	Name                      string
	Email                     pgtype.Text
	Preferredmatchgroup       pgtype.Int4
	Emailnotificationsenabled bool
	Isactive                  bool
	ID                        int32
	Userid                    pgtype.Int4
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, updatePlayer,
		arg.Name,
		arg.Email,
		arg.Preferredmatchgroup,
		arg.Emailnotificationsenabled,
		arg.Isactive,
		arg.ID,
		arg.Userid,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Email,
		&i.Createdat,
		&i.Updatedat,
		&i.Preferredmatchgroup,
		&i.Isactive,
		&i.Emailnotificationsenabled,
	)
	return i, err
}

const updateSeason = `-- name: UpdateSeason :one
UPDATE seasons
SET name = $1,
    startDate = $2,
    seasonType = $3,
    frequency = $4,
    isActive = $5,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $6 AND userId = $7
RETURNING id, userid, name, startdate, createdat, updatedat, isactive, seasontype, frequency
`

type UpdateSeasonParams struct {
	Name       string
	Startdate  pgtype.Date
	Seasontype string
	Frequency  string
	Isactive   bool
	ID         int32
	Userid     pgtype.Int4
}

func (q *Queries) UpdateSeason(ctx context.Context, arg UpdateSeasonParams) (Season, error) {
	row := q.db.QueryRow(ctx, updateSeason,
		arg.Name,
		arg.Startdate,
		arg.Seasontype,
		arg.Frequency,
		arg.Isactive,
		arg.ID,
		arg.Userid,
	)
	var i Season
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Name,
		&i.Startdate,
		&i.Createdat,
		&i.Updatedat,
		&i.Isactive,
		&i.Seasontype,
		&i.Frequency,
	)
	return i, err
}

const updateUserAppSettings = `-- name: UpdateUserAppSettings :exec
UPDATE users
SET jsonSettings = $1,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateUserAppSettingsParams struct {
	Jsonsettings pgtype.Text
	ID           int32
}

func (q *Queries) UpdateUserAppSettings(ctx context.Context, arg UpdateUserAppSettingsParams) error {
	_, err := q.db.Exec(ctx, updateUserAppSettings, arg.Jsonsettings, arg.ID)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET updatedAt = CURRENT_TIMESTAMP
WHERE email = $1
`

func (q *Queries) UpdateUserPassword(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, updateUserPassword, email)
	return err
}

const updateUserUserSettings = `-- name: UpdateUserUserSettings :exec
UPDATE users
SET jsonSettings = $1,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateUserUserSettingsParams struct {
	Jsonsettings pgtype.Text
	ID           int32
}

func (q *Queries) UpdateUserUserSettings(ctx context.Context, arg UpdateUserUserSettingsParams) error {
	_, err := q.db.Exec(ctx, updateUserUserSettings, arg.Jsonsettings, arg.ID)
	return err
}

const upsertPlayerCustomValue = `-- name: UpsertPlayerCustomValue :one
INSERT INTO player_custom_values (player_id, column_id, value)
VALUES ($1, $2, $3)
ON CONFLICT(player_id, column_id) DO UPDATE SET
    value = excluded.value,
    updatedAt = CURRENT_TIMESTAMP
RETURNING id, player_id, column_id, value, createdat, updatedat
`

type UpsertPlayerCustomValueParams struct {
	PlayerID pgtype.Int4
	ColumnID pgtype.Int4
	Value    pgtype.Text
}

func (q *Queries) UpsertPlayerCustomValue(ctx context.Context, arg UpsertPlayerCustomValueParams) (PlayerCustomValue, error) {
	row := q.db.QueryRow(ctx, upsertPlayerCustomValue, arg.PlayerID, arg.ColumnID, arg.Value)
	var i PlayerCustomValue
	err := row.Scan(
		&i.ID,
		&i.PlayerID,
		&i.ColumnID,
		&i.Value,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}
