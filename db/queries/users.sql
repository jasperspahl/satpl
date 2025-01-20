-- name: GetUserBySpotifyID :one
SELECT * FROM users
WHERE spotify_id = $1;

-- name: CreateUser :one
INSERT INTO users (
  spotify_id,
  display_name,
  email,
  access_token,
  refresh_token
) VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateAccessToken :exec
UPDATE users SET access_token = $2
WHERE id = $1;

-- name: UpdateRefreshToken :exec
UPDATE users SET refresh_token = $2
WHERE id = $1;

-- name: UpdateTokens :exec
UPDATE users SET access_token = $2, refresh_token = $3
WHERE id = $1;

