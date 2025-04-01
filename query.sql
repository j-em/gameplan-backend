-- query.sql with PostgreSQL parameter syntax

-- name: GetUserByEmail :one
SELECT id, email, isVerified
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
    stytchId, stripeId, name, email, phone, country, birthday, lang, isVerified
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET updatedAt = CURRENT_TIMESTAMP
WHERE email = $1;

-- name: GetPlayers :many
SELECT * FROM players
WHERE userId = $1 AND isActive = true;

-- name: CreatePlayer :one
INSERT INTO players (
    userId, name, email, preferredMatchGroup, emailNotificationsEnabled
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE id = $1 AND userId = $2;

-- name: UpdatePlayer :one
UPDATE players
SET name = $1,
    email = $2,
    preferredMatchGroup = $3,
    emailNotificationsEnabled = $4,
    isActive = $5,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $6 AND userId = $7
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1 AND userId = $2;

-- name: GetSeasons :many
SELECT * FROM seasons
WHERE userId = $1 AND isActive = true;

-- name: CreateSeason :one
INSERT INTO seasons (
    userId, name, startDate, seasonType, frequency
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSeason :one
SELECT * FROM seasons
WHERE id = $1 AND userId = $2;

-- name: UpdateSeason :one
UPDATE seasons
SET name = $1,
    startDate = $2,
    seasonType = $3,
    frequency = $4,
    isActive = $5,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $6 AND userId = $7
RETURNING *;

-- name: DeleteSeason :exec
DELETE FROM seasons
WHERE id = $1 AND userId = $2;

-- name: CreateMatch :one
INSERT INTO matches (
    seasonId, playerId1, playerId1Points, playerId2, playerId2Points, matchDate, winnerId, "group"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateMatch :one
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
RETURNING *;

-- name: DeleteMatch :exec
DELETE FROM matches
WHERE id = $1;

-- name: GetMatch :one
SELECT * FROM matches
WHERE id = $1;

-- name: GetSeasonScoreboard :many
SELECT p.id as player_id, p.name as player_name, COUNT(m.winnerId) as wins
FROM players p
LEFT JOIN matches m ON p.id = m.winnerId AND m.seasonId = $1
WHERE p.userId = (SELECT s.userId FROM seasons s WHERE s.id = $2)
GROUP BY p.id, p.name
ORDER BY wins DESC;

-- name: GetSeasonUpcomingMatches :many
SELECT * FROM matches
WHERE seasonId = $1 AND matchDate > CURRENT_TIMESTAMP
ORDER BY matchDate ASC
LIMIT 5;

-- name: GetPlayerCustomColumns :many
SELECT * FROM player_custom_columns
WHERE is_active = true
ORDER BY display_order;

-- name: CreatePlayerCustomColumn :one
INSERT INTO player_custom_columns (
    name, field_type, description, is_required, display_order
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPlayerCustomValues :many
SELECT pcv.*, pcc.name as column_name, pcc.field_type
FROM player_custom_values pcv
JOIN player_custom_columns pcc ON pcv.column_id = pcc.id
WHERE pcv.player_id = $1;

-- name: UpsertPlayerCustomValue :one
INSERT INTO player_custom_values (player_id, column_id, value)
VALUES ($1, $2, $3)
ON CONFLICT(player_id, column_id) DO UPDATE SET
    value = excluded.value,
    updatedAt = CURRENT_TIMESTAMP
RETURNING *;

-- name: DeletePlayerCustomColumn :exec
DELETE FROM player_custom_columns
WHERE id = $1;

-- name: GetUserAppSettings :one
SELECT jsonSettings FROM users
WHERE id = $1;

-- name: UpdateUserAppSettings :exec
UPDATE users
SET jsonSettings = $1,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: GetUserUserSettings :one
SELECT jsonSettings FROM users
WHERE id = $1;

-- name: UpdateUserUserSettings :exec
UPDATE users
SET jsonSettings = $1,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: GetUserSubscription :one
SELECT subscriptionTier, stripeId
FROM users
WHERE id = $1;

-- name: DeleteUserSubscription :exec
UPDATE users
SET subscriptionTier = 'free',
    stripeId = '',
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
UPDATE users SET isActive = false, updatedAt = CURRENT_TIMESTAMP WHERE id = $1;

-- name: UpdateMatchesBatch :exec
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
FROM (SELECT * FROM UNNEST ($1::matches[])) AS m
WHERE matches.id = m.id;