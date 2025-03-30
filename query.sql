-- query.sql

-- name: GetUserByEmail :one
SELECT id, email, isVerified -- Add password hash column if/when implemented
FROM users
WHERE email = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    stytchId, stripeId, name, email, phone, country, birthday, lang, isVerified
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUserPassword :exec -- Placeholder for password update logic
UPDATE users
SET -- password_hash = ? -- Add relevant fields
    updatedAt = CURRENT_TIMESTAMP
WHERE email = ?;

-- name: GetPlayers :many
SELECT * FROM players
WHERE userId = ? AND isActive = true; -- Add potential filters based on GetPlayersParams

-- name: CreatePlayer :one
INSERT INTO players (
    userId, name, email, preferredMatchGroup, emailNotificationsEnabled
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE id = ? AND userId = ?;

-- name: UpdatePlayer :one
UPDATE players
SET name = ?,
    email = ?,
    preferredMatchGroup = ?,
    emailNotificationsEnabled = ?,
    isActive = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ? AND userId = ?
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = ? AND userId = ?;

-- name: GetSeasons :many
SELECT * FROM seasons
WHERE userId = ? AND isActive = true;

-- name: CreateSeason :one
INSERT INTO seasons (
    userId, name, startDate, seasonType, frequency
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSeason :one
SELECT * FROM seasons
WHERE id = ? AND userId = ?;

-- name: UpdateSeason :one
UPDATE seasons
SET name = ?,
    startDate = ?,
    seasonType = ?,
    frequency = ?,
    isActive = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ? AND userId = ?
RETURNING *;

-- name: DeleteSeason :exec
DELETE FROM seasons
WHERE id = ? AND userId = ?;

-- name: CreateMatch :one
INSERT INTO matches (
    seasonId, playerId1, playerId1Points, playerId2, playerId2Points, matchDate, winnerId, "group"
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateMatch :one
UPDATE matches
SET seasonId = ?,
    playerId1 = ?,
    playerId1Points = ?,
    playerId2 = ?,
    playerId2Points = ?,
    matchDate = ?,
    winnerId = ?,
    "group" = ?,
    isActive = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteMatch :exec
DELETE FROM matches
WHERE id = ?;

-- name: GetMatch :one
SELECT * FROM matches
WHERE id = ?;

-- name: GetSeasonScoreboard :many
SELECT p.id as player_id, p.name as player_name, COUNT(m.winnerId) as wins
FROM players p
LEFT JOIN matches m ON p.id = m.winnerId AND m.seasonId = ? -- Filter join by season
WHERE p.userId = (SELECT userId FROM seasons WHERE id = ?) -- Ensure player belongs to the season owner
GROUP BY p.id, p.name
ORDER BY wins DESC;

-- name: GetSeasonUpcomingMatches :many
SELECT * FROM matches
WHERE seasonId = ? AND matchDate > CURRENT_TIMESTAMP -- Use CURRENT_TIMESTAMP for timestamp comparison
ORDER BY matchDate ASC
LIMIT 5; -- Limit can be parameterized if needed

-- name: GetPlayerCustomColumns :many
SELECT * FROM player_custom_columns
WHERE is_active = true
ORDER BY display_order;

-- name: CreatePlayerCustomColumn :one
INSERT INTO player_custom_columns (
    name, field_type, description, is_required, display_order
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetPlayerCustomValues :many
SELECT pcv.*, pcc.name as column_name, pcc.field_type
FROM player_custom_values pcv
JOIN player_custom_columns pcc ON pcv.column_id = pcc.id
WHERE pcv.player_id = ?;

-- name: UpsertPlayerCustomValue :one
INSERT INTO player_custom_values (player_id, column_id, value)
VALUES (?, ?, ?)
ON CONFLICT(player_id, column_id) DO UPDATE SET
    value = excluded.value,
    updatedAt = CURRENT_TIMESTAMP
RETURNING *;

-- name: DeletePlayerCustomColumn :exec
DELETE FROM player_custom_columns
WHERE id = ?; -- Add userId check if columns are user-specific

-- name: GetUserAppSettings :one -- Assuming settings are stored in users.jsonSettings
SELECT jsonSettings FROM users
WHERE id = ?;

-- name: UpdateUserAppSettings :exec
UPDATE users
SET jsonSettings = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: GetUserUserSettings :one -- Assuming settings are stored in users.jsonSettings
SELECT jsonSettings FROM users -- Or specific fields if not just JSON
WHERE id = ?;

-- name: UpdateUserUserSettings :exec
UPDATE users
SET jsonSettings = ?, -- Or specific fields
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: GetUserSubscription :one
SELECT subscriptionTier, stripeId -- Add other relevant subscription fields
FROM users
WHERE id = ?;

-- name: DeleteUserSubscription :exec -- Likely involves Stripe API + updating user record
UPDATE users
SET subscriptionTier = 'free', -- Or nullify stripeId
    stripeId = '', -- Or NULL
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :exec
-- Consider implications: deleting related data (players, seasons, matches)? Or just mark inactive?
-- Option 1: Mark inactive
UPDATE users SET isActive = false, updatedAt = CURRENT_TIMESTAMP WHERE id = ?;
-- Option 2: Actual deletion (requires cascading deletes or manual cleanup)
-- DELETE FROM users WHERE id = ?;

-- Add other queries as needed for specific handlers like:
-- PostMatchesUnassignPlayerFromMatch
-- PutMatchesBatches
-- GetPlayersPlayerIdSchedule
-- GetSeasonsTotalAmount
-- GetSeasonsSeasonIdPublicScheduleLink
-- Subscription handling (PostSubscriptions...) - These likely involve Stripe API calls more than direct DB updates initially.
-- PostSupportMessages - Might need a 'support_messages' table.
-- Verification/Magic Link handlers - Might involve temporary token tables or user flags.